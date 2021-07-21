package binarytrees

func FindNodesDistanceK(tree *BinaryTree, target int, k int) []int {
	nodesDistanceK := []int{}
	findDistanceFromNodeToTarget(tree, target, k, &nodesDistanceK)
	return nodesDistanceK
}

func findDistanceFromNodeToTarget(node *BinaryTree, target int, k int, nodesDistanceK *[]int) int {
	if node == nil {
		return -1
	}

	if node.Value == target {
		addSubtreeNodesAtDistanceK(node, 0, k, nodesDistanceK)
		return 1
	}

	leftDistance := findDistanceFromNodeToTarget(node.Left, target, k, nodesDistanceK)
	rightDistance := findDistanceFromNodeToTarget(node.Right, target, k, nodesDistanceK)

	if leftDistance == k || rightDistance == k {
		*nodesDistanceK = append(*nodesDistanceK, node.Value)
	}

	if leftDistance != -1 {
		addSubtreeNodesAtDistanceK(node.Right, leftDistance+1, k, nodesDistanceK)
		return leftDistance + 1
	}

	if rightDistance != -1 {
		addSubtreeNodesAtDistanceK(node.Left, rightDistance+1, k, nodesDistanceK)
		return rightDistance + 1
	}

	return -1
}

func addSubtreeNodesAtDistanceK(node *BinaryTree, distance int, k int, nodesDistanceK *[]int) {
	if node == nil {
		return
	}

	if distance == k {
		*nodesDistanceK = append(*nodesDistanceK, node.Value)
	} else {
		addSubtreeNodesAtDistanceK(node.Left, distance+1, k, nodesDistanceK)
		addSubtreeNodesAtDistanceK(node.Right, distance+1, k, nodesDistanceK)
	}
}
