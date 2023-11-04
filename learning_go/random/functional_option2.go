package main

import "fmt"

const (
	DefaultRegistry      = "index.docker.io"
	defaultRegistryAlias = "docker.io"
	DefaultTag           = "latest"
)

type options struct {
	strict          bool // weak by default
	insecure        bool // secure by default
	defaultRegistry string
	defaultTag      string
}

type Option func(*options)

func makeOptions(opts ...Option) options {
	opt := options{
		defaultRegistry: DefaultRegistry,
		defaultTag:      DefaultTag,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

//example of functional call

func WithRegistry(registry string) Option {
	//returns an anonymous function which sets the default registry
	//this instance confomrs to the Option interface
	return func(opt *options) {
		opt.defaultRegistry = registry
	}
}

func WithStrict(s bool) Option {
	return func(opt *options) {
		opt.strict = s
	}
}

func main() {
	// Usage example:
	opt := makeOptions(WithStrict(true), WithRegistry("myRegistry"))
	//OR using the functional call
	opt2 := NewTag("myTag", WithStrict(true), WithRegistry("myRegistry"))
	fmt.Println(opt)
	fmt.Println(opt2)
}

func NewTag(name string, opts ...Option) string {
	opt := makeOptions(opts...)
	fmt.Println(opt)
	return name
}
