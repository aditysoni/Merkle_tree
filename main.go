package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aditysoni/merkle-tree/tree" // Make sure this matches your `go.mod`
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("📦 Enter number of data blocks: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numBlocks, err := strconv.Atoi(input)
	if err != nil || numBlocks <= 0 {
		fmt.Println("❌ Invalid number.")
		return
	}

	var data [][]byte
	for i := 0; i < numBlocks; i++ {
		fmt.Printf("🔹 Enter data for block %d: ", i+1)
		blockData, _ := reader.ReadString('\n')
		blockData = strings.TrimSpace(blockData)
		data = append(data, []byte(blockData))
	}

	root := tree.NewMerkleTree(data)

	fmt.Println("\n🌳 Merkle Tree Structure:")
	tree.PrintMerkleTree(root, 0)

	fmt.Printf("\n🧮 Tree Depth: %d\n", tree.PrintDepth(root))
	fmt.Printf("🔗 Merkle Root Hash: %x\n", root.Hash)
}
