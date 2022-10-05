package addonerowtotree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addRow(node *TreeNode, val, depth int) {
	if node == nil {
		return
	}

	if depth-1 == 1 {
		node.Left = &TreeNode{
			Val:  val,
			Left: node.Left,
		}

		node.Right = &TreeNode{
			Val:   val,
			Right: node.Right,
		}

		return
	}

	addRow(node.Left, val, depth-1)
	addRow(node.Right, val, depth-1)
}

func AddOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		rootNode := &TreeNode{
			Val: val,
		}

		rootNode.Left = root
		root = rootNode
	} else {
		addRow(root, val, depth)
	}

	return root
}
