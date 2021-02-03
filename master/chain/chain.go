package chain

import (
	"fmt"
	"sync"
)

// Chain - Essentially a doubly linked list structure
type Chain struct {
	Head *Node
	Tail *Node
	sync.RWMutex
}

// NewChain - Create a new chain object
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
		for currentNode.successor != nil {
			currentNode = currentNode.successor
			size++
		}
		return size
	}
}

// Print - Nicely pront the current state of the chain (for debug)
func (chain *Chain) Print() {
	currentNode := chain.Head
	fmt.Printf("[HEAD]")
	for currentNode.successor != nil {
		fmt.Printf("->(%v)", currentNode.Address)
		currentNode = currentNode.successor
	}
	fmt.Printf("->(%v)", currentNode.Address)
	fmt.Printf("[TAIL]\n")
}

// GetNode - Get a node object for the given node address
func (chain *Chain) GetNode(address string) *Node {
	node := chain.Head
	for {
		if node == nil {
			return nil
		} else if node.Address == address {
			return node
		} else {
			node = node.successor
		}
	}
}

// RemoveNode - Delete a node from the chain (its removed through garbage collection following the pointers being updated of its neighbours)
func (chain *Chain) RemoveNode(node *Node) {
	if node.predecessor != nil {
		node.predecessor.successor = node.successor
	} else if node == chain.Head {
		chain.Head = node.successor
	}

	if node.successor != nil {
		node.successor.predecessor = node.predecessor
	} else if node == chain.Tail {
		chain.Tail = node.predecessor
	}
}
