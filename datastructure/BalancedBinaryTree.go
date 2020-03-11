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
	newNode := &BalancedBinaryNode{
		Data:  value,
		Bf:    0,
		Left:  nil,
		Right: nil,
	}
	node := tree.Root
	if node == nil {
		tree.Root = newNode
		return
	}
	var pre *BalancedBinaryNode
	for node != nil {
		if node.Data > newNode.Data {
			if node != nil {
				pre = node
			}
			node = node.Left
		} else if node.Data < newNode.Data {
			if node != nil {
				pre = node
			}
			node = node.Right
		} else {
			pre = nil
			break
		}
	}
	if pre != nil {
		if pre.Data > newNode.Data {
			pre.Left = newNode
		} else {
			pre.Right = newNode
		}
	}
}

//右旋
func (tree *BalancedBinaryTree) R_Rotate(P **BalancedBinaryNode) {
	var L *BalancedBinaryNode
	L = (*P).Left
	(*P).Left = L.Right
	L.Right = *P
	*P = L
}

//左旋
func (tree *BalancedBinaryTree) L_Rotate(P **BalancedBinaryNode) {
	var R *BalancedBinaryNode
	R = (*P).Right
	(*P).Right = R.Left
	R.Left = *P
	*P = R
}
