
// This is an example of a goyacc program.
// To build it:
// goyacc -p "expr" expr.y (produces y.go)
// go build -o expr y.go
// expr
// > <type an expression>

%{
    package syntax
    import (
        "fmt"
        "log"

    )
%}
%union{
    token *lexmachine.Token
    ast   *Node
}