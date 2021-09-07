package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for i, item := range record {
			lineNo, column := r.FieldPos(i)
			fmt.Printf("lineNo:%d column:%d pos:%d record:%s\n", lineNo, i, column, item)
		}
	}

}
