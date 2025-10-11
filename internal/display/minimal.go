package display

import "fmt"

func Minimal(text string) {
	fmt.Print("\x1b[?25l")

	fmt.Printf("%v\n", text)

}
