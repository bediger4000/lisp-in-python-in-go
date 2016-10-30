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
	GlobalEnv[S("define")] = Fn(defineSymbol)
	GlobalEnv[S("begin")] = Fn(evalBegin)
	GlobalEnv[S("quote")] = Fn(evalQuote)
	GlobalEnv[S("if")] = Fn(evalIf)
	GlobalEnv[S("car")] = Fn(evalCar)
	GlobalEnv[S("cdr")] = Fn(evalCdr)
	GlobalEnv[S("cons")] = Fn(evalCons)
	return GlobalEnv
}
