package mx3

import (
	"fmt"
	"log"
	"strings"
)

const (
	o = "o"
	x = "x"
)

// scanなら繰り返しやれるな。版
func main() {
	for {
		var in string
		fmt.Scan(&in)

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
}
