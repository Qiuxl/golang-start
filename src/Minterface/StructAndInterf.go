package Minterface

import (
	"fmt"
	"qzh-go-demo/src/MFunc"
)
type List []int

func (l List) Len() int {
	fmt.Printf("%v length is %v\n",l,len(l))
	return len(l)
}
func (l *List) Append(val int) {
	fmt.Printf("I'm appending! for value %v\n", val)
	*l = append(*l, val)
}
type Appender interface {
	Append(int)
}
func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

/**
 接口跨包实现，同名不冲突
 */
func LongEnough(l MFunc.Lener) bool {
	return l.Len()*10 > 42
}

func LongEnough2(l MFunc.Lener2) bool {
	return l.Len()*10 > 42
}

func Test(){
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	// List does not implement Appender (Append method has pointer receiver)
	// CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}
	(&lst).Append(100)
	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough2(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
	(*plst).Len()
}
