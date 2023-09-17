package main

import "fmt"

func main() {
	var row, col int
	var rowCount int
	var colCount int
	fmt.Scan(&row, &col)
	if (row >= 1 && row <= 31) && (col >= 1 && col <= 31) {
		fmt.Print("  |")
		if row > col {
			for rowCount = 1; rowCount < row; rowCount++ {
				fmt.Print("   ", rowCount)
			}
		} else if row == col {
			for rowCount = 1; rowCount <= row; rowCount++ {
				fmt.Print("   ", rowCount)
			}
		} else {
			for colCount = 1; colCount <= col; colCount++ {
				fmt.Print("   ", colCount)
			}
		}
		fmt.Println("")
		fmt.Print("---")
		if row > col {
			for rowCount = 1; rowCount < row; rowCount++ {
				fmt.Print("----")
			}
		} else if row == col {
			for rowCount = 1; rowCount <= row; rowCount++ {
				fmt.Print("----")
			}
		} else {
			for colCount = 1; colCount <= col; colCount++ {
				fmt.Print("----")
			}
		}
	}
	fmt.Println("")
	for i := 1; i <= row; i++ {
		fmt.Print(i, " |")
		for j := 1; j <= col; j++ {
			if i*j+1 < 10 {
				fmt.Print("   ", i*j)
			} else if i*j == 9 {
				fmt.Print("   ", i*j)
			} else {
				fmt.Print("  ", i*j)
			}
		}
		fmt.Println()
	}

	var input string
	fmt.Scanf("%v", &input)
	fmt.Scanf(" ")
}
