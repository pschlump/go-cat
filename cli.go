package main

import (
	"flag"

	catlib "github.com/pschlump/go-cat/lib"
)

var Output = flag.String("output", "", "Output - defaults to standard output") // 0
func init() {
	flag.StringVar(Output, "o", "", "Output - defaults to standard output") // 0
}

func main() {

	flag.Parse()
	fns := flag.Args()

	cfg := &catlib.GoCatCfg{}

	catlib.CatFiles(cfg, *Output, fns...)

}
