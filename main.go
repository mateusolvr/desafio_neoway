package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"neoway/utils"
)

func main() {

	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	csvfile, _ := os.Open("base_teste.txt")
	reader := csv.NewReader(bufio.NewReader(csvfile))
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	// fmt.Printf("%T\n", reader)

	utils.SendDBBULK(reader)
	// utils.SendDB(reader)

	elapsed := time.Since(start)

	log.Printf("Binomial took %s", elapsed)

}
