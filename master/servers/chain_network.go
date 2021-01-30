package servers

type Node struct {
	address     string
	successor   *Node
	predecessor *Node
	nodeType    NodeType
}

//Essentially a doubly linked list structure
type Chain struct {
	Head *Node
	Tail *Node
}

func NewChain(node *Node) *Chain {
	node.nodeType = HEAD_TAIL
	return &Chain{Head: node, Tail: node}
}

// func (chain *Chain) Print() {
// 	currentNode := chain.Head
// 	log.Printf("[HEAD]")
// 	for currentNode.successor != nil {
// 		log.Printf("->(%v)", currentNode.address)
// 		currentNode = currentNode.successor
// 	}
// 	log.Printf("[TAIL]")
// }
