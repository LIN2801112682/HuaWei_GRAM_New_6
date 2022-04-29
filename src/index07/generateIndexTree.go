package index07

import (
	"bufio"
	"dictionary_C"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

//According to a batch of log data, the VG is divided by the dictionary_C tree, and the index item set is constructed.
func GenerateIndexTree(filename string, qmin int, qmax int, root *dictionary_C.TrieTreeNode) (*IndexTree, *IndexTreeNode) {
	indexTree := NewIndexTree(qmin, qmax)
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Print(err)
	}
	buff := bufio.NewReader(data)
	var id int32 = 0
	var sum1 int64 = 0
	var sum2 int64 = 0
	var sum3 int64 = 0
	timeStamp := time.Now().Unix()
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		var vgMap map[int]string
		vgMap = make(map[int]string)
		id++
		timeStamp++
		sid := NewSeriesId(id, timeStamp)
		str := string(data)
		start1 := time.Now().UnixMicro()
		VGConsBasicIndex(root, qmin, qmax, str, vgMap)
		var keys = []int{}
		for key := range vgMap {
			keys = append(keys, key)
		}
		//对map中的key进行排序（map遍历是无序的）保证倒排表posList顺序
		sort.Sort(sort.IntSlice(keys))
		end1 := time.Now().UnixMicro()
		sum1 += (end1 - start1)

		start2 := time.Now().UnixMicro()
		var addr *IndexTreeNode
		for i := 0; i < len(keys); i++ {
			vgKey := keys[i]
			gram := vgMap[vgKey]
			start3 := time.Now().UnixMicro()
			//addr = nil
			addr = indexTree.InsertIntoIndexTree(gram, sid, vgKey)
			end3 := time.Now().UnixMicro()
			sum3 += (end3 - start3)
			if len(gram) > qmin && len(gram) <= qmax { //Generate all index entries between qmin+1 - len(gram)
				GramSubs = make([]SubGramOffset, 0)
				GenerateQmin2QmaxGrams(gram, qmin)
				start4 := time.Now().UnixMicro()
				indexTree.InsertOnlyGramIntoIndexTree(GramSubs, addr)
				end4 := time.Now().UnixMicro()
				sum3 += (end4 - start4)
			}
		}
		end2 := time.Now().UnixMicro()
		sum2 += (end2 - start2)
	}
	indexTree.cout = (int(id))
	indexTree.UpdateIndexRootFrequency()
	fmt.Println("构建索引项集总花费时间（us）：", sum1+sum2)
	fmt.Println("读取日志并划分索引项花费时间（us）：", sum1+sum2-sum3)
	fmt.Println("插入索引树花费时间（us）：", sum3)
	//indexTree.PrintIndexTree()
	return indexTree, indexTree.root
}

var GramSubs []SubGramOffset

//func GenerateQmin2QmaxGrams(gram string, qmin int)  {
//	len := len(gram)
//	for i := len - 1; i >= qmin ; i-- {
//		for j := 1; j <= len - i; j++ {
//			gramSub := gram[j:j+i]
//			GramSubs = append(GramSubs, NewSubGramOffset(gramSub,j))
//		}
//	}
//}

func GenerateQmin2QmaxGrams(gram string, qmin int) {
	length := len(gram)
	for i := 1; i <= length-qmin; i++ {
		gramSub := gram[i:length]
		GramSubs = append(GramSubs, NewSubGramOffset(gramSub, i))
	}
}

func VGConsBasicIndex(root *dictionary_C.TrieTreeNode, qmin int, qmax int, str string, vgMap map[int]string) {
	len1 := len(str)
	for p := 0; p < len1-qmin+1; p++ {
		tSub = ""
		FindLongestGramFromDic(root, str, p)
		t := tSub
		if t == "" || len(t) < qmin { //字典D中 qmin - qmax 之间都是叶子节点（索引项中不一定是）也就是说FindLongestGramFromDic找到的只要是长度大于qmin就都是VG的gram
			t = str[p : p+qmin]
		}
		if !IsSubStrOfVG(t, vgMap) {
			vgMap[p] = t
		}
	}
}

func IsSubStrOfVG(t string, vgMap map[int]string) bool {
	var flag = false
	var keys = []int{}
	for key := range vgMap {
		keys = append(keys, key)
	}
	//对map中的key进行排序（map遍历是无序的）
	sort.Sort(sort.IntSlice(keys))
	for i := 0; i < len(keys); i++ {
		vgKey := keys[i]
		str := vgMap[vgKey]
		if str == t {
			flag = false
		} else if i == (len(keys)-1) && strings.Contains(str, t) { //vgMap中最后一个gram包含了t 不划分
			flag = true
		} else if i < (len(keys)-1) && strings.Contains(str, t) { //vgMap中不是最后一个gram包含了t 划分  觉得这里会划分出很多
			flag = false
		}
	}
	return flag
}

var tSub string

func FindLongestGramFromDic(root *dictionary_C.TrieTreeNode, str string, p int) {
	if p < len(str) {
		c := str[p : p+1]
		if root.Children()[c[0]] != nil {
			tSub += c
			FindLongestGramFromDic(root.Children()[c[0]], str, p+1)
		} else {
			return
		}
	}
}
