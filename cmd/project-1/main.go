package main

import (
	"github.com/adamavixio/demo/pkg/pkg1"
	"github.com/adamavixio/demo/pkg/pkg2"
)

// Notice there is only a "Hello" method, but there is
// no "hello" lowercase method, as lowercase methods
// are private to the package

func main() {
	pkg1.Hello("testing package 1")
	pkg2.Hello("testing package 2")
}
