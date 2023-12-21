package linkedlist

// Struct para uma singly linked list
type Node struct {
	next *Node
	chave float64
}

type LinkedList struct {
	head *Node
}

// Struct para uma circular linked list
type CLinkedList struct {
	head *Node
}

func CreateCircularLinkedList(key float64) *CLinkedList {
	head := &Node{chave: key}
	cll := &CLinkedList{head}
	head.next = cll.head

	return cll
}

func (cl *CLinkedList) InsertKeyCircularLinkedList(key float64) {
	if cl.head == nil {
		node := &Node{chave: key}
		cl.head = node
		node.next = cl.head

		return
	}

	u := cl.head
	for ; u.next != cl.head ; {
		u = u.next
	}

	u.next = &Node{next: cl.head, chave: key}
}

func (cl *CLinkedList) ShowKeysCircularLinkedList() []float64 {
	if cl == nil || cl.head == nil {
		return []float64{}
	}

	out := make([]float64, 0)
	u := cl.head
	for ; u.next != cl.head ; {
		out = append(out, u.chave)
		u = u.next
	}

	out = append(out, u.chave)
	return out
}

func (cl *CLinkedList) ListLenghtCircularLinkedList() int {
	if cl == nil {
		return 0
	}

	count := 1
	u := cl.head
	for ; u.next != cl.head ; {
		count++
		u = u.next
	}

	return count
}

func (cl *CLinkedList) SearchKeyCircularLinkedList(key float64) bool {
	if cl == nil {
		return false
	}

	u := cl.head
	for u.next != cl.head {
		if u.chave == key {
			return true
		}
		u = u.next
	}

	if u.chave == key {
		return true
	}

	return false
}

func (cl *CLinkedList) DeleteKeyCircularLinkedList(key float64) bool {
	if !cl.SearchKeyCircularLinkedList(key) {
		return false
	}

	// Caso em que a chave a ser deletada esta na cabeca
	if cl.head.chave == key {
		if cl.ListLenghtCircularLinkedList() == 1 {
			cl.head = nil
			cl = nil
			return true
		}

		u := cl.head
		for u.next != cl.head {
			u = u.next
		}
		u.next = cl.head.next
		cl.head = cl.head.next
		return true
	}

	u := cl.head
	v := cl.head.next

	for v.chave != key {
		u = v
		v = v.next
	}

	u.next = v.next
	return true
}
func CreateLinkedList(key float64) *LinkedList {
	return &LinkedList{&Node{chave: key}}
}

func (l *LinkedList) InsertKey(key float64) {
	if l.head == nil {
		l.head = &Node{chave: key}
		return
	}

	u := l.head
	for ; u.next != nil ; {
		u = u.next
	}

	u.next = &Node{chave: key}
}

func (l *LinkedList) ShowKeys() ([]float64) {
	if l == nil {
		out := make([]float64, 0)
		return out
	}

	out := make([]float64, 0)
	u := l.head
	for ; u != nil ; {
		out = append(out, u.chave)
		u = u.next
	}

	return out
}

func (l *LinkedList) ListLenght() int {
	if l == nil {
		return 0
	}

	count := 1
	u := l.head

	for ; u.next != nil ; {
		u = u.next
		count++
	}

	return count
}

func (l *LinkedList) SearchKey(key float64) bool {
	if l == nil {
		return false
	}

	u := l.head
	for ; u.next != nil ; {
		if u.chave == key {
			return true
		}
		u = u.next
	}

	if u.chave == key {
		return true
	}

	return false
}

func (l *LinkedList) DeleteKey(key float64) bool {
	if !l.SearchKey(key) {
		return false
	}

	// Caso em que a chave a ser deletada está na cabeca
	if l.head.chave == key {
		l.head = l.head.next
		return true
	}

	u := l.head
	v := u.next
	for v.chave != key {
		u = v
		v = u.next
	}

	if v.next == nil {
		u.next = nil
	} 

	u.next = v.next

	return true
}

func TransformArray(arr []float64) *LinkedList {
	var ll *LinkedList
	flag := true
	for _, key := range arr {
		if flag {
			ll = CreateLinkedList(key)
			flag = false
		} else {
			ll.InsertKey(key)
		}
	}

	return ll
}

func (ll *LinkedList) TransformLinkedList() []float64 {
	if ll == nil {
		return []float64{}
	}
	list := make([]float64, 0)

	u := ll.head

	for ; u.next != nil ; {
		list = append(list, u.chave)
		u = u.next
	}

	list = append(list, u.chave)
	return list
}