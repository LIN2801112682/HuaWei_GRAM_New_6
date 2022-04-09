package indexMaintain

import (
	"bufio"
	"dictionary"
	"fmt"
	"index07"
	"io"
	"os"
	"sort"
	"time"
)

//根据一批日志数据通过字典树划分VG，增加到索引项集中
func AddIndex(filename string, qmin int, qmax int, root *dictionary.TrieTreeNode, indexTree *index07.IndexTree) *index07.IndexTree {
	start := time.Now().UnixMicro()
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Print(err)
	}
	buff := bufio.NewReader(data)
	id := indexTree.Cout()
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		var vgMap map[int]string
		vgMap = make(map[int]string)
		id++
		timeStamp := time.Now().Unix()
		sid := index07.NewSeriesId(int32(id), timeStamp)
		str := string(data)
		index07.VGConsBasicIndex(root, qmin, qmax, str, vgMap)
		var keys = []int{}
		for key := range vgMap {
			keys = append(keys, key)
		}
		//对map中的key进行排序（map遍历是无序的）
		sort.Sort(sort.IntSlice(keys))
		var addr *index07.IndexTreeNode
		for i := 0; i < len(keys); i++ {
			vgKey := keys[i]
			gram := vgMap[vgKey]
			addr = indexTree.InsertIntoIndexTree(gram, sid, vgKey)
			if len(gram) > qmin && len(gram) <= qmax { //Generate all index entries between qmin+1 - len(gram)
				index07.GramSubs = make([]index07.SubGramOffset, 0)
				index07.GenerateQmin2QmaxGrams(gram, qmin)
				indexTree.InsertOnlyGramIntoIndexTree(index07.GramSubs, addr)
			}
		}
	}
	indexTree.SetCout(id)
	indexTree.Root().SetFrequency(1)
	indexTree.UpdateIndexRootFrequency()
	end := time.Now().UnixMicro()
	fmt.Println("新增索引项集花费时间（us）：", end-start)
	//indexTree.PrintIndexTree()
	return indexTree
}
