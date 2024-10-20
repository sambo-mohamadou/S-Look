package main

import (
	"flag"
	"fmt"
)

func main() {

	deepnessFlag := flag.Bool("deep", false, "Tell us to look trough all the system or just the User's folder. False by default")

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("There is no argument !! Please specify the name of the file you want to FOK ;)")
		return
	}

	if *deepnessFlag {
		fmt.Println("I see you want to go deep")
	} else {
		fmt.Println("I see that you are a teaser")
	}

}
