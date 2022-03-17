
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
/* Special tokens */
%token SPACE
%token TAB
%token BANG
%token INTEGER_LITERAL
%token BOOLEAN_LITERAL
%token IDENTIFIER
%token SINGLE_LINE_COMMENT
%token MULTI_LINE_COMMENT
%token BREAK_LINE

/* Reserved words tokens */
%token class
%token public
%token static
%token void
%token main
%token String
%token System.out.println
%token return
%token int 
%token if 
%token for
%token else
%token while
%token this 
%token new
%token boolean 
%token length
%token extends

/* Literal tokens */
%token [
%token ] 
%token (
%token )
%token {
%token }
%token ,
%token ;
%token .
%token :

/* Operator tokens */
%token +
%token -
%token *
%token /
%token %
%token <
%token >
%token <=
%token >=
%token ==
%token !=
%token &&
%token ||
%token =

%% /* The grammar follows */
Program : MainClass ( ClassDeclaration )* <EOF>;