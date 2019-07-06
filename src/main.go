package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"qzh-go-demo/src/MFunc"
	"qzh-go-demo/src/Mchannel"
	"qzh-go-demo/src/Mgoruntine"
	"qzh-go-demo/src/Minterface"
	"qzh-go-demo/src/link"
	"qzh-go-demo/src/structTest"
)

const PI = 3.14

var (
	aa int
	bb float64
	cc bool
	// 变量要被外部使用，则需要将首个单词的首字母大写
	AA   string
	city = "hangzhou"
)

var a = "G"

func main(){
	Mchannel.TestChan()
	qzh := &Minterface.Qzh{}
	Minterface.DoDebug(qzh)

	http.HandleFunc("/hello",HelloServer)
	err := http.ListenAndServe(":8989",nil)
	if err != nil {
		fmt.Printf("http server error")
	}
}


func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w,"This is hello world")
}

func test1() {
	println(link.Link())
	map1 := make(map[string]string, 0)
	map1["qzh"] = "run"
	for i:=0; i < 10; i++{
		fmt.Println("My favorite number is", rand.Intn(100))
	}
	MFunc.Qzh("2333")

	shape := &structTest.Shape{
		Name: "qzh",
		Length: 3,
	}
	fmt.Println(shape.OutputLength())
	var arr1 [5]int
	arrTest(arr1[2:])


	var arr2 [5]int
	arr2[0] = 1
	size := 10
	testArr := make([]int, size)

	for i := 0; i < size; i ++ {
		testArr[i] = rand.Intn(100)
	}
	Mgoruntine.QuickSortQzh(testArr)
	for i,_ := range testArr {
		fmt.Printf("%d, ",testArr[i])
	}


}

func test(obj interface{}) {
	myMap := obj.(map[string]string)
	var key = "qzh233"
	val, ok := myMap[key]
	if ok {
		println(val)
	} else {
		fmt.Printf("key %v does not exist\n", key)
	}

}

func n() {
	print(a)
}

func arrTest(arr [] int){
	fmt.Println(len(arr))
}

func m() {
	// 这个  地方相当于 var a = "0" ，因此不覆盖全局变量的值
	a := "1"
	println(a)
}

func sum(a, b int64) int64 {
	var test int64 = 102
	return a + (b + test)
}
