package dictionary_C

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func GenerateDictionaryTree(filename string, qmin int, qmax int, T int) *TrieTreeNode {
	tree := NewTrieTree(qmin, qmax)
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Println(err)
	}
	buff := bufio.NewReader(data)
	var sum1 int64 = 0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		str := (string)(data)
		start1 := time.Now().UnixMicro()
		for i := 0; i < len(str)-qmax; i++ {
			substring := str[i : i+qmax]
			tree.InsertIntoTrieTree(substring)
		}
		for i := len(str) - qmax; i < len(str)-qmin+1; i++ {
			substring := str[i:len(str)]
			tree.InsertIntoTrieTree(substring)
		}
		end1 := time.Now().UnixMicro()
		sum1 = (end1 - start1) + sum1
	}
	var sum2 int64 = 0
	start2 := time.Now().UnixMicro()
	tree.PruneTree(T)
	end2 := time.Now().UnixMicro()
	sum2 = (end2 - start2) + sum2
	tree.UpdateRootFrequency()
	fmt.Println("构建字典树花费时间（us）：", sum1+sum2)
	fmt.Println("构建树花费时间（us）：", sum1)
	fmt.Println("剪枝花费时间（us）：", sum2)
	//tree.PrintTree()
	return tree.root
}
