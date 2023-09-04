package main

import (
	"fmt"

	"github.com/notblessy/pkg"
)

func main() {
	// run reverse binary tree
	reverseBinaryTree()

	// run staircase draw
	pkg.Staircase(int32(6))
}

func reverseBinaryTree() {
	// Example binary tree
	root := &pkg.TreeNode{
		Val: 1,
		Left: &pkg.TreeNode{
			Val: 2,
			Left: &pkg.TreeNode{
				Val: 4,
			},
			Right: &pkg.TreeNode{
				Val: 5,
			},
		},
		Right: &pkg.TreeNode{
			Val: 3,
		},
	}

	fmt.Println("Original binary tree pretty:")
	pkg.PrintTreeASCII(root)
	fmt.Println("Original binary tree inline:")
	pkg.PrintTree(root)

	// Reverse the binary tree
	reversedRoot := pkg.ReverseBinaryTree(root)

	fmt.Println("\nReversed binary tree pretty:")
	pkg.PrintTreeASCII(reversedRoot)

	fmt.Println("\nReversed binary tree inline:")
	pkg.PrintTree(reversedRoot)
}
