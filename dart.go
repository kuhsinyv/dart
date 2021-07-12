package dart

import (
	"sort"
)

const (
	ResizeDelta     = 1 << 6
	ResizeThreshold = 0.95
	RootNodeBase    = 1
	RootNodeIndex   = 0
	EndNodeBase     = -1
)

// Dart 构造双数组实现的 Trie
type Dart struct {
	dat          *DoubleArrayTrie
	llt          *LinkedListTrie
	used         []bool
	nextCheckPos int
	keys         [][]rune
	Output       map[int][]rune
}

// resize 动态分配空间，节省资源
func (d *Dart) resize(size int) {
	d.dat.resize(size)
	d.used = append(d.used, make([]bool, size-len(d.used))...)
}

// fetch 获取子节点
func (d *Dart) fetch(parent *LinkedListTrieNode) ([]*LinkedListTrieNode, error) {
	siblings := make([]*LinkedListTrieNode, 0, 2)

	var prev rune

	for i := parent.Left; i < parent.Right; i++ {
		if len(d.keys[i]) < parent.Depth {
			continue
		}

		tmpKey := d.keys[i]

		var curr rune
		if len(d.keys[i]) != parent.Depth {
			curr = tmpKey[parent.Depth] + 1
		}

		if prev > curr {
			return nil, ErrFetch
		}

		if curr != prev || len(siblings) == 0 {
			var subKey []rune
			if curr != 0 {
				subKey = append(parent.SubKey, curr-RootNodeBase)
			} else {
				subKey = parent.SubKey
			}

			tmpNode := NewLinkedListTrieNode(
				WithCode(curr),
				WithDepth(parent.Depth+1),
				WithLeft(i),
				WithSubKey(subKey),
			)

			if len(siblings) != 0 {
				siblings[len(siblings)-1].Right = i
			}

			siblings = append(siblings, tmpNode)

			if len(parent.Children) != 0 {
				parent.Children[len(parent.Children)-1].Right = i
			}

			parent.Children = append(parent.Children, tmpNode)
		}

		prev = curr
	}

	if len(siblings) != 0 {
		siblings[len(siblings)-1].Right = parent.Right
	}

	if len(parent.Children) != 0 {
		parent.Children[len(siblings)-1].Right = parent.Right
	}

	return parent.Children, nil
}

// maxInt 返回 a，b 中的较大值
func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// setNextCheckPos 根据情况判断是否需要更新 nextCheckPos
func (d *Dart) setNextCheckPos(nonZeroNum, pos int) {
	if float64(nonZeroNum)/float64(pos-d.nextCheckPos+1) >= ResizeThreshold {
		d.nextCheckPos = pos
	}
}

// insert 构建 Double Array Trie
func (d *Dart) insert(siblings []*LinkedListTrieNode) (int, error) {
	var begin int

	var nonZeroNum int

	var first bool

	pos := maxInt(int(siblings[0].Code)+1, d.nextCheckPos) - 1

	for {
	next:
		pos++

		if len(d.dat.Base) <= pos {
			d.resize(pos + 1)
		}

		if d.dat.Check[pos] > 0 {
			nonZeroNum++

			continue
		}

		if !first {
			d.nextCheckPos = pos
			first = true
		}

		begin = pos - int(siblings[0].Code)
		if len(d.dat.Base) <= begin+int(siblings[len(siblings)-1].Code) {
			d.resize(begin + int(siblings[len(siblings)-1].Code) + ResizeDelta)
		}

		if d.used[begin] {
			continue
		}

		for i := 1; i < len(siblings); i++ {
			if d.dat.Check[begin+int(siblings[i].Code)] != 0 {
				goto next
			}
		}

		break
	}

	d.setNextCheckPos(nonZeroNum, pos)
	d.used[begin] = true

	for i := 0; i < len(siblings); i++ {
		d.dat.Check[begin+int(siblings[i].Code)] = begin
	}

	for i := 0; i < len(siblings); i++ {
		newSiblings, err := d.fetch(siblings[i])
		if err != nil {
			return -1, err
		}

		if len(newSiblings) == 0 {
			d.dat.Base[begin+int(siblings[i].Code)] = -siblings[i].Left - 1
			d.Output[begin+int(siblings[i].Code)] = siblings[i].SubKey
			siblings[i].Base = EndNodeBase
			siblings[i].Index = begin + int(siblings[i].Code)
		} else {
			n, err := d.insert(newSiblings)
			if err != nil {
				return -1, err
			}

			d.dat.Base[begin+int(siblings[i].Code)] = n
			siblings[i].Index = begin + int(siblings[i].Code)
			siblings[i].Base = n
		}
	}

	return begin, nil
}

// Build 构建 Double Array Trie 和 Linked List Trie
func (d *Dart) Build(patterns []string) (*DoubleArrayTrie, *LinkedListTrie, error) {
	if len(patterns) == 0 {
		return nil, nil, ErrEmptyPatterns
	}

	d.dat = new(DoubleArrayTrie)
	d.resize(ResizeDelta)
	sort.Strings(patterns)

	for _, pattern := range patterns {
		d.keys = append(d.keys, []rune(pattern))
	}

	d.Output = make(map[int][]rune, len(d.keys))
	d.dat.Base[0] = RootNodeBase
	d.nextCheckPos = 0
	d.llt = new(LinkedListTrie)
	d.llt.Root = NewLinkedListTrieNode(WithRight(len(patterns)), WithIndex(RootNodeIndex))

	siblings, err := d.fetch(d.llt.Root)
	if err != nil {
		return nil, nil, err
	}

	for i, node := range siblings {
		if node.Code > 0 {
			siblings[i].SubKey = append(d.llt.Root.SubKey, node.Code-RootNodeBase)
		}
	}

	_, err = d.insert(siblings)
	if err != nil {
		return nil, nil, err
	}

	return d.dat, d.llt, nil
}
