package cache

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice()[]byte {
	return cloneBytes(v.b)
}

func (v ByteView)String()string{
	return string(v.b)
}

func cloneBytes(v []byte)[]byte  {
	c := make([]byte,len(v))
	copy(c,v)
	return c
}
