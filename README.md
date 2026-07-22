# Go CLI Calculator

A robust, interactive Command-Line Interface (CLI) calculator built entirely in Go. 

This project goes beyond simple string evaluation by implementing a custom **Recursive Descent Parser** from scratch. It mathematically parses expressions, strictly enforces operator precedence (BODMAS/PEMDAS), and features persistent session history using Go's native File I/O capabilities.

## ✨ Core Features

*   **Custom Recursive Descent Parser:** Built completely from scratch without relying on external evaluation libraries. It safely parses mathematical strings and evaluates them hierarchically.
*   **Operator Precedence:** Accurately handles mathematical rules, prioritizing Exponentiation (`^`), then Multiplication/Division (`*`, `/`), and finally Addition/Subtraction (`+`, `-`).
*   **Custom Lexer (Tokenizer):** Intelligently processes raw user input, handling spaces dynamically and grouping multi-digit and decimal numbers correctly.
*   **Persistent History Tracking:** 
    *   **In-Memory:** Tracks the current session's operations for quick terminal review.
    *   **File I/O:** Automatically logs all operations to a permanent `history.txt` file using `os.O_APPEND` and safe file locking (`defer file.Close()`).
*   **Defensive Programming:** Gracefully handles invalid inputs, prevents infinite loops during parsing, and ensures safe memory address sharing using Go Pointers.

## 🛠️ Technologies & Concepts Applied
*   **Language:** Go (Golang)
*   **Architecture:** Modular design (Separation of Concerns between CLI loop, parsing logic, mathematical operations, and data storage).
*   **Concepts:** Structs & Methods (Receivers), Pointers, File I/O (Read/Write permissions `0644`), String Formatting (`fmt.Sprintf`), and Error Handling.

## 🚀 Installation & Usage

1. **Clone the repository:**
   ```bash
   git clone https://github.com/MariamA-Helal/CLI-Calculator.git
   cd CLI-Calculator

2. **Run the application:**
  ```bash
  go run main.go
  ```

3. **Commands:**

```plaintext
Enter any mathematical expression (e.g., 2 + 3 * 4 ^ 2).

Type history to view all past operations.

Type exit or quit to terminate the program safely.
```

## 💻 Output & Test Cases
**Here is a live demonstration of the calculator handling various operator precedence rules and decimal outputs, directly from the terminal:**

```plaintext
==== CLI Calculator Started... ====
Type 'exit' or 'quit' to close the program.
Type 'history' to review program history.
--------------------------------------------
Calc> 2+3*4
 your result: 14.00
Calc> 5* 2^ 3
 your result: 40.00
Calc> history

 ______ Operation History ______
1: 2+3*4 = 14.00
2: 5* 2^ 3 = 40.00

Calc> 10 -4/2 + 3 ^ 2
 your result: 17.00
Calc> 20 - 5 - 3 +2
 your result: 14.00
Calc> 2.5 *4 / 2
 your result: 5.00
Calc> 2.5 * 5 / 3
 your result: 4.17
Calc> history

 ______ Operation History ______
1: 2+3*4 = 14.00
2: 5* 2^ 3 = 40.00
3: 10 -4/2 + 3 ^ 2 = 17.00
4: 20 - 5 - 3 +2 = 14.00
5: 2.5 *4 / 2 = 5.00
6: 2.5 * 5 / 3 = 4.17

Calc> exit
Exiting calculator ... 
```

## 📁 Project Structure

```plaintext
CLI-Calculator/
├── main.go                 # Main CLI loop and user input handling
├── history.txt             # Auto-generated persistent history log
├── calculator/
│   ├── evaluator.go        # Custom Tokenizer and Recursive Descent Parser
│   ├── history.go          # History slice and File I/O logic
│   └── operations.go       # Core mathematical functions (Add, Sub, Mul, Div, Power)
└── README.md
```