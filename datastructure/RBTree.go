package TestProject

const RED = true
const BLACK = false

type RBTree struct {
	Root *RBNode
}

type RBNode struct {
	Data  Map
	Left  *RBNode
	Right *RBNode
	Color bool
	Count int //这棵子树中的结点总数
}

func NewRBTree() *RBTree {
	return new(RBTree)
}

func NewRBNode(key string, value interface{}, N int, color bool) *RBNode {
	rbNode := &RBNode{
		Data:  make(Map),
		Left:  nil,
		Right: nil,
		Count: N,
		Color: color,
	}
	_map := CreateMap(key, value)
	rbNode.Data = _map
	return rbNode
}

func IsRed(node *RBNode) bool {
	if node == nil {
		return false
	}
	return node.Color == RED
}

//左旋转
func Left_Rotate(node **RBNode) {
	right := (*node).Right
	(*node).Right = right.Left
	right.Left = *node
	right.Color = (*node).Color
	(*node).Color = RED
	right.Count = (*node).Count
	(*node).Count = 1 + GetChildSize((*node).Left) + GetChildSize((*node).Right)
	*node = right
}

//右旋转
func Right_Rotate(node **RBNode) {
	left := (*node).Left
	(*node).Left = left.Right
	left.Right = *node
	left.Color = (*node).Color
	(*node).Color = RED
	left.Count = (*node).Count
	(*node).Count = 1 + GetChildSize((*node).Left) + GetChildSize((*node).Right)
	*node = left
}

func GetChildSize(node *RBNode) int {
	if node != nil {
		return node.Count
	}
	return 0
}

func FlipColors(node **RBNode) {
	(*node).Color = RED
	(*node).Left.Color = BLACK
	(*node).Right.Color = BLACK
}

func (tree *RBTree) Put(key string, value interface{}) {
	Put(&tree.Root, key, value)
	tree.Root.Color = BLACK
}

func Put(node **RBNode, key string, value interface{}) {
	if *node == nil {
		*node = NewRBNode(key, value, 1, RED)
	}
	k1 := (*node).Data.GetMapKey()
	cmp := CompareTo(key, k1.Interface())
	if cmp > 0 {
		Put(&(*node).Right, key, value)
	} else if cmp < 0 {
		Put(&(*node).Left, key, value)
	} else {
		(*node).Data = CreateMap(key, value)
	}
	if IsRed((*node).Right) && !IsRed((*node).Left) {
		Left_Rotate(node)
	}
	if IsRed((*node).Left) && IsRed((*node).Left.Left) {
		Right_Rotate(node)
	}
	if IsRed((*node).Left) && IsRed((*node).Right) {
		FlipColors(node)
	}
	(*node).Count = GetChildSize((*node).Left) + GetChildSize((*node).Right) + 1
}
