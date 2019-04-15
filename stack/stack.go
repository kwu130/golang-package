package stack

import (
    "fmt"
    "errors"
    "log"
)

func init() {
    log.SetPrefix("STACK: ")
    log.SetFlags(log.Llongfile)
}

type stack interface{
    Push(interface{})
    Pop()
    Top() (interface{}, error)
    IsEmpty() bool
    Clear()
    String() string
}

type Stack struct {
    size int
    stack []interface{} 
}

func New() Stack {
    return Stack{}
}

func push(s *Stack, a interface{}) {
    (*s).stack = append((*s).stack, a)
    (*s).size++
}

func (s *Stack)Push(a interface{}) {
    push(s, a)
}

func pop(s *Stack) error {
    if isempty(*s) {
        return errors.New("pop() is failed, stack is empty")
    }
    (*s).stack = (*s).stack[:(*s).size - 1]
    (*s).size--
    return nil
}

func (s *Stack)Pop() {
    if err := pop(s); err != nil {
        log.Fatal(err)
    }
}

func top(s Stack) (interface{}, error) {
    if isempty(s) {
        return nil, errors.New("top() is failed, stack is empty")
    }
    return  s.stack[s.size - 1], nil
}

func (s Stack)Top() interface{} {
    if val, err := top(s); err != nil {
        log.Fatal(err)
    } else {
        return val
    }
    return nil
}

func isempty(s Stack) bool {
    return s.size == 0
}

func (s Stack)IsEmpty() bool {
    return isempty(s)
}

func clear(s *Stack) {
    if (*s).size == 0 {
        return 
    } else {
        (*s).stack = (*s).stack[:0]
        (*s).size = 0
    }
}

func (s *Stack)Clear() {
    clear(s)
}

func (s Stack)String() string {
    return fmt.Sprintf("%v",s.stack)
}