package dart

// DoubleArrayTrie 双数组实现的 Trie
// https://linux.thai.net/~thep/datrie/
type DoubleArrayTrie struct {
	Base  []int
	Check []int
}

// resize 动态分配空间，节省资源
func (dat *DoubleArrayTrie) resize(size int) {
	dat.Base = append(dat.Base, make([]int, size-len(dat.Base))...)
	dat.Check = append(dat.Check, make([]int, size-len(dat.Check))...)
}

// ExactMatchSearch 精准匹配搜索
func (dat *DoubleArrayTrie) ExactMatchSearch(content []rune, pos int) bool {
	b := dat.Base[pos]

	var p int
	for _, c := range content {
		p = b + int(c) + 1
		if p >= len(dat.Check) {
			return false
		}

		if dat.Check[p] == b {
			b = dat.Base[p]
		} else {
			return false
		}
	}

	p = b
	n := dat.Base[p]

	if b == dat.Check[p] && n < 0 {
		return true
	}

	return false
}
