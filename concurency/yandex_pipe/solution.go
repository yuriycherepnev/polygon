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
