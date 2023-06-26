package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readCsvFile(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Could not open the requested file")
	}

	reader := csv.NewReader(f)

	data, rErr := reader.ReadAll()

	if rErr != nil {
		log.Fatal("Could not read the file")
	}

	return data
}

func main() {
	data := readCsvFile("./questions.csv")

	score := 0

	for _, questionAndAnswer := range data {
		var i int
		fmt.Print(fmt.Sprintf("%v = ", questionAndAnswer[0]))
		fmt.Scan(&i)

		correctAnswer, err := strconv.Atoi(questionAndAnswer[1])

		if err != nil {
			fmt.Printf("Your score is: %v\n", score)
			log.Fatal("Some server error, exiting quiz")
		}

		if i != correctAnswer {
			fmt.Println(fmt.Sprintf("Wrong answer, your score is: %v/%v", score, len(data)*10))
			os.Exit(1)
		} else {
			score += 10
		}
	}

	fmt.Println(fmt.Sprintf("Congrats for finishing the quiz, your score is: %v/%v", score, len(data)*10))
}
