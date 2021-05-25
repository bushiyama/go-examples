package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal("n is int only.")
	}
	return i
}

func main() {
	// 英小文字からなる長さ N の文字列 S が与えられます。
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := nextString()

	ret := logic(n, s)

	// 操作が終わった後の T を求めてください。
	fmt.Println(ret)
}

func logic(n int, s string) string {
	// また、文字列 T があり、 T は、はじめ空文字列です。
	t := make([]string, n)

	// あなたは以下の操作を i=1,2,…,N に対して、 i=1 から順に行います。{
	for _, v := range strings.Split(s, "") {

		// T から c と同じ文字を全て削除した後、
		tmp := strings.Join(t, "")
		rep := strings.ReplaceAll(tmp, v, "")
		t = strings.Split(rep, "")

		// T の末尾に c を追加する。
		t = append(t, v)
	}
	return strings.Join(t, "")
}
