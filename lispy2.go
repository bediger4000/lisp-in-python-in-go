package main
import (
	"os"
	"fmt"
	"lispy2"
)

func main() {
	list, err := lispy2.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("Problem with parse: %s\n", err)
		os.Exit(1)
	}

	lispy2.PrintNode(list)
	fmt.Printf("\n")

	env := lispy2.StandardEnv()

	n := lispy2.Eval(list, env)
	lispy2.PrintNode(n)
	fmt.Printf("\n")

	os.Exit(0)
}
