package Minterface

import "fmt"

type Persion interface {
	Name() string
}

type Coder interface {
	Persion
	Debug()
}

type Qzh struct {
}

func (*Qzh) Name() string {
	return "I'm qzh"
}

func (*Qzh) Debug()  {
	fmt.Println("I'm debugging!")
}

func DoDebug(c Coder)  {
	c.Debug()
}