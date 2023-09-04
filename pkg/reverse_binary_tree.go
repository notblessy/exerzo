package pkg

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ReverseBinaryTree reverses the binary tree by swapping left and right subtrees.
func ReverseBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap the left and right subtrees using a temporary variable.
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	// Recursively reverse the left and right subtrees.
	ReverseBinaryTree(root.Left)
	ReverseBinaryTree(root.Right)

	return root
}

// PrintTree prints the binary tree in a readable format (in-order traversal).
func PrintTree(root *TreeNode) {
	if root != nil {
		PrintTree(root.Left)
		fmt.Print(root.Val, " ")
		PrintTree(root.Right)
	}
}

// PrintTreeASCII prints the binary tree in ASCII format (level-order traversal).
func PrintTreeASCII(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	fmt.Println("QUEUE", &queue)

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				fmt.Print(node.Val, " ")

				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}

		fmt.Println()
	}
}

// // Example binary tree
// root := &TreeNode{
// 	Val: 1,
// 	Left: &TreeNode{
// 		Val: 2,
// 		Left: &TreeNode{
// 			Val: 4,
// 		},
// 		Right: &TreeNode{
// 			Val: 5,
// 		},
// 	},
// 	Right: &TreeNode{
// 		Val: 3,
// 	},
// }

// fmt.Println("Original binary tree:")
// printTreeASCII(root)

// // Reverse the binary tree
// reversedRoot := reverseBinaryTree(root)

// fmt.Println("\nReversed binary tree:")
// printTreeASCII(reversedRoot)
