package servers

import (
	"fmt"
)

type Node struct {
	debug       string
	address     string
	successor   *Node
	predecessor *Node
}

//Essentially a doubly linked list structure
type Chain struct {
	Head *Node
	Tail *Node
}

func NewChain(node *Node) *Chain {
	return &Chain{Head: node, Tail: node}
}

func (chain *Chain) Print() {
	currentNode := chain.Head
	fmt.Printf("[HEAD]")
	for currentNode.successor != nil {
		fmt.Printf("->(%v)", currentNode.debug)
		currentNode = currentNode.successor
	}
	fmt.Printf("->(%v)", currentNode.debug)
	fmt.Printf("[TAIL]\n")
}
