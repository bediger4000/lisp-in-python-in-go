package lispy2
import (
)

type Environment map[Node]Node

var GlobalEnv Environment

func StandardEnv() (Environment) {
	GlobalEnv = make(map[Node]Node)
	GlobalEnv[S("twelve")] = F(12)
	GlobalEnv[S("+")] = Fn(addNodes)
	GlobalEnv[S("-")] = Fn(subtractNodes)
	GlobalEnv[S("*")] = Fn(multiplyNodes)
	GlobalEnv[S("/")] = Fn(divideNodes)
	GlobalEnv[S("define")] = Fn(defineSymbol)
	GlobalEnv[S("begin")] = Fn(evalBegin)
	return GlobalEnv
}
