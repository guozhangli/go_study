package TestProject

type BinaryTreeNode struct {
	Data  interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root  *BinaryTreeNode
	index int
}

func NewBinaryTreeNode(vaule interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data:  vaule,
		Left:  nil,
		Right: nil,
	}
}

func checkBinaryTree(tree *BinaryTree) {
	if tree == nil {
		panic("no binaryTree created")
	}
}
func (tree *BinaryTree) CreateBinaryTree(n *BinaryTreeNode, str []string) *BinaryTreeNode {
	checkBinaryTree(tree)
	if str[tree.index] != "#" {
		n = NewBinaryTreeNode(str[tree.index])
		tree.index++
		n.Left = tree.CreateBinaryTree(n.Left, str)
		tree.index++
		n.Right = tree.CreateBinaryTree(n.Right, str)
	}
	tree.Root = n
	return n
}
