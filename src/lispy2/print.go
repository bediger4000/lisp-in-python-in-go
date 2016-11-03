package lispy2
import (
	"fmt"
)
func PrintList(list List) {
	fmt.Printf("(")
	comma := ""
	for _, item := range list {
		fmt.Printf("%s", comma)
		switch item.(type) {
		case List:
			PrintList(item.(List))
		default:
			PrintNode(item)
		}
		comma = " "
	}
	fmt.Printf(")")
}
func PrintNode(n Node) {
	if n == nil {
		fmt.Printf("nil")
		return
	}
	switch n.(type) {
	case B:
		if n.(B) {
			fmt.Printf("#t")
		} else {
			fmt.Printf("#f")
		}
	case I:
		fmt.Printf("%d", n)
	case F:
		fmt.Printf("%g", n)
	case S:
		fmt.Printf("%s", n)
	case List:
		PrintList(n.(List))
	case Fn:
		fmt.Printf("function")
	case Pr:
		fmt.Printf("(lambda ")
		PrintNode(n.(Pr).params)
		fmt.Printf(" ")
		PrintNode(n.(Pr).body)
		fmt.Printf(")")
	default:
		fmt.Printf("%v - %T", n, n)
	}
}
