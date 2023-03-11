module github.com/xiaoqinglee/pub_module_demo/baz

go 1.20

require (
	github.com/k0kubun/pp/v3 v3.2.0
	github.com/xiaoqinglee/pub_module_demo v0.4.1
	github.com/xiaoqinglee/pub_module_demo/bar v0.4.1
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)

replace github.com/xiaoqinglee/pub_module_demo => ../

replace github.com/xiaoqinglee/pub_module_demo/bar => ../bar

retract [v0.0.0, v0.4.0]
