/*
 * @Author: your name
 * @Date: 2021-02-03 19:50:48
 * @LastEditTime: 2021-02-08 11:24:58
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /learn/BRPC_learn/MPSC.go
 */
package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type TaskNode struct {
	data interface{}
	Next *TaskNode
}

var UNCONNECTED *TaskNode = new(TaskNode)

func NewExecutionQueue(_func func(interface{})) *ExecutionQueue {
	return &ExecutionQueue{
		head:          nil,
		_execute_func: _func,
		locker:        sync.Mutex{},
		pool: &sync.Pool{New: func() interface{} {
			return new(TaskNode)
		}},
	}
}

type ExecutionQueue struct {
	head          *TaskNode         `json:"head"`
	_execute_func func(interface{}) `json:"_execute_func"`
	locker        sync.Mutex        `json:"locker"`
	pool          *sync.Pool        `json:"pool"`
}

func (ex *ExecutionQueue) AddTaskNode(data interface{}) {
	node := ex.pool.Get().(*TaskNode)
	// node := new(TaskNode)
	node.data = data
	node.Next = UNCONNECTED

	preHead := atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&ex.head)), unsafe.Pointer(node))
	// ex.locker.Lock()
	// preHead := ex.head
	// ex.head = node
	// ex.locker.Unlock()

	if preHead != nil {
		node.Next = (*TaskNode)(preHead)
		return
	}

	node.Next = nil
	// 任务不多直接执行，防止线程切换
	ex._execute_func(node.data)
	if !ex.moreTasks(node) {
		// ex.pool.Put(node)
		// atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&ex.head)), nil)
		return
	}
	go ex.exectueTasks(node)

}

func (this *ExecutionQueue) toString() string {
	rs, _ := json.Marshal(this)
	return string(rs)
}

func (ex *ExecutionQueue) moreTasks(oldNode *TaskNode) bool {

	newHead := oldNode

	if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&ex.head)), unsafe.Pointer(newHead), nil) {
		return false
	}
	newHead = (*TaskNode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&ex.head))))

	// ex.locker.Lock()
	// if ex.head == newHead {
	// 	ex.head = nil
	// 	ex.locker.Unlock()
	// 	return false
	// } else {
	// 	// 有新插入值
	// 	newHead = ex.head
	// }
	// ex.locker.Unlock()

	// newTail 为结尾
	var tail *TaskNode

	p := newHead
	for {
		for {
			if p.Next != UNCONNECTED {
				break
			} else {
				runtime.Gosched()
			}
		}
		saved_next := p.Next
		p.Next = tail
		tail = p
		p = saved_next

		if p == oldNode {
			oldNode.Next = tail
			return true
		}
	}
}

func (ex *ExecutionQueue) exectueTasks(taskNode *TaskNode) {
	// defer singalexit.Done()
	// singalexit.Add(1)
	for {
		tmp := taskNode

		taskNode = taskNode.Next
		tmp.Next = nil
		ex.pool.Put(tmp)
		ex._execute_func(taskNode.data)

		if taskNode.Next == nil && !ex.moreTasks(taskNode) {
			// ex.pool.Put(taskNode)
			// atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&ex.head)), nil)
			return
		}
	}
}

var count int64 = 0

func print(data interface{}) {
	// a := count
	// _ = a
	_ = data.(int) * data.(int)
	atomic.AddInt64(&count, 1)
	// fmt.Println(data.(int))
}
func Test1() {
	var singalexit = sync.WaitGroup{}
	ex := NewExecutionQueue(print)
	start := time.Now()

	for k := 0; k < 20; k++ {
		for i := 0; i < 10000; i++ {
			singalexit.Add(1)
			go func(i int, singalexit *sync.WaitGroup) {
				defer singalexit.Done()
				for j := 0; j < 90; j++ {
					ex.AddTaskNode(i*100 + j)
				}

			}(i, &singalexit)

		}
	}
	singalexit.Wait()
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed/20)
	time.Sleep(2 * time.Second)
	fmt.Println(atomic.LoadInt64(&count))
}

func Test2() {
	var singalexit sync.WaitGroup
	data := make(chan int, 2000)
	var count1 int64 = 0
	go func() {
		for {
			<-data
			atomic.AddInt64(&count1, 1)
		}

	}()
	start := time.Now()
	func() {
		for i := 0; i < 10000; i++ {
			singalexit.Add(1)
			go func(i int) {

				defer singalexit.Done()
				for j := 0; j < 90; j++ {
					data <- (i*100 + j)
				}
			}(i)
		}
	}()
	singalexit.Wait()
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
	time.Sleep(2 * time.Second)
	fmt.Println(atomic.LoadInt64(&count1))

}

func main() {
	for i := 0; i < 10; i++ {
		count = 0
		Test1()
	}
	// for i := 0; i < 10; i++ {
	// 	Test2()
	// }

	// Test2()
}
