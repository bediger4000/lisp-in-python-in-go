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
	switch n.(type) {
	case B:
		fmt.Printf("%v", n)
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
	default:
		fmt.Printf("%v - %T", n, n)
	}
}
