package main

import (
	"github.com/adamavixio/demo/pkg/pkg1"
	"github.com/adamavixio/demo/pkg/pkg2"
)

func main() {
	pkg1.Hello("testing package 1 from project 2 now")
	pkg2.Hello("testing package 2 from project 2 now")
}
