package dart

// LinkedListTrieNode 链表实现的 Trie Node
type LinkedListTrieNode struct {
	// Code = 字符 UTF-8 编码值 + 1
	Code  rune
	Depth int
	// Left = 当前深度的与本字符相同的字符所在 pattern 在 patterns 中的索引起始 [Left, Right)
	Left int
	// Right = 当前深度的与本字符相同的字符所在 pattern 在 patterns 中的索引结束 [Left, Right)
	Right    int
	Index    int
	Base     int
	SubKey   []rune
	Children []*LinkedListTrieNode
}

// LinkedListTrie 链表实现的 Trie
type LinkedListTrie struct {
	Root *LinkedListTrieNode
}

// LinkedListTrieNodeOption 构造 LinkedListTrieNode 时添加可选参数的函数类型
type LinkedListTrieNodeOption func(node *LinkedListTrieNode)

// WithCode 构造 LinkedListTrieNode 时指定 Code
func WithCode(code rune) LinkedListTrieNodeOption {
	return func(node *LinkedListTrieNode) {
		node.Code = code
	}
}

// WithDepth 构造 LinkedListTrieNode 时指定 Depth
func WithDepth(depth int) LinkedListTrieNodeOption {
	return func(node *LinkedListTrieNode) {
		node.Depth = depth
	}
}

// WithLeft 构造 LinkedListTrieNode 时指定 Left
func WithLeft(left int) LinkedListTrieNodeOption {
	return func(node *LinkedListTrieNode) {
		node.Left = left
	}
}

// WithRight 构造 LinkedListTrieNode 时指定 Right
func WithRight(right int) LinkedListTrieNodeOption {
	return func(node *LinkedListTrieNode) {
		node.Right = right
	}
}

// WithIndex 构造 LinkedListTrieNode 时指定 Index
func WithIndex(index int) LinkedListTrieNodeOption {
	return func(node *LinkedListTrieNode) {
		node.Index = index
	}
}

// WithSubKey 构造 LinkedListTrieNode 时指定 SubKey
func WithSubKey(subKey []rune) LinkedListTrieNodeOption {
	return func(node *LinkedListTrieNode) {
		node.SubKey = make([]rune, len(subKey))
		copy(node.SubKey, subKey)
	}
}

// NewLinkedListTrieNode 构造 LinkedListTrieNode （参数可选）
func NewLinkedListTrieNode(opts ...LinkedListTrieNodeOption) *LinkedListTrieNode {
	node := new(LinkedListTrieNode)
	for _, opt := range opts {
		opt(node)
	}

	return node
}
