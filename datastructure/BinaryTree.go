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
func (tree *BinaryTree) CreatePerBinaryTree(n *BinaryTreeNode, str []string) *BinaryTreeNode {
	checkBinaryTree(tree)
	if str[tree.index] != "#" {
		n = NewBinaryTreeNode(str[tree.index])
		tree.index++
		n.Left = tree.CreatePerBinaryTree(n.Left, str)
		tree.index++
		n.Right = tree.CreatePerBinaryTree(n.Right, str)
	}
	tree.Root = n
	return n
}

func (tree *BinaryTree) CreateMidBinaryTree(n *BinaryTreeNode, str []string) *BinaryTreeNode {
	checkBinaryTree(tree)
	if str[tree.index] == "#" {
		tree.index++
		n = NewBinaryTreeNode(str[tree.index])
		//n.Left=tree.CreateMidBinaryTree(n,str)

		tree.index++
		n.Right = tree.CreateMidBinaryTree(n, str)
	}
	tree.Root = n
	return n
}
