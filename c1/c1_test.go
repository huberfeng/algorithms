package c1

import (
	"testing"
	"fmt"
		)

func TestStack_Stack(t *testing.T) {
	s := new(Stack)
	s.Stack()
	fmt.Println(s.Size())
	fmt.Println(s.Isempty())

	s.Push(1)
	s.Push(2)
	fmt.Println(s.Size())
	fmt.Println(s.Isempty())
	fmt.Println(s.Pop())

	fmt.Println(s.Size())
	fmt.Println(s.Isempty())
	fmt.Println(s.Pop())

}

func TestStack_StackP(t *testing.T) {
	s := new(StackP)
                                 //expect
	fmt.Println(s.Size())        //0
	fmt.Println(s.Isempty())    //true

	s.Push(1)
	s.Push(2)
	fmt.Println(s.Size())      //2
	fmt.Println(s.Isempty())   //false
	fmt.Println(s.Pop())       //2

	fmt.Println(s.Size())      //1
	fmt.Println(s.Isempty())   //false
	fmt.Println(s.Pop())       //1
	fmt.Println(s.Size())      //0
	fmt.Println(s.Pop())       //nil
	fmt.Println(s.Size())      //0
}

func TestBag_Bag(t *testing.T) {
	b := new(Bag)
	b.Bag()
	fmt.Println(b.Size())
	fmt.Println(b.Isempty())

	b.Add("qwer")
	fmt.Println(b.Size())
	fmt.Println(b.Isempty())
}

func TestQueue_Queue(t *testing.T) {
	q := new(Queue)
	q.Queue()
	fmt.Println(q.Size())
	fmt.Println(q.Isempty())

	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Size())
	fmt.Println(q.Isempty())

	fmt.Println(q.Dequeue())
	fmt.Println(q.Size())
	fmt.Println(q.Isempty())

}

func TestQueue_QueueL(t *testing.T) {
	q := new(QueueL)
	fmt.Println(q.Size())
	fmt.Println(q.Isempty())

	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Size())
	fmt.Println(q.Isempty())

	fmt.Println(q.Dequeue())
	fmt.Println(q.Size())
	fmt.Println(q.Isempty())

	fmt.Println(q.Dequeue())
	fmt.Println(q.Size())
	fmt.Println(q.Isempty())

	fmt.Println(q.Dequeue())
	fmt.Println(q.Size())
	fmt.Println(q.Isempty())
}

func TestCn3(t *testing.T) {
	a := []int{1,2,3,4,5}
	Cn3(a)
}

func TestAn3(t *testing.T) {
	a := []int{1,2,3,4,5}
	An3(a)
}

