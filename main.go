package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("==== CLI Calculator Started... ====")
	fmt.Println("Type 'exit' or 'quit' to close the program.")
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

		fmt.Printf(" you entered: %s\n ", input)

	}

	// sum := calculator.Add(5.73, 4.538)
	// fmt.Printf("Test add: 5.73 + 4.538 = %.2f \n", sum)

	// decAns, remainder, err := calculator.Div(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("Test div: 10 / 3 = %.2f \n(Remainder: %d) \n", decAns, remainder)
	// }

}

//
