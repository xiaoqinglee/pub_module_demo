package pub_module_demo

import (
	"fmt"
	"github.com/xiaoqinglee/pub_module_demo/bar"
	"github.com/xiaoqinglee/pub_module_demo/bar/bar_shared_within_project"
	"github.com/xiaoqinglee/pub_module_demo/baz"
	"github.com/xiaoqinglee/pub_module_demo/baz/baz_shared_within_project"
	"github.com/xiaoqinglee/pub_module_demo/foo_shared_within_project"
	_ "go.opentelemetry.io/otel"
)

func Foo() {
	fmt.Println("Foo")
	fmt.Println(foo_shared_within_project.Pika)
}

func VisitBarFromParent() {
	bar.Bar()
	bar.VisitSibling()
	bar.VisitParent()
	fmt.Println(bar_shared_within_project.Pika)
}

func VisitBazFromParent() {
	baz.Baz()
	baz.VisitSibling()
	baz.VisitParent()
	fmt.Println(baz_shared_within_project.Pika)
}
