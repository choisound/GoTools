package main

import (
	"fmt"
)

type test struct {
	a int64
	b int64
	c int64
}

func testFor(queue <-chan int) {

	for i := 0; i < 10; i++ {
		fmt.Printf("%d %d %v %v\n", i, &i, i, &i)
		i := <-queue
		fmt.Printf("%d %d %v %v\n", i, &i, i, &i)

		fmt.Printf("消费了%d\n", i)
	}
}

type fast7_desc struct {
	scantime string
	scantype string
	desc     string
}
type tst struct {
	mailno           string
	status           int
	transport_status int
	fast7_desc       []fast7_desc
}
type ttt struct {
	code int
	msg  string
	data tst
}

// func main() {
// 	str := `{"code":0,"msg":"success","data":[{"mailno":"777700032006","status":0,"transport_status":0,"fast7_desc":[{"scantime":"2019-12-04 17:27:03","scantype":"\u6536\u4ef6","desc":"\u3010shenzhen branch\u3011\u7684\u3010Dai Jihui\u3011\u5df2\u6536\u4ef6"},{"scantime":"2019-12-04 17:31:10","scantype":"\u53d1\u4ef6","desc":"\u5728\u3010shenzhen branch\u3011\u626b\u63cf\u53d1\u5f80\u3010HongKong Hub\u3011"},{"scantime":"2019-12-06 16:52:48","scantype":"\u53d1\u4ef6","desc":"\u5728\u3010HongKong Hub\u3011\u626b\u63cf\u53d1\u5f80\u3010Manila Branch\u3011"},{"scantime":"2019-12-06 17:14:26","scantype":"\u7559\u4ed3\u4ef6","desc":"\u5df2\u88ab\u3010Manila Branch\u3011\u7559\u4ed3,\u539f\u56e0\u662f:\u3010COUTOMS CLEARING (\u8d27\u7269\u5230\u8fbe\u9a6c\u5c3c\u62c9\u673a\u573a-\u6e05\u5173\u4e2d)\u3011"},{"scantime":"2019-12-06 22:27:58","scantype":"\u5230\u4ef6","desc":"\u5230\u8fbe\u3010Manila Branch\u3011,\u4e0a\u4e00\u7ad9\u662f\u3010HongKong Hub\u3011"}],"transport_desc":[],"scantype":"\u5230\u4ef6"}]}`
// 	// var test ttt
// 	var sg map[string]interface{}
// 	err := json.Unmarshal([]byte(str), &sg)
// 	fmt.Printf("%+v\n", err)
// 	a, b := sg["data"].([]interface{})
// 	for _, item := range a {
// 		x, y := item.(tst)
// 		fmt.Println(item)
// 		fmt.Printf("%+v %+v\n", x, y)
// 	}
// 	fmt.Printf("%+v %+v", a, b)
// 	// st := sg["data"]
// }
func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
