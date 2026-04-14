package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	readerr := bufio.NewReader(os.Stdin)
	
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
	fmt.Println("Press Enter to start quiz...")
	readerr.ReadString('\n')

	for _, record := range records{
		var userAnswer string

		fmt.Println("%s = ?\n" ,record[0])
		fmt.Scanln(&userAnswer)

		if strings.ToLower(strings.TrimSpace(userAnswer)) == strings. ToLower(strings.TrimSpace(record[1])){
			fmt.Println("Correct!")
			score++
		}else{
			fmt.Println("Wrong! correct answer is: %s\n", record[1])
		}
	}
	fmt.Printf("\nYou got %d out of %d correct\n", score, len(records))
}
