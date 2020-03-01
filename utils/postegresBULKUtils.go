package utils

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

// ReplaceSQL : Replace the "?" with an increasing $n sequence
func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

// SendDBBULK : Send data to database
func SendDBBULK(reader *csv.Reader, db *sql.DB) {

	i := 0
	n := 0
	rowsPerInsert := 200
	endOfData := false

	// Loop to insert only the maximum number of rows defined by the variable 'rowsPerInsert'
	for {
		sqlStr := "INSERT INTO public.analise_compra_usuario (cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, " +
			"loja_ultima_compra, inserido_em, cpf_valido, cnpj_mais_frequente_valido, cnpj_ultima_compra_valido) " +
			"VALUES "
		vals := []interface{}{}
		n++

		// Loop to bulk several rows and late insert them together
		for {
			line, error := reader.Read()
			if error == io.EOF {
				endOfData = true
				break
			} else if error != nil {
				log.Fatal(error)
			}

			i++

			// Ignore first row (headers)
			if i == 1 {
				continue
			}

			// Split into columns
			columns := strings.Fields(line[0])

			// Clean and validate the CPF
			cleanCPF, flagCPF := CPFUtils(columns[0])

			// Clean and validate the CNPJ
			cleanCNPJLojaMaisFreq, flagCNPJLojaMaisFreq := CNPJUtils(columns[6])
			cleanCNPJLojaUltCompra, flagCNPJLojaUltCompra := CNPJUtils(columns[7])

			//Get current time to insert in the DB
			currentTime := time.Now()

			sqlStr += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
			vals = append(vals, cleanCPF, columns[1], columns[2], StringUtils(columns[3]), StringUtils(columns[4]), StringUtils(columns[5]), StringUtils(cleanCNPJLojaMaisFreq), StringUtils(cleanCNPJLojaUltCompra), currentTime, flagCPF, flagCNPJLojaMaisFreq, flagCNPJLojaUltCompra)

			if i == rowsPerInsert*n {
				break
			}
		}
		fmt.Println("Number of inserted rows: ", i)
		// Trim the last ,
		sqlStr = strings.TrimSuffix(sqlStr, ",")

		// Replacing ? with $n sequence
		sqlStr = ReplaceSQL(sqlStr, "?")

		// Prepare the statement
		stmt, _ := db.Prepare(sqlStr)

		// Execute
		_, err := stmt.Exec(vals...)
		if err != nil {
			panic(err)
		}
		// If there's no more data in the file, exit loop
		if endOfData == true {
			break
		}
	}

}
