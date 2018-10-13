package c1


//栈 接口定义
type S interface {
	Push(p interface{})
	Pop() interface{}
	Isempty() bool
	Size() int
}

// 栈，数组实现
type  Stack struct {
	items []interface{}
}

func (s *Stack) Stack() {
	s.items = make([]interface{}, 0, 5)
}

func (s *Stack) Push(p interface{}) {
	s.items = append(s.items, p)
}

func (s *Stack) Pop() interface{}{
	if s.Isempty() {
		return nil
	}
	item := s.items[s.Size() -1]
	s.items = s.items[:s.Size() - 1]
	return item
}

func (s *Stack) Isempty() bool{
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return len(s.items)
}

//栈，链表实现
type StackP struct {
	first *node    //栈顶指针
	size int       //栈大小
}
//栈节点
type node struct {
	item interface{}
	next *node
}

func (s *StackP) Push(p interface{}) {
	oldNode := s.first
	newNode := new(node)
	newNode.item = p
	newNode.next = oldNode
	s.first = newNode
	s.size ++
}

func (s *StackP) Pop() interface{}{
	if s.first ==  nil {
		return nil
	}
	item := s.first.item
	s.first = s.first.next
	s.size --
	return item
}

func (s *StackP) Isempty() bool{
	return s.first == nil
}

func (s *StackP) Size() int {
	return s.size
}


// 背包 切片实现
type Bag struct {
	bag []interface{}
}

func (b *Bag) Bag() {
	b.bag = make([]interface{}, 0, 5)
}

func (b *Bag) Add(p interface{}) {
	b.bag = append(b.bag, p)
}

func (b *Bag) Isempty() bool {
	return len(b.bag) == 0
}

func (b *Bag) Size() int {
	return len(b.bag)
}


//先进先出队列接口定义
type Q interface {
	Enqueue(p interface{})
	Dequeue() interface{}
	Isempty() bool
	Size() int
}
// 队列，数组实现
type Queue struct{
	queue []interface{}
}

func (q *Queue) Queue() {
	q.queue = make([]interface{}, 0, 5)
}

func (q *Queue) Enqueue(p interface{}) {
	q.queue = append(q.queue, p)
}

func (q *Queue) Dequeue() interface{}{
	if q.Isempty() {
		return nil
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *Queue) Isempty() bool{
	return len(q.queue) == 0
}

func (q *Queue) Size() int {
	return len(q.queue)
}

//链表实现

type Qnode struct {
	item interface{}
	next *Qnode
}
type QueueL struct {
	first *Qnode    //队头
	last *Qnode     //队尾
	size int
}

func (q *QueueL) Enqueue(p interface{}) {
	newNode := new(Qnode)
	newNode.item = p
	if q.first == nil {        //当只有一个节点时，队首即是队尾
		q.first = newNode
		q.last = newNode
	}else {
	    q.last.next = newNode
	    q.last = newNode
	}
	q.size ++
}

func (q *QueueL) Dequeue() interface{}{
	if q.Isempty() {
		return nil
	}
	item := q.first.item
	q.first = q.first.next
	if q.first == nil {     //队列出完，队尾指针置为nil
		q.last = nil
	}
	q.size --
	return item
}

func (q *QueueL) Isempty() bool{
	return q.first == nil
}

func (q *QueueL) Size() int {
	return q.size
}