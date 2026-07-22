package calculator

import (
	"fmt"
	"os"
)

var historyList []string

func AddRecord(expression string, result float64) {
	record := fmt.Sprintf("%s = %.2f", expression, result)

	historyList = append(historyList, record)

	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer file.Close()
		file.WriteString(record + "\n")
	} else {
		fmt.Println("Warning: Couldn't save to history file ! ")
	}

}

func PrintHistory() {
	if len(historyList) == 0 {
		fmt.Println(" History is empty.")
		return
	}

	fmt.Println("\n ______ Operation History ______")

	for i, record := range historyList {
		fmt.Printf("%d: %s\n", i+1, record)
	}

}
