package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MariamA-Helal/CLI-Calculator/calculator"
)

func main() {
	fmt.Println("==== CLI Calculator Started... ====")
	fmt.Println("Type 'exit' or 'quit' to close the program.")
	fmt.Println("Type 'history' to reveiw program history.")
	fmt.Println("--------------------------------------------")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Calc> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if strings.ToLower(input) == "exit" || strings.ToLower(input) == "quit" {
			fmt.Println("Exiting calculator ... ")
			break
		}

		if strings.ToLower(input) == "history" {
			calculator.PrintHistory()
			continue
		}

		result := calculator.Evaluate(input)

		fmt.Printf(" your result: %.2f\n ", result)

		calculator.AddRecord(input, result)

	}

}
