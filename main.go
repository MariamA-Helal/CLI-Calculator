package main

import (
	"fmt"

	"github.com/MariamA-Helal/CLI-Calculator/calculator"
)

func main() {
	fmt.Println("==== CLI Calculator Started... ====")

	sum := calculator.Add(5.73, 4.538)
	fmt.Printf("Test add: 5.73 + 4.538 = %.2f \n", sum)

	decAns, remainder, err := calculator.Div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test div: 10 / 3 = %.2f \n(Remainder: %d) \n", decAns, remainder)
	}

}
