package chain

import (
	"fmt"
	"sync"
)

type Node struct {
	Address     string
	Successor   *Node
	Predecessor *Node
}

// Chain - Essentially a doubly linked list structure
type Chain struct {
	Head *Node
	Tail *Node
	sync.RWMutex
}

func NewChain() *Chain {
	return &Chain{}
}

// Len - The length of the chain (the number of nodes within the chain)
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

// Print ...
func (chain *Chain) Print() {
	currentNode := chain.Head
	fmt.Printf("[HEAD]")
	for currentNode.Successor != nil {
		fmt.Printf("->(%v)", currentNode.Address)
		currentNode = currentNode.Successor
	}
	fmt.Printf("->(%v)", currentNode.Address)
	fmt.Printf("[TAIL]\n")
}

// RemoveNode ...
func (chain *Chain) RemoveNode(address string) {
	currentNode := chain.Head
	for currentNode.Successor != nil {
		if currentNode.Address != address {
			currentNode = currentNode.Successor
		}
		currentNode.Predecessor.Successor = currentNode.Successor
		currentNode.Successor.Predecessor = currentNode.Predecessor
		return
	}
}
