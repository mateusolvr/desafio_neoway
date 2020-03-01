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

	// Coonect with the PostgreSQL database
	db := utils.ConnectDB()
	defer db.Close()

	// Scan the files in 'neoway/myfiles/' folder
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

	// Loop through the files
	for _, file := range files {
		myfilepath := dirname + "/" + file.Name()
		fmt.Println("Reading file: ", myfilepath)

		// Check if the file is a .csv or .txt, if not, continue
		extension := filepath.Ext(myfilepath)
		if extension != ".csv" && extension != ".txt" {
			fmt.Println("Extension not valid.")
			continue
		}

		// Open the file
		csvfile, _ := os.Open(myfilepath)
		reader := csv.NewReader(bufio.NewReader(csvfile))
		reader.Comma = '\t'
		reader.FieldsPerRecord = -1
		utils.SendDBBULK(reader, db)
	}

	elapsed := time.Since(start)

	log.Printf("Elapsed time: %s", elapsed)

}
