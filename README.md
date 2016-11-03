#Lisp in Python in Go

Peter Norvig's [Lisp in Python](http://norvig.com/lispy.html) implemented
in Go.

I didn't try to be Scheme spec compliant, I just wanted to see how hard this
would be in a different programming language.

Golang's type system was probably the major barrier.

## Design Specifics

This set of Go types seems to work. I'm not sure it's any better than using
a list of `interface{}` as the List type.

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
    type Pr struct {
        body Node
        params Node
        env *Environment
    }
    func (pr Pr) isNode() { }


It took me a few minutes, and a good bit of googling, but the Golang interface
means that I didn't have to do anything dodgy to get a `Node` type that could
have a `[]Node` as one of its sub-types.

Norvig's Python code takes good advantage of Python magic like `__call__()` methods
to make executing a lambda expression look exactly the same as executing a built-in
primitive like '*'. Very clever.
