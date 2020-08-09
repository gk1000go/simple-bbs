package main

import (
	"fmt"
	"sync"
)

var m sync.Map

type data struct {
	d int
}

func main()  {
	m.Store("abc",&data{123})
	d,_:=m.Load("abc")
	fmt.Println(d)
	d.(*data).d=456
	e,_:=m.Load("abc")
	fmt.Println(e)
}