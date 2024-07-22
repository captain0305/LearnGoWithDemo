package escape

import (
	"bytes"
)

// escapeValue 转义后的替代值
const (
	separatorByte byte = '|'
	escapeValue        = '&'
)

var escapeMap = map[byte]string{
	'|': "&#x7C;",
}
var unescapeMap = map[string]byte{
	"&#x7C;": '|',
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

// UnescapeString 反转义字符串
// 参数：data 待反转义数据
// 返回：反转义后的数据
func UnescapeString(data string) string {
	dataLen := len(data)
	if dataLen <= 0 {
		return ""
	}
	buf := GetBuf()
	defer PutBuf(buf)
	binData := []byte(data)
	index := 0
	for index = 0; index < dataLen-1; index++ {
		unescapeByte4String(buf, binData, &index)
	}
	if index == dataLen-1 {
		value := binData[index]
		buf.WriteByte(value)
	}
	ret := buf.String()
	return ret
}

func unescapeByte4String(buf *bytes.Buffer, binData []byte, index *int) {
	value := binData[*index]
	switch value {
	case escapeValue:
		buffer := GetSmallBuf()
		defer PutSmallBuf(buffer)
		if *index+6 <= len(binData) {
			buffer.Write(binData[*index:(*index + 6)])
		}
		unescapeByte(buf, value, buffer.String(), index)
	default:
		buf.WriteByte(value)
	}
}

func unescapeByte(buf *bytes.Buffer, value byte, nextValue string, i *int) {
	unescapeByte, ok := unescapeMap[nextValue]
	if ok {
		buf.WriteByte(unescapeByte)
		*i = *i + 5
	} else {
		buf.WriteByte(value)
	}
}
