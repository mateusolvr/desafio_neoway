package utils

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

// ReplaceSQL replaces the instance occurrence of any string pattern with an increasing $n based sequence
func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

// SendDBBULK : Send data to database
func SendDBBULK(reader *csv.Reader) {

	db := ConnectDB()
	defer db.Close()

	i := 0
	n := 0
	linesPerInsert := 200
	endOfData := false

	for {
		sqlStr := "INSERT INTO public.analise_compra_usuario (cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, " +
			"loja_ultima_compra, inserido_em, cpf_valido, cnpj_mais_frequente_valido, cnpj_ultima_compra_valido) " +
			"VALUES "
		vals := []interface{}{}
		n++

		for {
			line, error := reader.Read()
			if error == io.EOF {
				endOfData = true
				break
			} else if error != nil {
				log.Fatal(error)
			}

			i++

			if i == 1 {
				continue
			}

			columns := strings.Fields(line[0])

			cleanCPF, flagCPF := CPFUtils(columns[0])

			cleanCNPJLojaMaisFreq, flagCNPJLojaMaisFreq := CNPJUtils(columns[6])
			cleanCNPJLojaUltCompra, flagCNPJLojaUltCompra := CNPJUtils(columns[7])

			currentTime := time.Now()

			sqlStr += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
			vals = append(vals, cleanCPF, columns[1], columns[2], StringUtils(columns[3]), StringUtils(columns[4]), StringUtils(columns[5]), StringUtils(cleanCNPJLojaMaisFreq), StringUtils(cleanCNPJLojaUltCompra), currentTime, flagCPF, flagCNPJLojaMaisFreq, flagCNPJLojaUltCompra)

			if i == linesPerInsert*n {
				break
			}
		}
		//trim the last ,
		sqlStr = strings.TrimSuffix(sqlStr, ",")

		//Replacing ? with $n for postgres
		sqlStr = ReplaceSQL(sqlStr, "?")

		//prepare the statement
		stmt, _ := db.Prepare(sqlStr)

		//format all vals at once
		_, err := stmt.Exec(vals...)
		if err != nil {
			panic(err)
		}
		if endOfData == true {
			break
		}
	}

}
