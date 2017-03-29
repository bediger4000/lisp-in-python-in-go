# Lisp in Python in Go

Peter Norvig's [Lisp in Python](http://norvig.com/lispy.html) implemented
in Go.

I didn't try to be Scheme spec compliant, I just wanted to see how hard this
would be in a different programming language. I certainly wrote more Golang
than Norvig wrote Python.

The fun is doing 3 languages all at once. Bearing in mind that I'm not a Lisp
programmer, I got to write an interpreter for a Lisp-style language in Go,
while understanding the original interpreter in Python.  Golang's type system
was probably the major barrier.

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

## Verification

Beyond just typing in simple forms to see if the pieces work, I decided to quit
when it could do two things:

1. Calculate factorial 
2. Execute a self-replicating expression.

It can run two variations on factorial:

	(define fact (lambda (x) (if (== x 1) 1 (* x (fact (- x 1))))))
	(fact 5)
	(define fact2 (lambda (y x) (if (== x 1) 1 (* x (y y (- x 1))))))
	(fact2 fact2 5)

And it runs a self-replicating program:

    ((lambda (y) (cons y y)) (lambda (y) (cons y y)))

This isn't exactly the same as the "canonical" Scheme self-replicating
program, which has `list` where I have `cons`. I'm not sure I understand
the subtle differences between lists made with the two function.
