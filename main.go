package main

import (
	"fmt"
	"github.com/aditysoni/merkle-tree/tree"
)

func main() {
	data := [][]byte{
		[]byte("block1"),
		[]byte("block2"),
		[]byte("block3"),
	}

	root := tree.NewMerkleTree(data)
	fmt.Printf("Merkle Root: %x\n", root.Hash)
}