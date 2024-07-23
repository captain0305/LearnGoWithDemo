package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

func main() {
	s := `hello|world\n`
	all := strings.ReplaceAll(s, "|", "&#x7C;")
	all = strings.ReplaceAll(all, `\n`, "&#x0A;")
	fmt.Println(all)
	replaceAll := strings.ReplaceAll(all, "&#x7C;", "|")
	replaceAll = strings.ReplaceAll(replaceAll, "&#x0A;", `\n`)
	fmt.Println(replaceAll)
}

var escapeMap = map[byte]string{
	'|': "&#x7C;",
}

var unEscapeMap = map[string]byte{
	"&#x7C;": '|',
}
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

// EscapeString 转义字符串
// 参数：data 待转义数据
// 返回：转义后的数据
func EscapeString(data string) string {
	buf := GetBuf()
	escapeString(buf, data)
	ret := buf.String()
	PutBuf(buf)
	return ret
}

// escapeString 转义字符串
func escapeString(buf *bytes.Buffer, data string) int {
	dataLen := len(data)
	if dataLen <= 0 {
		return 0
	}
	lastEscapeIndex := 0
	for i := 0; i < dataLen; i++ {
		escapeByte(buf, data, &i, &lastEscapeIndex)
	}
	if lastEscapeIndex != dataLen {
		buf.WriteString(data[lastEscapeIndex:])
	}
	return len(buf.Bytes())
}

func escapeByte(buf *bytes.Buffer, data string, dataIndex *int, lastEscapeIndex *int) {
	key := data[*dataIndex]
	value, ok := escapeMap[key]
	if ok {
		buf.WriteString(data[*lastEscapeIndex:*dataIndex])
		buf.Write([]byte(value))
		*lastEscapeIndex = *dataIndex + 1
	}
}
