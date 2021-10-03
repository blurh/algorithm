package structure

import (
    "errors"
    "sync"
)

type stack struct {
    data []interface{}
    lock sync.RWMutex
}

func InitStack() *stack {
    s := new(stack)
    return s
}

func (s *stack) Push(value interface{}) bool {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.data = append(s.data, value)
    return true
}

func (s *stack) Pop() (interface{}, error) {
    s.lock.Lock()
    defer s.lock.Unlock()
    if len(s.data) <= 0 {
        errMsg := errors.New("stack is empty")
        return nil, errMsg
    }
    popValue := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return popValue, nil
}

func (s *stack) Len() int {
    s.lock.Lock()
    defer s.lock.Unlock()
    return len(s.data)
}

func (s *stack) Clear() bool {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.data = []interface{}{}
    return true
}
