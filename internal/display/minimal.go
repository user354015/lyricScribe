package display

import "fmt"

func Display(text string) {
	fmt.Print("\x1b[?25l")

	fmt.Printf("%v\n", text)

}
