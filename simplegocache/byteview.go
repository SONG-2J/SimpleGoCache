package simplegocache

type ByteView struct {
	b []byte // 存储真实的缓冲值
}

// Len 返回ByteView的长度
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice 以字节片的形式返回数据的副本
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String 以字符串形式返回数据，如果需要可以进行复制。
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
