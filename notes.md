TESTING

filename: xxx_test.go
start test functions with 'Test'
takes arg t *testing.T
coverage: go test -cover

BENCHMARKING

When the benchmark code is executed, it runs b.N times and measures how long it takes.
To run: go test -bench=.


STRINGS

%q - wrap in quotes


FUNCTIONS

args of same type: e.g. (got string, want string) you can shorten to (got, want string)