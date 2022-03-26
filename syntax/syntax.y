
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
%token LEFTBRACKET
%token RIGHTBRACKET
%token LEFTANGLEBRACKET
%token RIGHTANGLEBRACKET
%token COMMA
%token SEMICOLON
%token COLON
%token LEFTPARENTHESIS
%token RIGHTPARENTHESIS
%token PERIOD

/* Operator tokens */
%token PLUS
%token ASTERIX
%token DIVISION
%token MODULO
%token DOUBLEQUAL
%token DIFFERENT
%token LESS
%token LESSOREQUALS
%token GREATER
%token GREATEROREQUALS
%token LOGICALAND
%token LOGICALOR
%token EQUAL
%token MINUS

%% /* The grammar follows */
Program : MainClass
 	| MainClass ClassDeclaration   ;

MainClass : class IDENTIFIER LEFTANGLEBRACKET public static void main LEFTPARENTHESIS String LEFTBRACKET RIGHTBRACKET IDENTIFIER RIGHTPARENTHESIS LEFTANGLEBRACKET Statement RIGHTANGLEBRACKET RIGHTANGLEBRACKET
ClassDeclaration : class IDENTIFIER Extension LEFTANGLEBRACKET  VarDeclaration   MethodDeclaration  RIGHTANGLEBRACKET
Extension:
	 | extends IDENTIFIER
	 ;
VarDeclaration :
		| Type IDENTIFIER SEMICOLON VarDeclaration;
Statement : LEFTANGLEBRACKET  Statement  RIGHTANGLEBRACKET
            | if LEFTPARENTHESIS Expression RIGHTPARENTHESIS Statement else Statement
            | while LEFTPARENTHESIS Expression RIGHTPARENTHESIS Statement
            | System.out.println LEFTPARENTHESIS Expression RIGHTPARENTHESIS SEMICOLON
            | IDENTIFIER EQUAL Expression SEMICOLON
            | IDENTIFIER LEFTBRACKET Expression RIGHTBRACKET EQUAL Expression SEMICOLON
            | ;
Type: int LEFTBRACKET RIGHTBRACKET
	| boolean
	| int
	| IDENTIFIER ;
MethodTypeDeclaration:
		| Type IDENTIFIER
		| Type IDENTIFIER COMMA MethodTypeDeclaration ;
MethodDeclaration : public Type IDENTIFIER LEFTPARENTHESIS MethodTypeDeclaration RIGHTPARENTHESIS LEFTANGLEBRACKET VarDeclaration Statement return Expression SEMICOLON RIGHTANGLEBRACKET
		  | public Type IDENTIFIER LEFTPARENTHESIS MethodTypeDeclaration RIGHTPARENTHESIS LEFTANGLEBRACKET VarDeclaration Statement return Expression SEMICOLON RIGHTANGLEBRACKET MethodDeclaration
		  | ;
Expression :
	    | Expression LOGICALAND Expression
	    | Expression LESS Expression
	    | Expression PLUS Expression
	    | Expression MINUS Expression
	    | Expression ASTERIX Expression
            |  Expression LEFTBRACKET Expression RIGHTBRACKET
            |  Expression PERIOD  length
            |  Expression PERIOD IDENTIFIER LEFTPARENTHESIS MethodExpressionSignature RIGHTPARENTHESIS
            |  INTEGER_LITERAL
            |  BOOLEAN_LITERAL
            |  IDENTIFIER
            | this
            | new int LEFTBRACKET Expression RIGHTBRACKET
            | new IDENTIFIER LEFTPARENTHESIS RIGHTPARENTHESIS
            | BANG Expression
            | LEFTPARENTHESIS Expression RIGHTPARENTHESIS
       	    ;
MethodExpressionSignature : Expression
			  | MethodExpressionSignature COMMA Expression
			  ;