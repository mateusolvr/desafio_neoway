package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

// SendDB : Send data to database
func SendDB(reader *csv.Reader) {

	db := ConnectDB()
	defer db.Close()

	i := 0
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		i++
		fmt.Println(i)

		if i == 1 {
			continue
		}

		columns := strings.Fields(line[0])

		cleanCPF, flagCPF := CPFUtils(columns[0])

		cleanCNPJLojaMaisFreq, flagCNPJLojaMaisFreq := CNPJUtils(columns[6])
		cleanCNPJLojaUltCompra, flagCNPJLojaUltCompra := CNPJUtils(columns[7])

		currentTime := time.Now()

		sqlStatement := "INSERT INTO public.analise_compra_usuario (cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, " +
			"loja_ultima_compra, inserido_em, cpf_valido, cnpj_mais_frequente_valido, cnpj_ultima_compra_valido) " +
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
		_, err := db.Exec(sqlStatement, cleanCPF, columns[1], columns[2], StringUtils(columns[3]), StringUtils(columns[4]), StringUtils(columns[5]), StringUtils(cleanCNPJLojaMaisFreq), StringUtils(cleanCNPJLojaUltCompra), currentTime, flagCPF, flagCNPJLojaMaisFreq, flagCNPJLojaUltCompra)
		// sqlStatement := "INSERT INTO public.analise_compra_usuario (cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, " +
		// 	"loja_ultima_compra, inserido_em, cpf_valido, cnpj_mais_frequente_valido, cnpj_ultima_compra_valido) " +
		// 	"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
		// _, err := db.Exec(sqlStatement, columns[0], columns[1], columns[2], nil, nil, nil, columns[6], columns[7], currentTime, true, true, true)

		if err != nil {
			panic(err)
		}
		// break
	}

}
