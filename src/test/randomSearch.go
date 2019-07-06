package test

import "math/rand" // 引入一个标准库包

const MaxRand = 16 // 声明一个有名整型常量

// 一个函数声明
/*
 StatRandomNumbers生成一些不大于MaxRand的非负
 随机整数，并统计和返回小于和大于MaxRand/2的随机数
 个数。输入参数numRands指定了要生成的随机数的总数。
*/
func StatRandomNumbers(numRands int) (int, int) {
	// 声明了两个变量（类型都为int，初始值都为0）
	var a, b int
	// 一个for循环代码块
	for i := 0; i < numRands; i++ {
		// 一个if-else条件控制代码块
		if rand.Intn(MaxRand) < MaxRand/2 {
			a = a + 1
		} else {
			b++ // 等价于：b = b + 1
		}
	}
	return a, b // 此函数返回两个结果
}
