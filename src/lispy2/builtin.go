package lispy2
import (
	"fmt"
)

func addNodes(n Node) (Node) {
	sum := 0.0
	l := n.(List)
	for _, node := range l[1:] {
		a :=Eval(node, GlobalEnv)
		switch a.(type) {
		case F:
			sum += float64(a.(F))
		default:
			fmt.Printf("Addend %v of type %T, expecting type F\n", a, a)
		}
	}
	return F(sum)
}
func subtractNodes(n Node) (Node) {
	switch n.(type) {
	case List:
		l := n.(List)
		sum := float64(Eval(l[1], GlobalEnv).(F))
		for _, node := range l[2:] {
			sum -= float64(Eval(node, GlobalEnv).(F))
		}
		return F(sum)
	default:
	}
	return nil
}
func multiplyNodes(n Node) (Node) {
	multi := 1.0
	l := n.(List)
	for _, node := range l[1:] {
		a :=Eval(node, GlobalEnv)
		switch a.(type) {
		case F:
			multi *= float64(a.(F))
		default:
			fmt.Printf("Mulitiplican %v of type %T, expecting type F\n", a, a)
		}
	}
	return F(multi)
}
func divideNodes(n Node) (Node) {
	l := n.(List)
	var answer float64
	a := Eval(l[1], GlobalEnv)
	switch a.(type) {
	case F:
		b := Eval(l[2], GlobalEnv)
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
func equalNodes(n Node) (Node) {
	l := n.(List)
	a := Eval(l[1], GlobalEnv)
	switch a.(type) {
	case F:
		b := l[1].(F)
		c := l[2].(F)
		return B(b == c)
	case S:
		b := l[1].(S)
		c := l[2].(S)
		return B(b == c)
	case B:
		b := l[1].(B)
		c := l[2].(B)
		return B(b == c)
	}
	return B(false)
}
func cdr (n Node) (Node) {
	l := n.(List)
	a := Eval(l[1], GlobalEnv)
	switch a.(type) {
	case List:
		b := a.(List)[1:]
		return b
	default:
	}
	return List(nil)
}
func car (n Node) (Node) {
	l := n.(List)
	a := Eval(l[1], GlobalEnv)
	switch a.(type) {
	case List:
		b := a.(List)[0]
		return b
	default:
	}
	return List(nil)
}
func cons (n Node) (Node) {
	l := n.(List)
	var r List
	r = append(r, Eval(l[1], GlobalEnv))
	nxt := Eval(l[2], GlobalEnv)
	switch nxt.(type) {
	case List:
		l2 := nxt.(List)
		r = append(r, l2...)
	default:
		r = append(r, nxt)
	}
	return r
}
