package lispy2
import (
	"fmt"
)

func addNodes(n Node, e *Environment) (Node) {
	sum := 0.0
	l := n.(List)
	for _, node := range l[1:] {
		a :=Eval(node, e)
		switch a.(type) {
		case F:
			sum += float64(a.(F))
		default:
			fmt.Printf("Addend %v of type %T, expecting type F\n", a, a)
		}
	}
	return F(sum)
}
func subtractNodes(n Node, e *Environment) (Node) {
	switch n.(type) {
	case List:
		l := n.(List)
		sum := float64(Eval(l[1], e).(F))
		for _, node := range l[2:] {
			sum -= float64(Eval(node, e).(F))
		}
		return F(sum)
	default:
	}
	return nil
}
func multiplyNodes(n Node, e *Environment) (Node) {
	multi := 1.0
	l := n.(List)
	for _, node := range l[1:] {
		a :=Eval(node, e)
		switch a.(type) {
		case F:
			multi *= float64(a.(F))
		default:
			fmt.Printf("Mulitiplican %v of type %T, expecting type F\n", a, a)
		}
	}
	return F(multi)
}
func divideNodes(n Node, e *Environment) (Node) {
	l := n.(List)
	var answer float64
	a := Eval(l[1], e)
	switch a.(type) {
	case F:
		b := Eval(l[2], e)
		switch b.(type) {
		case F:
			answer = float64(a.(F))
			answer /= float64(b.(F))
			return F(answer)
		default:
			fmt.Printf("Denominator %v of type %T not a float\n", b, b)
		}
	default:
		fmt.Printf("Numerator %v of type %T not a float\n", a, a)
	}
	return F(0)
}
func equalNodes(n Node, e *Environment) (Node) {
	l := n.(List)
	a := Eval(l[1], e)
	switch a.(type) {
	case F:
		b := Eval(l[2], e).(F)
		return B(a == b)
	case S:
		b := Eval(l[2], e).(S)
		return B(a == b)
	case B:
		b := Eval(l[2],e).(B)
		return B(b == a)
	}
	return B(false)
}
func cdr (n Node, e *Environment) (Node) {
	l := n.(List)
	a := Eval(l[1], e)
	switch a.(type) {
	case List:
		b := a.(List)[1:]
		return b
	default:
	}
	return List(nil)
}
func car (n Node, e *Environment) (Node) {
	l := n.(List)
	a := Eval(l[1], e)
	switch a.(type) {
	case List:
		b := a.(List)[0]
		return b
	default:
	}
	return List(nil)
}
func cons (n Node, e *Environment) (Node) {
	l := n.(List)
	var r List
	r = append(r, Eval(l[1], e))
	nxt := Eval(l[2], e)
	switch nxt.(type) {
	case List:
		l2 := nxt.(List)
		r = append(r, l2...)
	default:
		r = append(r, nxt)
	}
	return r
}
