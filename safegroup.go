package safegroup

import (
	"fmt"
	"runtime"

	"golang.org/x/sync/errgroup"
)

// PanicBufLen panic stack buffer size，default size is 2048
var PanicBufLen = 2048

// safegroup define
type safegroup struct {
	g errgroup.Group
}

// Go run a goroutine with recover panic
func (s *safegroup) Go(f func() error) {
	s.g.Go(func() error {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, PanicBufLen)
				buf = buf[:runtime.Stack(buf, false)]
				fmt.Printf("[PANIC]:\t%+v\n%s\n", r, buf)
			}
		}()
		return f()
	})
}

// Wait wait goroutines finish
func (s *safegroup) Wait() error {
	return s.g.Wait()
}

// NewSafeGroup new a safegroup
func NewSafeGroup() *safegroup {
	return &safegroup{
		g: errgroup.Group{},
	}
}
