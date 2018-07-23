package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// cache := make(map[string]string)
	// cache["name"] = "azen"
	// fmt.Println(cache)

	//	-- array部分 --
	// fmt.Println("array部分开始啦啦啦...")
	// arrayInit()
	// gridArray()
	// loop()
	// rangeLoop()

	//	-- slice部分 --
	// fmt.Println("slice部分开始啦啦啦...")
	// sliceCreate()
	// updateSlice()
	// reslice()
	// sliceExpand()
	// sliceAppend()
	// sliceCreateOps()
	// zeroSliceAppend()
	// initSlice()
	// sliceCopy()
	// sliceDelete()

	//	-- map部分 --
	// mapCreate()
	// mapEmptyCreate()
	// mapLoop()
	// mapNoValue()
	// mapNoValueProtect()
	// deleteKeyForMap()

	//	-- rune部分 --
	lengthTest()
	wrongIndexForString()
	runeSliceIndex()
	getRuneCount()
}

func getRuneCount() {
	fmt.Println("我是Miku~", utf8.RuneCountInString("我是Miku~"))
	// 我是Miku~ 7
}

func runeSliceIndex() {
	content := "我是Miku~"
	for i, r := range []rune(content) {
		fmt.Printf("\n(index=%d rune=%c char=%X)", i, r, r)
	}

	/*
		(index=0 rune=我 char=6211)
		(index=1 rune=是 char=662F)
		(index=2 rune=M char=4D)
		(index=3 rune=i char=69)
		(index=4 rune=k char=6B)
		(index=5 rune=u char=75)
		(index=6 rune=~ char=7E)
	*/
}

func wrongIndexForString() {
	content := "我是Miku~"
	for i, ch := range content {
		fmt.Printf("\n(index=%d content=%X)", i, ch)
	}
	/*
		(index=0 content=6211)
		(index=3 content=662F)
		(index=6 content=4D)
		(index=7 content=69)
		(index=8 content=6B)
		(index=9 content=75)
		(index=10 content=7E)
	*/
}

func lengthTest() {
	content := "我是Miku~"
	println(len(content)) // 11

	for _, b := range []byte(content) {
		fmt.Printf("%X ", b)
	}
	//	E6 88 91 E6 98 AF 4D 69 6B 75 7E
}

/**
* map部分
 */

func deleteKeyForMap() {
	mikuMap := map[string]string{
		"name":     "Miku",
		"type":     "VOC",
		"age":      "16",
		"birthday": "8月31日",
	}
	delete(mikuMap, "type")
	fmt.Println(mikuMap)
}

func mapNoValueProtect() {
	mikuMap := map[string]string{
		"name":     "Miku",
		"type":     "VOC",
		"age":      "16",
		"birthday": "8月31日",
	}
	if color, ok := mikuMap["color"]; ok {
		fmt.Println(color)
	} else {
		fmt.Println("ㄟ...不知道miku的色值...")
	}
}

func mapNoValue() {
	mikuMap := map[string]string{
		"name":     "Miku",
		"type":     "VOC",
		"age":      "16",
		"birthday": "8月31日",
	}
	fmt.Println("miku的色值为：", mikuMap["color"])
}

func mapLoop() {
	mikuMap := map[string]string{
		"name":     "Miku",
		"type":     "VOC",
		"age":      "16",
		"birthday": "8月31日",
	}
	for k, v := range mikuMap {
		fmt.Println(k, v)
	}
}

func mapEmptyCreate() {
	var mikuMap map[string]string
	fmt.Println(mikuMap) // map[]
	mikuMapAnother := make(map[string]string)
	fmt.Println(mikuMapAnother) // map[]
}

func mapCreate() {
	mikuMap := map[string]string{
		"name":     "Miku",
		"type":     "VOC",
		"age":      "16",
		"birthday": "8月31日",
	}
	fmt.Println(mikuMap)
}

/**
* slice部分
 */

func sliceDelete() {
	s := []int{0, 1, 2}
	s = append(s[:1], s[2:]...)
	fmt.Println(s)
}

func sliceCopy() {
	s1 := []int{88, 22}
	copy(s1[0:], s1[1:])
	fmt.Println(s1)
}

func initSlice() {
	s := []int{686}
	fmt.Println(s)
}

func zeroSliceAppend() {
	var s []int
	s2 := append(s, 888)
	fmt.Println(s2)
}

func sliceCreateOps() {
	var s1 []int
	fmt.Println(s1) // []
}

func sliceAppend() {
	arr := [...]int{0, 1, 2}
	fmt.Println(arr)
	s1 := arr[:1]
	s2 := s1[:2]
	s3 := append(s2, 438)
	s4 := append(s3, 10086)
	s5 := append(s4, 233)
	fmt.Println(arr, s1, s2, s3, s4, s5)
	s4[0] = 999
	fmt.Println(arr, s1, s2, s3, s4, s5)
}

func sliceExpand() {
	arr := [...]int{0, 1, 2}
	s1 := arr[:1] // 0
	s2 := s1[1:3] // 1,2
	fmt.Println(s1, s2)
}

func reslice() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr)

	s := arr[:]

	func(s []int) {
		s[0] = 100
	}(s)
	fmt.Println(s, arr)

	func(s []int) {
		s[0] = 100
	}(s[2:])
	fmt.Println(s, arr) // [100 1 100 3 4 5] [100 1 100 3 4 5]
}

func updateSlice() {
	arr := [...]int{0, 1, 2, 3}
	fmt.Println(arr)
	func(s []int) {
		s[0] = 100
	}(arr[:])
	fmt.Println(arr) // 100,1,2,3
}

func sliceCreate() {
	arr := [...]int{0x66CCFF, 666, 233}
	s := arr[1:2]
	fmt.Println(s)
	fmt.Println(arr[:2])
	fmt.Println(arr[1:])
	fmt.Println(arr[:])
}

/**
* array部分
 */
func rangeLoop() {
	array := [...]int{666, 888, 0x66CCFF}
	for i := range array {
		fmt.Println(i)
	}
}

func loop() {
	array := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func gridArray() {
	arr4 := [3][2]int{{1, 2}}
	fmt.Println(arr4)
}

func arrayInit() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{12, 34, 56, 7}

	fmt.Println(arr1, arr2, arr3)
}
