TESTING

filename: xxx_test.go
start test functions with 'Test'
takes arg t *testing.T
coverage: go test -cover

t.Run: t.Run("make the sums of some slices", func(t *testing.T) {


BENCHMARKING

When the benchmark code is executed, it runs b.N times and measures how long it takes.
To run: go test -bench=.


STRINGS

%q - wrap in quotes


FUNCTIONS

args of same type: e.g. (got string, want string) you can shorten to (got, want string)

ARRAYS/SLICES

slices.Equal - use for equality

INTERFACES

e.g. 
type Shape interface {
	Area() float64
}
In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

FOREVER

forever go run . - loop execution forever