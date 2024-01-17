package gogame

import (
	"fmt"
)

// ReadInput is a function, which allows you to get a single character input from the user
func ReadInput(text string) rune {
	print(text)
	var input string
	fmt.Scanln(&input)
	return []rune(input)[0]
}
