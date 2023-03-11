package main

import (
	"github.com/xiaoqinglee/pub_module_demo"
	"github.com/xiaoqinglee/pub_module_demo/bar"
	"github.com/xiaoqinglee/pub_module_demo/baz"
)

func main() {
	pub_module_demo.Foo()
	pub_module_demo.VisitBarFromParent()
	pub_module_demo.VisitBazFromParent()
	bar.Bar()
	bar.VisitParent()
	bar.VisitSibling()
	baz.Baz()
	baz.VisitParent()
	baz.VisitSibling()
}
