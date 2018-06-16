package helloworld

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello, World.")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
