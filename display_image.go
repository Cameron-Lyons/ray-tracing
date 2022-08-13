package main

import (
	"fmt"
)

func main() {

	const image_width int = 256
	const image_height int = 256

	fmt.Print("P3\n", image_width, ' ', image_height, "\n255\n")

}
