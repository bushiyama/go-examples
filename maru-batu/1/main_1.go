package mx1

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

const (
	o = "o"
	x = "x"
)

var flagvar string

// 標準入力を受ける要件とずれるけど、パラメタにしちゃったほうが取り回しいいよ版
func main() {
	flag.StringVar(&flagvar, "p", "", "[ox]{5}")
	flag.Parse()

	retMap := make(map[string]int, 3)
	var drawFlg bool
	for _, v := range strings.Split(flagvar, "") {
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
