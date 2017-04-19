package handler

import "sort"

// InitTree sets all parents in tree
func InitTree(tree PathPart) {
	initTreeRecursive(tree, nil)
}

func initTreeRecursive(tree, parent PathPart) {
	tree.setParent(parent)
	sort.Sort(sorter(tree.Children()))
	for _, child := range tree.Children() {
		initTreeRecursive(child, tree)
	}
}

type sorter []PathPart

func (s sorter) Len() int {
	return len(s)
}

func (s sorter) Less(i, j int) bool {
	if len(s[i].Children()) == 0 && len(s[j].Children()) > 0 {
		return false
	}
	return len(s[i].Name()) < len(s[j].Name())
}

func (s sorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
