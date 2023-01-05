/*
Write a program which mimics the Unix program cat.
Make it support the n flag, where each line is numbered.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var n = flag.Bool("n", false, "whether to print line numbers")

func cat(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("err:", err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	i := 1
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if *n {
			w.WriteString(strconv.Itoa(i) + " ")
			i++
		}
		w.WriteString(str)
	}

}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("no param")
		return
	}
	for i := 0; i < flag.NArg(); i++ {
		// "/Users/chajiu/Desktop/project/golangLearning/readme.md"
		cat(flag.Arg(i))
	}
	// run:
	// ./q22  -n ../readme.md
}
