package main

import "fmt"

type Node struct {
	Next  *Node
	Value int
}

func (node *Node) Count() (count int) {
	count = 0

	for node != nil {
		count++
		node = node.Next
	}

	return count
}

type LinkedList struct {
	Root   *Node
	Length int
}

// Reference https://stackoverflow.com/a/37135458/8094831
func NewLinkedList() LinkedList {
	list := LinkedList{Root: nil, Length: 0}
	return list
}

// Reference http://bulkan-evcimen.com/writing_linked_list_using_golang_part_one.html
func (list *LinkedList) Add(data int) {
	if list.Root == nil {
		list.Root = &Node{Value: data, Next: nil}
	} else {
		tmp := &Node{Value: data, Next: list.Root}
		list.Root = tmp
	}
	list.Length++
}

func (list *LinkedList) AddNode(node Node) {
	list.Length += node.Count()

	if list.Root == nil {
		list.Root = &node
	} else {
		tmp := &node

		for tmp != nil {
			if tmp.Next == nil {
				tmp.Next = list.Root
				break
			} else {
				tmp = tmp.Next
			}
		}

		list.Root = &node
	}
}

func (list *LinkedList) Print() {
	node := list.Root
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}

func (list *LinkedList) Count() (count int) {
	node := list.Root
	count = 0

	for node != nil {
		count++
		node = node.Next
	}

	return count
}

func printStats(list LinkedList) {
	fmt.Printf("\nLength of list: %v\nCount of list: %v\n", list.Length, list.Count())
}

func main() {
	var node Node = Node{nil, 5}
	fmt.Println(node.Value)

	anotherNode := Node{&Node{nil, 33}, 10}
	fmt.Println(anotherNode)

	node.Next = &anotherNode

	var list LinkedList = NewLinkedList()

	list.Add(14)
	list.Add(29)
	list.Add(55)

	printStats(list)
	list.AddNode(node)
	printStats(list)

	fmt.Println("\nPrint list of nodes")
	list.Print()

	fmt.Println()
	fmt.Println(list.Root)
	fmt.Println(list.Root.Next)
}
