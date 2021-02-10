package chain

import (
	"fmt"
	"sync"
)

// Chain - Essentially a doubly linked list structure
type Chain struct {
	nodes []*Node
	sync.RWMutex
}

// NewChain - Create a new chain object
func NewChain() *Chain {
	return &Chain{}
}

// Len - The length of the chain (the number of nodes within the chain)
func (c *Chain) Len() int {
	return len(c.nodes)
}

// Print - Nicely pront the current state of the chain (for debug)
func (chain *Chain) Print() {
	fmt.Printf("(HEAD)")
	for _, node := range chain.nodes {
		fmt.Printf("[%v]->", node.Address)
	}
	fmt.Printf("(TAIL)\n")

	fmt.Println("---------[ Neighbour Info ]----------")
	for _, node := range chain.nodes {
		fmt.Printf("%v<-[%v]->%v\n", node.GetPredAddress(), node.Address, node.GetSuccAddress())
	}
	fmt.Println("-------------------")
}

// GetHead -
func (chain *Chain) GetHead() *Node {
	return chain.nodes[0]
}

// GetHeadAddress -
func (chain *Chain) GetHeadAddress() string {
	head := chain.GetHead()
	if head != nil {
		return head.Address
	}
	return ""
}

// GetTail -
func (chain *Chain) GetTail() *Node {
	l := len(chain.nodes)
	if l > 0 {
		return chain.nodes[l-1]
	}
	return nil
}

// GetTailAddress -
func (chain *Chain) GetTailAddress() string {
	tail := chain.GetTail()
	if tail != nil {
		return tail.Address
	}
	return ""
}

// GetNodeIndex - ...
func (chain *Chain) GetNodeIndex(address string) int {
	for i, node := range chain.nodes {
		if node.Address == address {
			return i
		}
	}
	return -1
}

// GetNode - ...
func (chain *Chain) GetNode(address string) *Node {
	for _, node := range chain.nodes {
		if node.Address == address {
			return node
		}
	}
	return nil
}

// AddNode - ...
func (chain *Chain) AddNode(address string) *Node {
	tail := chain.GetTail()
	node := NewNode(address, tail, nil)

	if tail != nil {
		tail.SetSucc(node)
	}

	chain.nodes = append(chain.nodes, node)
	return node
}

// RemoveNode - ...
func (chain *Chain) RemoveNode(address string) {
	i := chain.GetNodeIndex(address)
	if i != -1 {
		chain.nodes = append(chain.nodes[:i], chain.nodes[i+1:]...)
	}
}

// Fix - Attempt to fix the chain
func (chain *Chain) Fix() error {
	for i, node := range chain.nodes {
		if i == 0 {
			node.SetPred(nil)
		} else {
			node.SetPred(chain.nodes[i-1])
		}

		if i == chain.Len()-1 {
			node.SetSucc(nil)
		} else {
			node.SetSucc(chain.nodes[i+1])
		}

		err := node.UpdateNeighbours(node.GetPredAddress(), node.GetSuccAddress())
		if err != nil {
			return err
		}
	}

	return nil
}
