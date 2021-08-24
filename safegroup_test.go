package safegroup

import (
	"testing"
)

func TestNewSafeGroup(t *testing.T) {
	s := NewSafeGroup()
	s.Go(func() error {
		t.Fatal("test 1...")
		return nil
	})

	s.Go(func() error {
		t.Fatal("test 2...")
		return nil
	})

	s.Wait()

	t.Fatal("final ...")
}

func TestGo(t *testing.T) {
	s := NewSafeGroup()
	s.Go(func() error {
		t.Fatal("test 1...")
		return nil
	})

	s.Go(func() error {
		t.Fatal("test 2...")
		return nil
	})
	s.Go(func() error {
		panic("this is a panic")
	})

	s.Wait()

	t.Fatal("final ...")
}
