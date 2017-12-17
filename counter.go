package counter

import "sync/atomic"

// Countor はインクリメントしかできないカウンターオブジェクトです。
//
// Cはゴルーチンセーフな実装ですがNの10倍時間がかかります。
type Countor interface {
	Up(n uint64)
	Count() uint64
}

var (
	_ Countor = (*C)(nil)
	_ Countor = (*N)(nil)
)

// C カウンターでUint64のラップです。
type C struct {
	count uint64
}

// Up では指定した数だけカウントアップします。
func (c *C) Up(n uint64) {
	atomic.AddUint64(&c.count, n)
}

// Count では、カウントアップした数を
func (c C) Count() uint64 {
	return c.count
}

// N カウンターでUint64のラップです。
type N struct {
	count uint64
}

// Up では指定した数だけカウントアップします。
func (c *N) Up(n uint64) {
	c.count += n
}

// Count では、カウントアップした数を
func (c N) Count() uint64 {
	return c.count
}
