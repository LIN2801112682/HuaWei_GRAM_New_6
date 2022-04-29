package main

import (
	"dictionary_C"
	"fmt"
	"index07"
	"matchQuery"
	"runtime"
)

func TraceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	fmt.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func main() {

	fmt.Println("字典树D：===============================================================")
	fmt.Println("字典树D内存占用大小：")
	//TraceMemStats()
	fmt.Println()
	root := dictionary_C.GenerateDictionaryTree("src/resources/5000Dic.txt", 2, 12, 40) //
	fmt.Println()
	//TraceMemStats()
	fmt.Println()

	fmt.Println("索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	TraceMemStats()
	fmt.Println()
	_, indexTreeNode := index07.GenerateIndexTree("src/resources/500000Index.txt", 2, 12, root) //
	fmt.Println()
	TraceMemStats()
	fmt.Println()

	//fmt.Println(unsafe.Sizeof(indexRoot))
	/*indexTreeNode.FixInvertedIndexSize()
	sort.SliceStable(index07.Res, func(i, j int) bool {
		if index07.Res[i] < index07.Res[j]  {
			return true
		}
		return false
	})
	fmt.Println(index07.Res)
	fmt.Println(len(index07.Res))
	sum := 0
	for _,val := range index07.Res{
		sum += val
	}
	fmt.Println(index07.Res[0])
	fmt.Println(index07.Res[len(index07.Res)-1])
	fmt.Println(index07.Res[len(index07.Res)/2])
	fmt.Println(sum/len(index07.Res))*/

	/*indexTreeNode.SearchGramsFromIndexTree()
	//fmt.Println(index07.Grams)
	fmt.Println(len(index07.Grams))
	var numsOfgrams2_12 [13]int
	for _,val := range index07.Grams{
		numsOfgrams2_12[len(val)]++
	}
	fmt.Println(numsOfgrams2_12)*/

	/*fmt.Println("新增索引后的索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	//TraceMemStats()
	fmt.Println()
	indexMaintain.AddIndex("src/resources/add2000.txt", 2, 6, root, indexTree)
	fmt.Println()
	//TraceMemStats()
	fmt.Println()*/

	var searchQuery1 = [10]string{"GET", "GET /english", "GET /english/images/", "GET /images/", "GET /english/images/team_hm_header_shad.gif HTTP/1.0", "GET /images/s102325.gif HTTP/1.0", "GET /english/history/history_of/images/cup/", "/images/space.gif", "GET / HTTP/1.0", "11187"}
	for i := 0; i < 10; i++ {
		resInt := matchQuery.MatchSearch(searchQuery1[i], root, indexTreeNode, 2, 12) //get english venues
		//fmt.Println(resInt)
		fmt.Println(len(resInt))
		fmt.Println("==================================================")
	}

	/*var searchQuery2 = [10]string{"french", "nav_tickets_off.gif", "ticket_quest_bg2", "HTTP/1.1", "1.0", "football.gif", "HTTP", "images", "s102438", "venue_paris_stad_header.gif"}
	for i := 0; i < 10; i++ {
		resInt := matchQuery.MatchSearch(searchQuery2[i], root, indexTreeNode, 2, 12) //get english venues
		//fmt.Println(resInt)
		fmt.Println(len(resInt))
		fmt.Println("==================================================")
	}

	var searchQuery3 = [10]string{"nav_history_off.gif", "mascot.html", "venues", "index.html", "space.gif", "GET /english/frntpage.htm HTTP/1.0", "comp_stage2_brc_topr.gif", "hm_linkf.gif", "nav_bg_bottom.jpg", "cal_paris.gif"}
	for i := 0; i < 10; i++ {
		resInt := matchQuery.MatchSearch(searchQuery3[i], root, indexTreeNode, 2, 12) //get english venues
		//fmt.Println(resInt)
		fmt.Println(len(resInt))
		fmt.Println("==================================================")
	}*/

	/*resInt := matchQuery.MatchSearch("GET / HTTP/1.0", root, indexTreeNode, 2, 12)
	fmt.Println(resInt)
	fmt.Println(len(resInt))*/

	/*map1 := make(map[index07.SeriesId][]int)
	map2 := make(map[index07.SeriesId][]int)
	var a = []int{0,1,2}
	var b = []int{0,1,6}
	var c = []int{0,1,2,3,4}
	var d = []int{0,1}
	map1[index07.NewSeriesId(1,12)] = a
	map1[index07.NewSeriesId(2,12)] = b
	map2[index07.NewSeriesId(2,12)] = c
	map2[index07.NewSeriesId(3,12)] = d
	fmt.Println(map1)
	fmt.Println(map2)
	matchQuery.MergeMapsInvertLists(map1,map2)

	fmt.Println(map1)
	fmt.Println(map2)*/

}
