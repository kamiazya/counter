package counter

import (
	"sync"
	"testing"
)

func TestCountor(t *testing.T) {
	count := 10000
	t.Run("C", func(t *testing.T) {
		wg := new(sync.WaitGroup)
		c := new(C)
		for i := 0; count > i; i++ {
			wg.Add(1)
			go func() {
				c.Up(1)
				wg.Done()
			}()
		}
		wg.Wait()
		if c.Count() != uint64(count) {
			t.Error("atomicな操作ができていません。\nwant: ", uint64(count), "actual: ", c.Count())
		}
		t.Log("C構造体ではatomicなができます。\nwant: ", uint64(count), "actual: ", c.Count())
	})
	t.Run("N", func(t *testing.T) {
		wg := new(sync.WaitGroup)
		n := new(N)
		for i := 0; count > i; i++ {
			wg.Add(1)
			go func() {
				n.Up(1)
				wg.Done()
			}()
		}
		wg.Wait()
		t.Log("N構造体ではatomicな操作ができません。\nwant: ", uint64(count), "actual: ", n.Count())
	})
}

func BenchmarkCountor(b *testing.B) {
	b.Run("C", func(b *testing.B) {
		c := new(C)
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			c.Up(1)
		}
	})
	b.Run("N", func(b *testing.B) {
		n := new(N)
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			n.Up(1)
		}
	})
}
