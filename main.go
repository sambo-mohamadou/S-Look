package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {

	deepnessFlag := flag.Bool("deep", false, "Tell us to look trough all the system or just the User's folder. False by default")

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("There is no argument !! Please specify the name of the file you want to FOK ;)")
		return
	}

	if *deepnessFlag {

		// pattern := home + "\\*\\" + flag.Arg(0) + "*"
		// fmt.Println(pattern)
		// fmt.Println(search.GlobSearch(pattern))
		err := dipSearch(flag.Arg(0))
		if err != nil {
			fmt.Println("There has been an issue : ", err)
			return
		}
	} else {
		fmt.Println(tipSearch(flag.Arg(0)))
	}

}

func tipSearch(file string) string {
	return file
}

func dipSearch(target string) error {
	//target := file

	// matches, err := filepath.Glob(file)
	// if err != nil {
	// 	return nil, err
	// } else {
	// 	return matches, nil
	// }
	home, er := os.UserHomeDir()
	if er != nil {
		fmt.Println("Error getting home directory:", er)
		return er
	}

	err := filepath.WalkDir(home, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if ok, _ := filepath.Match("*"+target+"*", d.Name()); ok {
			fmt.Println(" ", path, d.IsDir())
		}
		return nil
	})
	return err
}
