package search

import "path/filepath"

func TipSearch(file string) string {
	return "Tip Search"
}

func GlobSearch(file string) ([]string, error) {
	matches, err := filepath.Glob(file)
	if err != nil {
		return nil, err
	} else {
		return matches, nil
	}
}
