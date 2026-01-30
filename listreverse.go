package piscine

func ListReverse(l *List) {
	var prev *NodeL = nil
	current := l.Head
	l.Tail = l.Head // the old head becomes the new tail

	for current != nil {
		next := current.Next // store next
		current.Next = prev  // reverse link
		prev = current       // move prev forward
		current = next       // move current forward
	}

	l.Head = prev // prev is new head
}
