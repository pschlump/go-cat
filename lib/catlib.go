package catlib

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pschlump/Go-FTL/server/sizlib"
)

type GoCatCfg struct {
	// add line numbering
	// add de-tabbing
	// others...
}

func CatFiles(cfg *GoCatCfg, output string, files ...string) {
	var err error

	of := os.Stdout
	if output != "" {
		of, err = sizlib.Fopen(output, "w")
		if err != nil {
			fmt.Fprintf(os.Stderr, "go-cat: Error: Unable to open output file %s, %s\n", output, err)
			os.Exit(1)
		}
	}

	for _, vv := range files {
		buf, err := ioutil.ReadFile(vv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "go-cat: Error: Unable to read file %s, %s\n", vv, err)
		}
		fmt.Fprintf(of, "%s", buf)
	}

}
