package lispy2
import (
)

func addNodes(n Node) (Node) {
	switch n.(type) {
	case List:
		sum := 0.0
		l := n.(List)
		for _, node := range l[1:] {
			sum += float64(Eval(node, GlobalEnv).(F))
		}
		return F(sum)
	default:
	}
	return nil
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
	switch n.(type) {
	case List:
		l := n.(List)
		answer := float64(Eval(l[1], GlobalEnv).(F))
		for _, node := range l[2:] {
			answer *= float64(Eval(node, GlobalEnv).(F))
		}
		return F(answer)
	default:
	}
	return nil
}
func divideNodes(n Node) (Node) {
	switch n.(type) {
	case List:
		l := n.(List)
		answer := float64(Eval(l[1], GlobalEnv).(F))
		answer /= float64(Eval(l[2], GlobalEnv).(F))
		return F(answer)
	default:
	}
	return nil
}
func defineSymbol(n Node) (Node) {
	switch n.(type) {
	case List:
		l := n.(List)
		symbol := Eval(l[1], GlobalEnv)
		value  := Eval(l[2], GlobalEnv)
		GlobalEnv[symbol] = value
		return value
	default:
	}
	return nil
}
func evalBegin(n Node) (Node) {
	l := n.(List)
	var answer Node
	for _, node := range l {
		answer = Eval(node, GlobalEnv)
	}
	return answer
}
