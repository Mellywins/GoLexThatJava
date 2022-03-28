
// This is an example of a goyacc program.
// To build it:
// goyacc -p "expr" expr.y (produces y.go)
// go build -o expr y.go
// expr
// > <type an expression>

%{
    package main
    import (
	"github.com/timtadh/lexmachine"
    )
%}
%union{
    token *lexmachine.Token
    ast   *Node
}
%token BANG
%token INTEGER_LITERAL
%token BOOLEAN_LITERAL
%token IDENTIFIER
%token CLASS
%token PUBLIC
%token STATIC
%token VOID
%token MAIN
%token STRING
%token SYSTEMOUTPRINTLN
%token RETURN
%token INT
%token IF
%token FOR
%token ELSE
%token WHILE
%token THIS
%token NEW
%token BOOLEAN
%token LENGTH
%token EXTENDS
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
Program : MainClass { yylex.(*golex).stmts = append(yylex.(*golex).stmts, $1.ast) }
 	| MainClass ClassDeclaration
 	{
 	yylex.(*golex).stmts=append(yylex.(*golex).stmts,$1.ast)
 	yylex.(*golex).stmts=append(yylex.(*golex).stmts,$2.ast)
 	} ;

MainClass : CLASS IDENTIFIER LEFTANGLEBRACKET PUBLIC STATIC VOID MAIN LEFTPARENTHESIS STRING LEFTBRACKET RIGHTBRACKET IDENTIFIER RIGHTPARENTHESIS LEFTANGLEBRACKET Statement RIGHTANGLEBRACKET RIGHTANGLEBRACKET
	{
	$$.ast=NewNode("MAINCLASS: ",nil).
	AddKid(NewNode("",$1.token)).
	AddKid(NewNode("",$2.token)).
	AddKid(NewNode("{",$3.token)).
	AddKid(NewNode("",$4.token)).
	AddKid(NewNode("",$5.token)).
	AddKid(NewNode("",$6.token)).
	AddKid(NewNode("",$7.token)).
	AddKid(NewNode("(",$8.token)).
	AddKid(NewNode("",$9.token)).
	AddKid(NewNode("[",$10.token)).
	AddKid(NewNode("]",$11.token)).
	AddKid(NewNode("",$12.token)).
	AddKid(NewNode(")",$13.token)).
	AddKid(NewNode("{",$14.token)).
	AddKid($15.ast).
	AddKid(NewNode("}",$16.token)).
	AddKid(NewNode("}",$16.token))


	}
;
ClassDeclaration : CLASS IDENTIFIER Extension LEFTANGLEBRACKET  VarDeclaration   MethodDeclaration  RIGHTANGLEBRACKET
{
	**
}
Extension: EXTENDS IDENTIFIER;
VarDeclaration : Type IDENTIFIER SEMICOLON VarDeclaration
		| Type IDENTIFIER SEMICOLON;
Statement : LEFTANGLEBRACKET  Statement  RIGHTANGLEBRACKET {
		$$.ast=NewNode("bracketed statement",nil).
			AddKid(NewNode("{",$1.token)).
			AddKid($2.ast).
			AddKid(NewNode("}",$3.token))
		}
            | IF LEFTPARENTHESIS Expression RIGHTPARENTHESIS Statement ELSE Statement
            | WHILE LEFTPARENTHESIS Expression RIGHTPARENTHESIS Statement
            | SYSTEMOUTPRINTLN LEFTPARENTHESIS Expression RIGHTPARENTHESIS SEMICOLON
            | IDENTIFIER EQUAL Expression SEMICOLON
            | IDENTIFIER LEFTBRACKET Expression RIGHTBRACKET EQUAL Expression SEMICOLON
            | {$$.ast=NewNode("empty statement content",nil) } ;
Type: INT LEFTBRACKET RIGHTBRACKET
	| BOOLEAN
	| INT
	| IDENTIFIER ;
MethodTypeDeclaration:
		| Type IDENTIFIER
		| Type IDENTIFIER COMMA MethodTypeDeclaration ;
MethodDeclaration : PUBLIC Type IDENTIFIER LEFTPARENTHESIS MethodTypeDeclaration RIGHTPARENTHESIS LEFTANGLEBRACKET VarDeclaration Statement RETURN Expression SEMICOLON RIGHTANGLEBRACKET
		  | PUBLIC Type IDENTIFIER LEFTPARENTHESIS MethodTypeDeclaration RIGHTPARENTHESIS LEFTANGLEBRACKET VarDeclaration Statement RETURN Expression SEMICOLON RIGHTANGLEBRACKET MethodDeclaration
		  | ;
Expression : IDENTIFIER
	    | Expression LOGICALAND Expression
	    | Expression LESS Expression
	    | Expression PLUS Expression
	    | Expression MINUS Expression
	    | Expression ASTERIX Expression
            | INTEGER_LITERAL
            | BOOLEAN_LITERAL
            |  Expression LEFTBRACKET Expression RIGHTBRACKET
            |  Expression PERIOD  LENGTH
            |  Expression PERIOD IDENTIFIER LEFTPARENTHESIS MethodExpressionSignature RIGHTPARENTHESIS
            | THIS
            | NEW INT LEFTBRACKET Expression RIGHTBRACKET
            | NEW IDENTIFIER LEFTPARENTHESIS RIGHTPARENTHESIS
            | BANG Expression
            | LEFTPARENTHESIS Expression RIGHTPARENTHESIS
       	    ;
MethodExpressionSignature : Expression
			  | MethodExpressionSignature COMMA Expression
			  ;
%%