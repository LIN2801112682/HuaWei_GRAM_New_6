package index07

import "reflect"

type IndexTree struct {
	qmin int
	qmax int
	cout int
	root *IndexTreeNode
}

func (i *IndexTree) Qmin() int {
	return i.qmin
}

func (i *IndexTree) SetQmin(qmin int) {
	i.qmin = qmin
}

func (i *IndexTree) Qmax() int {
	return i.qmax
}

func (i *IndexTree) SetQmax(qmax int) {
	i.qmax = qmax
}

func (i *IndexTree) Cout() int {
	return i.cout
}

func (i *IndexTree) SetCout(cout int) {
	i.cout = cout
}

func (i *IndexTree) Root() *IndexTreeNode {
	return i.root
}

func (i *IndexTree) SetRoot(root *IndexTreeNode) {
	i.root = root
}

func NewIndexTree(qmin int, qmax int) *IndexTree {
	return &IndexTree{
		qmin: qmin,
		qmax: qmax,
		cout: 0,
		root: NewIndexTreeNode(""),
	}
}

//Insert gram into IndexTree  position:The starting position of the strat in the statement
func (tree *IndexTree) InsertIntoIndexTree(gram string, sid SeriesId, position int) *IndexTreeNode {
	node := tree.root
	var addr *IndexTreeNode
	var childIndex int8 = -1
	for i := 0; i < len(gram); i++ {
		childIndex = GetIndexNode(node.children, gram[i])
		if childIndex == -1 {
			currentNode := NewIndexTreeNode(string(gram[i]))
			node.children[gram[i]] = currentNode
			node = currentNode
		} else {
			node = node.children[uint8(childIndex)]
			node.frequency++
		}
		if i == len(gram)-1 { //Leaf node, need to hook up linkedList
			node.isleaf = true
			if _, ok := node.invertedIndex[sid]; !ok { //There is no sid in the key, to create an inverted structure corresponding to the sid
				node.InsertSidAndPosArrToInvertedIndexMap(sid, position)
			} else { //Find the same sid and add posArray[j]
				node.InsertPosArrToInvertedIndexMap(sid, position)
			}
			addr = node
		}
	}
	return addr
}

func (tree *IndexTree) InsertOnlyGramIntoIndexTree(gramSubs []SubGramOffset, addr *IndexTreeNode) {
	var childIndex int8 = -1
	for k := 0; k < len(gramSubs); k++ {
		gram := gramSubs[k].subGram
		offset := gramSubs[k].offset
		node := tree.root
		for i := 0; i < len(gram); i++ {
			childIndex = GetIndexNode(node.children, gram[i])
			if childIndex == -1 {
				currentNode := NewIndexTreeNode(string(gram[i]))
				node.children[gram[i]] = currentNode
				node = currentNode
			} else {
				node = node.children[uint8(childIndex)]
				node.frequency++
			}
			if i == len(gram)-1 { //Leaf node, need to hook up linkedList
				node.isleaf = true
				if _, ok := node.addrOffset[addr]; !ok {
					node.addrOffset[addr] = offset
				}
			}
		}
	}
}

func (tree *IndexTree) PrintIndexTree() {
	tree.root.PrintIndexTreeNode(0)
}

func (tree *IndexTree) UpdateIndexRootFrequency() {
	for _, child := range tree.root.children {
		tree.root.frequency += child.frequency
	}
	tree.root.frequency--
}

//Calculate the length of each invertedList
var Res []int

func (root *IndexTreeNode) FixInvertedIndexSize() {
	for _, child := range root.children {
		if child.isleaf == true {
			Res = append(Res, len(child.invertedIndex)) //The append function must be used, and i cannot be used for variable addition, because there is no make initialization
		}
		child.FixInvertedIndexSize()
	}
}

//Calculate the length of each gram
var Grams []string
var temp string

func (root *IndexTreeNode) SearchGramsFromIndexTree() {
	if root == nil {
		return
	}
	for _, child := range root.children {
		if child != nil {
			temp += child.data
			if child.isleaf == true {
				Grams = append(Grams, temp)
			}
			child.SearchGramsFromIndexTree()
			temp = temp[0 : len(temp)-1]
		}
	}
}

//remove the same index entry
func RemoveSliceInvertIndex(grams []string) (ret []string) {
	n := len(grams)
	for i := 0; i < n; i++ {
		state := false
		for j := i + 1; j < n; j++ {
			if j > 0 && reflect.DeepEqual(grams[i], grams[j]) {
				state = true
				break
			}
		}
		if !state {
			ret = append(ret, grams[i])
		}
	}
	return
}
