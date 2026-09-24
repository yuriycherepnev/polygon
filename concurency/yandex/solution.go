package main

const MaxItems = 9999

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

func Pipe(p Producer, c Consumer) error {
	buf := make([]any, 0, MaxItems)
	cookieQnt := 0

	for {
		items, cookie, err := p.Next()

		if len(buf)+len(items) > MaxItems {
			err = c.Process(buf)
			if err != nil {
				return err
			}
			err = p.Commit(cookieQnt)
			if err != nil {
				return err
			}
			buf = buf[:0]
			cookieQnt = 0
		}
		buf = append(buf, items...)
		cookieQnt += cookie
	}

}
