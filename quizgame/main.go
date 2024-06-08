package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func readCsvFile(name string) [][]string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("Unable to Open Csv file. Error: ", err.Error())
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unalbe to read records from csv file. Error: ", err.Error())
	}
	return records
}

func printProblem(id int, input []string) {
	fmt.Printf("Problem #%d: %s ?", id, input[0])
}

func printScore(score, totalscore int) {
	fmt.Printf("You have scored %d out of %d", score, totalscore)
}

func main() {

	var (
		filename   string
		score      int
		timeout    int
		totalscore int
	)

	flag.StringVar(&filename, "filename", "problem.csv", "Please specify file from where quiz game will be loaded")
	flag.IntVar(&timeout, "timeout", 600, "Provide timeout of Quiz game. After provide time, game will stopped and result will be declared")

	flag.Parse()

	timer := time.After(time.Duration(timeout) * time.Second)
	go func() {
		<-timer
		fmt.Println()
		printScore(score, totalscore)
		os.Exit(0)
	}()
	records := readCsvFile("problem.csv")
	totalscore = len(records)
	for i, r := range records {
		printProblem(i+1, r)
		var answer string
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			fmt.Println()
			printScore(score, totalscore)
			log.Fatal(err.Error())
		}
		if r[1] == strings.TrimSpace(answer) {
			score++
		}
	}
	// fmt.Println("You have scored ", correct)

}
