package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	o = "o"
	x = "x"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

// 標準入力を受けとる版
func main() {
	sc.Split(bufio.ScanWords)
	in := nextString()

	retMap := make(map[string]int, 3)
	var drawFlg bool
	for _, v := range strings.Split(in, "") {
		if v != o && v != x {
			log.Fatal("o x only")
		}
		retMap[v]++
		if retMap[v] == 3 {
			fmt.Println(v + ": win!")
			drawFlg = true
			break
		}
	}
	if !drawFlg {
		fmt.Println("draw")
	}
}
