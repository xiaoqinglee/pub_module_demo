package baz

import (
	"fmt"
	"github.com/k0kubun/pp/v3"
	"github.com/xiaoqinglee/pub_module_demo/bar/bar_shared_within_project"
	"github.com/xiaoqinglee/pub_module_demo/foo_shared_within_project"
)

func Baz() {
	pp.Println("Baz")
}

func VisitParent() {
	fmt.Println(foo_shared_within_project.Pika)
}

func VisitSibling() {
	fmt.Println(bar_shared_within_project.Pika)
}
