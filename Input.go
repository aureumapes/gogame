package gogame

import "fmt"

func ReadInput(text string) rune {
	print(text)
	var input string
	fmt.Scanln(&input)
	return []rune(input)[0]
}
