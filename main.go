package main

import (
	"flag"
	"fmt"
	"os"

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
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}
		pattern := home + "\\.*\\" + flag.Arg(0) + "*"
		fmt.Println(pattern)
		fmt.Println(search.GlobSearch(pattern))
	} else {
		fmt.Println(search.TipSearch(flag.Arg(0)))
	}

}
