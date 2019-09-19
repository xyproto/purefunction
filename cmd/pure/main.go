package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/purefunction"
	"github.com/xyproto/textoutput"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"
)

func main() {
	o := textoutput.NewTextOutput(true, true)
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Please provide one or more filenames for analysis.")
		os.Exit(1)
	}
	filenames := flag.Args()
	sort.Strings(filenames)

	// Output the pure functions, per filename
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	for _, filename := range filenames {
		pureFunctions, err := purefunction.PureFunctions(filename)
		if err != nil {
			o.ErrExit(err.Error())
		}
		if len(pureFunctions) == 0 {
			fmt.Fprintf(tw, o.LightTags("<darkgray>[<cyan>%s<darkgray>]<off>\t<white>%s<off>\n"), filepath.Base(filename), "None")
			continue
		}
		sort.Strings(pureFunctions)
		for _, name := range pureFunctions {
			fmt.Fprintf(tw, o.LightTags("<darkgray>[<cyan>%s<darkgray>]<off>\t<white>%s<off>\n"), filepath.Base(filename), name)
		}
	}
	tw.Flush()
}
