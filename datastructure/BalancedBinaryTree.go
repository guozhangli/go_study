package TestProject

type BalancedBinaryTree struct {
	Root *BalancedBinaryNode
}

type BalancedBinaryNode struct {
	Data        int
	Bf          int //平衡因子
	Left, Right *BalancedBinaryNode
}

func InitBalancedBinaryTree() *BalancedBinaryTree {
	return new(BalancedBinaryTree)
}

func checkBalancedBinaryTree(tree *BalancedBinaryTree) {
	if tree == nil {
		panic("未创建平衡二叉树")
	}
}

func (tree *BalancedBinaryTree) InsertNodeInBalancedBinaryTree(value int) {
	checkBalancedBinaryTree(tree)
	node := &BalancedBinaryNode{
		Data:  value,
		Bf:    0,
		Left:  nil,
		Right: nil,
	}
	tree.Root = node
}
