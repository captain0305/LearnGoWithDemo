package escape

import (
	"bytes"
	"sync"
)

var bufferPool sync.Pool
var smallBuffer sync.Pool

func init() {
	bufferPool.New = func() interface{} {
		return &bytes.Buffer{}
	}
	smallBuffer.New = func() interface{} {
		return &bytes.Buffer{}
	}
}

// GetBuf 获取缓存
func GetBuf() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

// PutBuf 归还缓存
func PutBuf(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}

// GetSmallBuf 获取缓存
func GetSmallBuf() *bytes.Buffer {
	return smallBuffer.Get().(*bytes.Buffer)
}

// PutSmallBuf 归还缓存
func PutSmallBuf(buf *bytes.Buffer) {
	buf.Reset()
	smallBuffer.Put(buf)
}
