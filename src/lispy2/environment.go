package lispy2
import (
	"math"
)

type Environment map[Node]Node

var GlobalEnv Environment

func StandardEnv() (Environment) {
	GlobalEnv = make(map[Node]Node)
	GlobalEnv[S("+")] = Fn(addNodes)
	GlobalEnv[S("-")] = Fn(subtractNodes)
	GlobalEnv[S("*")] = Fn(multiplyNodes)
	GlobalEnv[S("/")] = Fn(divideNodes)
	GlobalEnv[S("==")] = Fn(equalNodes)
	GlobalEnv[S("pi")] = F(math.Pi)
	GlobalEnv[S("cdr")] = Fn(cdr)
	GlobalEnv[S("car")] = Fn(car)
	return GlobalEnv
}
