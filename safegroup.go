package safegroup

import (
	"fmt"
	"runtime"

	"golang.org/x/sync/errgroup"
)

// PanicBufLen panic调用栈日志buffer大小，默认2048
var PanicBufLen = 2048

// safegroup panic 安全的group结构定义
type safegroup struct {
	g errgroup.Group
}

// Go 启动一个协程，会接住panic
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

// Wait等待协程结束
func (s *safegroup) Wait() error {
	return s.g.Wait()
}

// NewSafeGroup 生成一个实例
func NewSafeGroup() *safegroup {
	return &safegroup{
		g: errgroup.Group{},
	}
}
