package main

import (
	"context"
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

func printScore(score int) {
	fmt.Println("You have scored ", score)
}

func main() {

	var (
		filename string
		correct  int
		timeout  int
	)

	flag.StringVar(&filename, "filename", "problem.csv", "Please specify file from where quiz game will be loaded")
	flag.IntVar(&timeout, "timeout", 600, "Provide timeout of Quiz game. After provide time, game will stopped and result will be declared")

	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println()
			printScore(correct)
			os.Exit(0)
		}
	}()
	records := readCsvFile("problem.csv")
	for i, r := range records {
		printProblem(i+1, r)
		var answer string
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			fmt.Println()
			fmt.Println("You have scored ", correct)
			log.Fatal(err.Error())
		}
		if r[1] == strings.TrimSpace(answer) {
			correct++
		}
	}
	// fmt.Println("You have scored ", correct)

}
