package utils

import "fmt"

type TestCase struct {
	FileName string
	Expected int64
	Fn       func(string) int64
}

func Run(testCases []TestCase) {
	for _, testCase := range testCases {
		actual := testCase.Fn(testCase.FileName)
		if testCase.Expected == -1 {
			fmt.Println(actual)
		} else if testCase.Expected != actual {
			fmt.Printf("Expected: %d Actual: %d\n", testCase.Expected, actual)
			break
		} else {
			fmt.Println("Passed!")
		}
	}
}
