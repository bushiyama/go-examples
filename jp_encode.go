package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// UTF-8 -> ShiftJIS
func utf8_to_sjis(str string) (string, error) {
	iostr := strings.NewReader(str)
	rio := transform.NewReader(iostr, japanese.ShiftJIS.NewEncoder())
	ret, err := ioutil.ReadAll(rio)
	if err != nil {
		return "", err
	}
	return string(ret), err
}

// UTF-8 <- ShiftJIS
func sjis_to_utf8(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}

func main() {
	str := "髙橋"
	msg, err := utf8_to_sjis(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)

	msg, err = sjis_to_utf8(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
}
