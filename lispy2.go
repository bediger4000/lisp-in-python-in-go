package main
import (
	"os"
	"bufio"
	"fmt"
	"lispy2"
)

func main() {
	env := lispy2.StandardEnv()
	input:= bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	for input.Scan() {
		list, err := lispy2.Parse(input.Text())
		if err != nil {
			fmt.Printf("Problem with parse of %q: %s: %s\n", input.Text(), err)
			continue
		}

		lispy2.PrintNode(list)
		fmt.Printf("\n")


		n := lispy2.Eval(list, env)
		lispy2.PrintNode(n)
		fmt.Printf("\n> ")
	}

	os.Exit(0)
}
