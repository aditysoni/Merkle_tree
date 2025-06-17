package tree
import (
	"encoding/hex"
	"fmt"
	"strings"
)
// NewMerkleTree constructs a tree from a slice of data blocks
// constructs the tree in bottoms up way 
func NewMerkleTree(data [][]byte) *MerkleNode {
	var nodes []MerkleNode

	for _, d := range data {
		node := NewMerkleNode(nil, nil, d)
		nodes = append(nodes, *node)
	}

	// Handle odd number of nodes
	if len(nodes)%2 != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
	}

	for len(nodes) > 1 {
		var newLevel []MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			left := &nodes[i]
			right := &nodes[i+1]
			parent := NewMerkleNode(left, right, nil)
			newLevel = append(newLevel, *parent)
		}

		if len(newLevel)%2 != 0 && len(newLevel) != 1 {
			newLevel = append(newLevel, newLevel[len(newLevel)-1])
		}

		nodes = newLevel
	}

	return &nodes[0]
}


func PrintMerkleTree(root *MerkleNode, level int) {
   if root == nil {
	return 
   }

   indent := strings.Repeat("  ", level)
   hash := hex.EncodeToString(root.Hash)
   fmt.Printf("%sLevel %d: Hash = %s\n", indent, level, hash)
   PrintMerkleTree(root.Left, level - 1)
   PrintMerkleTree(root.Right, level - 1)

}

func PrintDepth(root *MerkleNode) int {
	if root == nil {
		return 0 
	}
	leftDepth := PrintDepth(root.Left)
	rightDepth := PrintDepth(root.Right)

	if leftDepth > rightDepth {
		fmt.Printf("Depth %d", leftDepth + 1)
		return leftDepth + 1
	}
		fmt.Printf("Depth %d", rightDepth + 1)
		return rightDepth + 1
	
}