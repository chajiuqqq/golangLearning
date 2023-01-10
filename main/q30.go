package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var countC, countW, countL int
	r := bufio.NewReader(os.Stdin)
	for {
		switch str, err := r.ReadString('\n'); true {
		case err != nil:
			fmt.Println("ch:%d,words:%d,lines:%d\n", countC, countW, countL)
			return
		default:
			countC += len([]rune(str))
			countW += len(strings.Fields(str))
			countL++
		}

	}

}
