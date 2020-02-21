package main

import "fmt"

//TestSlice test
func TestSlice() {
	arr := [2]int{1, 2}
	arr1 := [...]int{1, 2}
	fmt.Printf("%+v\n", arr == arr1)
	arr2 := make([]int, 2, 2)
	arr2[0] = 1
	arr2[1] = 2
	sarr := arr[:]
	sarr1 := append(sarr, 3)
	sarr[0] = 0
	fmt.Printf("sarr1 %+v %d %d\n", sarr1, len(sarr1), cap(sarr1))
	sarr2 := sarr1[:2]
	sarr2 = append(sarr2, 9)
	sarr2 = append(sarr2, 4)
	sarr2 = append(sarr2, 0)
	fmt.Printf("sarr1 %+v %d %d\n sarr2  %+v %d %d\n", sarr1, len(sarr1), cap(sarr1), sarr2, len(sarr2), cap(sarr2))
	sarr1[2] = 7
	fmt.Printf("sarr1 %+v %d %d\n sarr2  %+v %d %d\n", sarr1, len(sarr1), cap(sarr1), sarr2, len(sarr2), cap(sarr2))
}
