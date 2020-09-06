package fixer

type BufferSpy struct {
	buffer    []byte
	writeChan chan (bool)
}

func (b *BufferSpy) Write(p []byte) (n int, err error) {
	b.buffer = p
	b.writeChan <- true
	return len(p), err
}
func (b *BufferSpy) String() string {
	return string(b.buffer)
}

func NewBufferSpy() *BufferSpy {
	ch := make(chan (bool), 1)
	return &BufferSpy{
		writeChan: ch,
	}
}
