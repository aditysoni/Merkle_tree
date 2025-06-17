package tree

import (
	"crypto/sha256"
)

// MerkleNode represents a node in the Merkle Tree
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  []byte
}

// NewMerkleNode creates a new node (leaf or parent)
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := &MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Hash = hash[:]
	} else {
		prevHashes := append(left.Hash, right.Hash...)
		hash := sha256.Sum256(prevHashes)
		node.Hash = hash[:]
	}

	node.Left = left
	node.Right = right

	return node
}
