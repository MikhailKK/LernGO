package consolereader

import (
	"bufio"
	"fmt"
	"os"
)

func consolereader() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
