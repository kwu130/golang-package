package queue

import (
    "fmt"
    "log"
    "errors"
)
//初始化日志输出
func init() {
    log.SetPrefix("QUEUE: ")
    log.SetFlags(log.Llongfile)
}

type Queue struct {
    head int
    tail int
    queue []interface{}
}

type queue interface {
    Push(interface{})
    Pop()
    Front() interface{}
    Back() interface{}
    IsEmpty() bool
    Size() int
    Clear()
    String() string
}

func New() Queue {
    return Queue{}
}
//在队尾插入一个元素
func push(q *Queue, a interface{}) {
    (*q).queue = append((*q).queue, a)
    (*q).tail++
    //while the number of empty space more than 2 times of non-empty space, moving element to save memory
    //
    if ((*q).tail - (*q).head) * 2 < (*q).head {
        (*q).queue = (*q).queue[(*q).head:(*q).tail]
        (*q).tail = (*q).tail - (*q).head
        (*q).head = 0
    }
}

func (q *Queue)Push(a interface{}) {
    push(q, a)
}
//从队首删除一个元素
func pop(q *Queue) error {
    if isempty(*q) {
        return errors.New("pop() is failed, queue is empty")
    }
    (*q).head++

    return nil
}

func (q *Queue)Pop() {
    if err := pop(q); err != nil {
        log.Fatal(err)
    }
}
//返回队列第一个元
func front(q Queue) (interface{}, error) {
    if isempty(q) {
        return nil, errors.New("front() is failed, queue is empty")
    }
    return q.queue[q.head], nil
}

func (q Queue)Front() interface{} {
    if val, err := front(q); err != nil {
        log.Fatal(err)
    } else {
        return val
    }
    return nil
}
//返回队列最后一个元素
func back(q Queue) (interface {}, error) {
    if isempty(q) {
        return nil, errors.New("back() is failed, queue is empty")
    }
    return q.queue[q.tail-1], nil
}

func (q Queue)Back() interface{} {
    if val, err := back(q); err != nil {
        log.Fatal(err)
    } else {
        return val
    }
    return nil
}
//判断队列是否为空
func isempty(q Queue) bool {
    if q.head >= q.tail {
        return true
    }
    return false
}

func (q Queue)IsEmpty() bool {
    return isempty(q)
}
//返回队列长度
func size(q Queue) int {
    if q.head > q.tail {
        return -1
    }
    return q.tail - q.head
}

func (q Queue)Size() int {
    if length := size(q); length == -1 {
        log.Fatal("queue is overflow")
    } else {
        return length
    }
    return 0
}
//清空队列
func clear(q *Queue) {
    (*q).queue = (*q).queue[:0]
    (*q).head = 0
    (*q).tail = 0
}

func (q *Queue)Clear() {
    clear(q)
}

func (q Queue)String() string {
    return fmt.Sprintf("%v", q.queue[q.head:q.tail])
}
//测试用
func output(q Queue) {
    fmt.Printf("head:%d tail:%d val:%v\n", q.head, q.tail, q.queue[q.head:q.tail])
}
