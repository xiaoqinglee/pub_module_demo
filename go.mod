module github.com/xiaoqinglee/pub_module_demo

go 1.20

require (
	github.com/xiaoqinglee/pub_module_demo/bar v0.4.1
	github.com/xiaoqinglee/pub_module_demo/baz v0.4.1
	go.opentelemetry.io/otel v1.14.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/k0kubun/pp/v3 v3.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)

replace github.com/xiaoqinglee/pub_module_demo/bar => ./bar

replace github.com/xiaoqinglee/pub_module_demo/baz => ./baz

retract [v0.0.0, v0.4.0]
