module github.com/xiaoqinglee/pub_module_demo/bar

go 1.20

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/xiaoqinglee/pub_module_demo v0.4.1
	github.com/xiaoqinglee/pub_module_demo/baz v0.4.1
)

replace github.com/xiaoqinglee/pub_module_demo => ../

replace github.com/xiaoqinglee/pub_module_demo/baz => ../baz

retract [v0.0.0, v0.4.0]
