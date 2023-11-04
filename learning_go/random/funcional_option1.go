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

func makeOptions() options {
	//basically to update opts before and then make all the fucntional calls, neat trick to update the struct.
	opt := options{
		defaultRegistry: DefaultRegistry,
		defaultTag:      DefaultTag,
	}
	return opt
}

//example of functional call

func WithRegistry(opt *options, registry string) {
	opt.defaultRegistry = registry
}

func WithStrict(opt *options, s bool) {
	opt.strict = s
}

func main() {
	// Usage example:
	opt := makeOptions()
	WithRegistry(&opt, "myRegistry")
	WithStrict(&opt, true)
	fmt.Println(opt)
}

//if we have too many options, we can use a map instead of a struct
//but we lose the type safety

//explain opt map[string]interface{}
