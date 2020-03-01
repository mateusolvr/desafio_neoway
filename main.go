package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"neoway/utils"
	"os"
	"path/filepath"
	"time"
)

func main() {

	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	db := utils.ConnectDB()
	defer db.Close()

	dirname := "myfiles"

	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		myfilepath := dirname + "/" + file.Name()
		fmt.Println("Reading file: ", myfilepath)
		extension := filepath.Ext(myfilepath)
		if extension != ".csv" && extension != ".txt" {
			fmt.Println("Extension not valid.")
			continue
		}
		csvfile, _ := os.Open(myfilepath)
		reader := csv.NewReader(bufio.NewReader(csvfile))
		reader.Comma = '\t'
		reader.FieldsPerRecord = -1
		utils.SendDBBULK(reader, db)
	}

	elapsed := time.Since(start)

	log.Printf("Binomial took %s", elapsed)

}
