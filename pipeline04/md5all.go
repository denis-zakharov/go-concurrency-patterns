package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"sort"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func main() {
	bound := flag.Int("n", 0, "define an upper-bound for parallel file md5 digestion")
	flag.Parse()

	var m map[string][md5.Size]byte
	var err error
	if *bound > 0 {
		m, err = MD5AllBounded(flag.Arg(0), *bound)
	} else {
		m, err = MD5All(flag.Arg(0))
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
