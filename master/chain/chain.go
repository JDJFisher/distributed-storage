package chain

import (
	"fmt"
	"sync"
)

type Node struct {
	Debug       string
	Address     string
	Successor   *Node
	Predecessor *Node
}

//Essentially a doubly linked list structure
type Chain struct {
	Head *Node
	Tail *Node
	sync.RWMutex
}

func NewChain() *Chain {
	return &Chain{}
}

//Len - The length of the chain (the number of nodes within the chain)
func (c *Chain) Len() int {
	if c.Head == nil {
		return 0
	} else {
		currentNode := c.Head
		size := 1
		for currentNode.Successor != nil {
			currentNode = currentNode.Successor
			size++
		}
		return size
	}
}

func (chain *Chain) Print() {
	currentNode := chain.Head
	fmt.Printf("[HEAD]")
	for currentNode.Successor != nil {
		fmt.Printf("->(%v)", currentNode.Debug)
		currentNode = currentNode.Successor
	}
	fmt.Printf("->(%v)", currentNode.Debug)
	fmt.Printf("[TAIL]\n")
}

func (chain *Chain) RemoveNode(debugName string) {
	currentNode := chain.Head
	for currentNode.Successor != nil {
		if currentNode.Debug == debugName {
			//remove the node
		}
		currentNode = currentNode.Successor
	}
}
