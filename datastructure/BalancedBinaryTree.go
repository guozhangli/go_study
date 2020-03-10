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

func (tree *BalancedBinaryTree) InsertNodeInBalancedBinaryTree(value int) {

}
