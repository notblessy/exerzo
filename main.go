package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/notblessy/pkg"
)

type JobApplication struct {
	StudentID string `json:"studentId"`
	JobID     string `json:"jobId"`
}

func main() {
	// run reverse binary tree
	// reverseBinaryTree()

	// run staircase draw
	// pkg.Staircase(int32(6))
	data, err := os.ReadFile("job_application.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Create a variable of the struct type
	var jobApp JobApplication

	// Parse the JSON data into the struct
	if err := json.Unmarshal(data, &jobApp); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Access the values
	fmt.Println("Student ID:", jobApp.StudentID)
	fmt.Println("Job ID:", jobApp.JobID)
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
