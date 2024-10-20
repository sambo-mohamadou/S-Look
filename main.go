package main

import (
	"flag"
	"fmt"

	"github.com/sambo-mohamadou/S-Lookup/search"
)

func main() {

	deepnessFlag := flag.Bool("deep", false, "Tell us to look trough all the system or just the User's folder. False by default")

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("There is no argument !! Please specify the name of the file you want to FOK ;)")
		return
	}

	if *deepnessFlag {
		fmt.Println(search.DipSearch("as"))
	} else {
		fmt.Println(search.TipSearch("as"))
	}

}
