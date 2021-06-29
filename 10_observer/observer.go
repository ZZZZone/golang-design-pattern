package observer

import "fmt"

// subject 相当于老师
type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

// 内容变更后,学生关注到讲课内容变更， 但是老师不需要关心知道内容变更的对象具体要做什么
func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

// 老师讲课， 内容发生变更
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

// Reader 可以看作是学生
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

// 学生记录笔记
func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}

// 可以有多类观察者, 旁听的老师，
type Lisener struct {
}

// 旁听的老师 记录讲课内容， 给讲课老师打分
func (r *Lisener) Update(s *Subject) {
	fmt.Printf("receive %s, make point\n", s.context)
}
