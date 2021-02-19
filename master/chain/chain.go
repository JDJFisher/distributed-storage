package chain

import (
	"fmt"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// Chain - Essentially a doubly linked list structure
type Chain struct {
	Nodes   []*Node
	Pending []uuid.UUID
	sync.RWMutex
}

// NewChain - Create a new chain object
func NewChain() *Chain {
	return &Chain{}
}

// Len - The length of the chain (the number of nodes within the chain)
func (c *Chain) Len() int {
	return len(c.Nodes)
}

// Print - Nicely pront the current state of the chain (for debug)
func (chain *Chain) Print() {
	fmt.Printf("(HEAD)")
	for _, node := range chain.Nodes {
		fmt.Printf("[%v]->", node.Address)
	}
	fmt.Printf("(TAIL)\n")

	fmt.Println("---------[ Neighbour Info ]----------")
	for _, node := range chain.Nodes {
		fmt.Printf("%v<-[%v]->%v\n", node.GetPredAddress(), node.Address, node.GetSuccAddress())
	}
	fmt.Println("-------------------")
}

// GetHead -
func (chain *Chain) GetHead() *Node {
	return chain.Nodes[0]
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
	l := len(chain.Nodes)
	if l > 0 {
		return chain.Nodes[l-1]
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
	for i, node := range chain.Nodes {
		if node.Address == address {
			return i
		}
	}
	return -1
}

// GetNode - ...
func (chain *Chain) GetNode(address string) *Node {
	for _, node := range chain.Nodes {
		if node.Address == address {
			return node
		}
	}
	return nil
}

// AddNode - ...
func (chain *Chain) AddNode(address string) (*Node, error) {
	tail := chain.GetTail()
	node := NewNode(address, tail, nil)

	if tail != nil {
		tail.SetSucc(node)
		err := tail.UpdateNeighbours(tail.GetPredAddress(), address)
		if err != nil {
			return nil, err
		}
	}

	chain.Nodes = append(chain.Nodes, node)
	return node, nil
}

// RemoveNode - ...
func (chain *Chain) RemoveNode(address string) {
	i := chain.GetNodeIndex(address)
	if i != -1 {
		chain.Nodes = append(chain.Nodes[:i], chain.Nodes[i+1:]...)
	}
}

// Fix - Attempt to fix the chain
func (chain *Chain) Fix() error {
	for i, node := range chain.Nodes {
		if i == 0 {
			node.SetPred(nil)
		} else {
			node.SetPred(chain.Nodes[i-1])
		}

		if i == chain.Len()-1 {
			node.SetSucc(nil)
		} else {
			node.SetSucc(chain.Nodes[i+1])
		}

		err := node.UpdateNeighbours(node.GetPredAddress(), node.GetSuccAddress())
		if err != nil {
			return err
		}
	}

	return nil
}

func (chain *Chain) RemoveUUIDFromPending(uid uuid.UUID) {
	for idx, val := range chain.Pending {
		if val == uid {
			chain.Pending = append(chain.Pending[:idx], chain.Pending[idx+1:]...)
			return
		}
	}
}

func (chain *Chain) IsInPending(uid uuid.UUID) bool {
	for _, val := range chain.Pending {
		if val == uid {
			return true
		}
	}
	return false
}
