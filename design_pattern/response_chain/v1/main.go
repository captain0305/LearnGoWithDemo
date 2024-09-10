package main

import (
	"fmt"
)

// Request 请求结构体
type Request struct {
	Type int
	Data string
}

// Handler 处理器接口
type Handler interface {
	HandleRequest(req *Request) bool
	SetNext(handler Handler)
}

// ConcreteHandlerA 具体处理器A
type ConcreteHandlerA struct {
	nextHandler Handler
}

// HandleRequest 处理请求
func (h *ConcreteHandlerA) HandleRequest(req *Request) bool {
	if req.Type == 1 {
		fmt.Printf("ConcreteHandlerA handled request with data: %s\n", req.Data)
		return true
	}
	if h.nextHandler != nil {
		return h.nextHandler.HandleRequest(req)
	}
	return false
}

// SetNext 设置下一个处理器
func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.nextHandler = handler
}

// ConcreteHandlerB 具体处理器B
type ConcreteHandlerB struct {
	nextHandler Handler
}

// HandleRequest 处理请求
func (h *ConcreteHandlerB) HandleRequest(req *Request) bool {
	if req.Type == 2 {
		fmt.Printf("ConcreteHandlerB handled request with data: %s\n", req.Data)
		return true
	}
	if h.nextHandler != nil {
		return h.nextHandler.HandleRequest(req)
	}
	return false
}

// SetNext 设置下一个处理器
func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.nextHandler = handler
}

func main() {
	// 创建处理器链
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerA.SetNext(handlerB)

	// 发送请求
	requests := []*Request{
		{Type: 1, Data: "Data1"},
		{Type: 2, Data: "Data2"},
		{Type: 3, Data: "Data3"},
	}

	for _, req := range requests {
		handlerA.HandleRequest(req)
	}
}
