package TestProject

const (
	LH = 1
	RH = -1
)

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

func NewBalanceBinaryNode(value int) *BalancedBinaryNode {
	return &BalancedBinaryNode{
		Data:  value,
		Bf:    0,
		Left:  nil,
		Right: nil,
	}
}

func (tree *BalancedBinaryTree) InsertNodeInBalancedBinaryTree(node **BalancedBinaryNode, value int) {
	checkBalancedBinaryTree(tree)
	newNode := NewBalanceBinaryNode(value)
	if *node == nil {
		*node = newNode
	} else {
		if (*node).Data > value {
			tree.InsertNodeInBalancedBinaryTree(&(*node).Left, value)
			(*node).Bf = GetHight((*node).Left) - GetHight((*node).Right)
			if (*node).Bf > LH {
				if (*node).Left.Bf == LH {
					R_Rotate(node)
				} else {
					LR_Rotate(node)
				}
			}
		} else if (*node).Data < value {
			tree.InsertNodeInBalancedBinaryTree(&(*node).Right, value)
			(*node).Bf = GetHight((*node).Left) - GetHight((*node).Right)
			if (*node).Bf < RH {
				if (*node).Right.Bf == RH {
					L_Rotate(node)
				} else {
					RL_Rotate(node)
				}
			}
		} else {
			return
		}
	}
}

func (tree *BalancedBinaryTree) DeleteNodeInBalancedBinaryTree(element int) {
	checkBalancedBinaryTree(tree)

}

func (tree *BalancedBinaryTree) SearchMaxElementInBalancedBinaryTree() *BalancedBinaryNode {
	checkBalancedBinaryTree(tree)
	node := tree.Root
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

//递归求最大节点
func (tree *BalancedBinaryTree) SearchMaxElementInBalancedBinaryTree_1(node *BalancedBinaryNode) *BalancedBinaryNode {
	checkBalancedBinaryTree(tree)
	if node == nil {
		return nil
	}
	if node.Right == nil {
		return node
	}
	return tree.SearchMaxElementInBalancedBinaryTree_1(node.Right)
}

func (tree *BalancedBinaryTree) SearchMinElementInBalancedBinaryTree() *BalancedBinaryNode {
	checkBalancedBinaryTree(tree)
	node := tree.Root
	if node == nil {
		return nil
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
}

//递归求最小节点
func (tree *BalancedBinaryTree) SearchMinElementInBalancedBinaryTree_1(node *BalancedBinaryNode) *BalancedBinaryNode {
	checkBalancedBinaryTree(tree)
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node
	}
	return tree.SearchMinElementInBalancedBinaryTree_1(node.Left)
}

//单旋转（右旋）
//左左插入导致的不平衡
func R_Rotate(P **BalancedBinaryNode) {
	var L *BalancedBinaryNode
	L = (*P).Left
	(*P).Left = L.Right
	L.Right = *P
	(*P).Bf = GetHight((*P).Left) - GetHight((*P).Right)
	L.Bf = GetHight(L.Left) - GetHight(L.Right)
	*P = L
}

//单旋转（左旋）
//右右插入导致的不平衡
func L_Rotate(P **BalancedBinaryNode) {
	var R *BalancedBinaryNode
	R = (*P).Right
	(*P).Right = R.Left
	R.Left = *P
	(*P).Bf = GetHight((*P).Left) - GetHight((*P).Right)
	R.Bf = GetHight(R.Left) - GetHight(R.Right)
	*P = R
}

//双旋转
//插入点位于P的左儿子的右子树
func LR_Rotate(P **BalancedBinaryNode) {
	//双旋转可以通过两次单旋转实现
	// 对P的左结点进行R旋转，再对根节点进行L旋转
	L_Rotate(&(*P).Left)
	R_Rotate(P)
}

//双旋转
//插入点位于P的右儿子的左子树
func RL_Rotate(P **BalancedBinaryNode) {
	//双旋转可以通过两次单旋转实现
	// 对P的右结点进行L旋转，再对根节点进行R旋转
	R_Rotate(&(*P).Right)
	L_Rotate(P)
}

func GetHight(root *BalancedBinaryNode) int {
	if root == nil {
		return 0
	} else {
		queue := NewQueueLinked()
		queue.EnQueue(root)
		queue.EnQueue(nil)
		var level int
		for !queue.IsEmpty() {
			var node *BalancedBinaryNode
			temp := queue.DeQueue().(*Node).Data
			if temp != nil {
				node = temp.(*BalancedBinaryNode)
			}
			if node == nil {
				if !queue.IsEmpty() {
					queue.EnQueue(nil)
				}
				level++
			} else {
				if node.Left != nil {
					queue.EnQueue(node.Left)
				}
				if node.Right != nil {
					queue.EnQueue(node.Right)
				}
			}
		}
		return level
	}
}
