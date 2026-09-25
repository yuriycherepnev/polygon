package main

import (
	"context"
	"time"
)

const MaxItems = 9999

var (
	ticker = time.NewTicker(time.Second * 60)
	bufCh  = make(chan BatchMsg, 10)
)

type Producer interface {
	// Next returns:
	// - batch of items to be processed
	// - cookie to be committed when processing is done
	// - error
	Next() (items []any, cookie int, err error)
	// Commit is used to mark data batch as processed
	Commit(cookie int) error
}

type Consumer interface {
	Process(items []any) error
}

type Batch struct {
	buf     []any
	cookies []int
}

type BatchMsg struct {
	items  []any
	cookie int
	err    error
}

func NewBatch() *Batch {
	return &Batch{
		buf:     make([]any, 0, MaxItems),
		cookies: make([]int, 0),
	}
}

func (b *Batch) Add(batch BatchMsg) {
	b.buf = append(b.buf, batch.items...)
	b.cookies = append(b.cookies, batch.cookie)
}

func (b *Batch) Flush(p Producer, c Consumer) error {
	err := c.Process(b.buf)
	if err != nil {
		return err
	}
	for _, v := range b.cookies {
		err = p.Commit(v)
		if err != nil {
			return err
		}
	}
	b.buf = b.buf[:0]
	b.cookies = b.cookies[:0]
	return nil
}

func Pipe(ctx context.Context, p Producer, c Consumer) error {
	b := NewBatch()
	defer ticker.Stop()

	go func() {
		for {
			items, cookie, err := p.Next()
			select {
			case bufCh <- BatchMsg{items, cookie, err}:
			case <-ctx.Done():
				return
			}
			if err != nil {
				return
			}
		}
	}()

	go func() {
		select {
		case batchMsg, ok := <-bufCh:
			if !ok {
				return
			}
			if len(b.buf)+len(batchMsg.items) > MaxItems {
				err := b.Flush(p, c)
				if err != nil {
					return
				}
			}
			b.Add(batchMsg)
		case <-ctx.Done():
			_ = b.Flush(p, c)
			return
		}

	}()

	return nil
}
