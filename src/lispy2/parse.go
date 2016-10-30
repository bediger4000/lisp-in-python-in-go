package lispy2
import (
	"fmt"
	"strconv"
)

type Node interface {
	isNode()
}
type List []Node
func (l List) isNode() { }

type I int
func (i I) isNode() { }
type F float64
func (f F) isNode() { }
type S string
func (s S) isNode() { }
type Fn func(Node)(Node)
func (fn Fn) isNode() { }
type B bool
func (b B) isNode() { }

func Parse(program string) (Node, error) {
	_, l, e := realParse(Tokenize(program))
	return l, e
}

func realParse(tokens []string) ([]string, Node, error) {
	if len(tokens) < 1 {
		return nil, nil, fmt.Errorf("unexpected EOF while reading")
	}
	switch tokens[0] {
	case "(":
		tokens = tokens[1:]
		var l List
		var item Node
		var e error
		for tokens[0] != ")" {
			tokens, item, e = realParse(tokens)
			if e != nil {
				return tokens, l, e
			}
			l = append(l, item)
		}
		tokens = tokens[1:]
		return tokens, l, nil
	case ")":
		return nil, nil, fmt.Errorf("unexpected ')'")
	}


	token := tokens[0]
	tokens = tokens[1:]
	return tokens, atom(token), nil
}

func atom(token string) (Node) {
	floatval, err := strconv.ParseFloat(token, 64)
	if err == nil { return F(floatval) }
	intval, err := strconv.Atoi(token)
	if err == nil { return I(intval) }
	if token == "true" { return B(true) }
	if token == "false" { return B(false) }
	return S(token)
}
