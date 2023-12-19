package linkedlist

type Node struct {
	next *Node
	chave float64
}

type LinkedList struct {
	head *Node
}

func CreateLinkedList(key float64) *LinkedList {
	return &LinkedList{&Node{chave: key}}
}

func (l *LinkedList) InsertKey(key float64) {
	if l.head == nil {
		l.head = &Node{chave: key}
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

	u := l.head
	v := u.next
	for ; v.chave != key ; {
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