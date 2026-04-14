package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	file,err := os.Open("problems.csv")
	if err != nil{
		fmt.Println("Erro opening file")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records,err := reader.ReadAll()
	if err != nil{
		fmt.Println("Error reading CSV")
		return
	}
	score := 0

	for _, record := range records{
		var userAnswer string

		fmt.Println(record[0] + "= ?")
		fmt.Scanln(&userAnswer)

		if strings.TrimSpace(userAnswer) == strings.TrimSpace(record[1]){
			score++
		}
	}
	fmt.Println("your score:" , score, "/", len(records))
}
