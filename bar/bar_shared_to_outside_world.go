package bar

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/xiaoqinglee/pub_module_demo/baz/baz_shared_within_project"
	"github.com/xiaoqinglee/pub_module_demo/foo_shared_within_project"
)

func Bar() {
	spew.Println("Bar")
}

func VisitParent() {
	fmt.Println(foo_shared_within_project.Pika)
}

func VisitSibling() {
	fmt.Println(baz_shared_within_project.Pika)
}
