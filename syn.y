
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
%%

Program : Statement {yylex.(*golex).stmts=append(yylex.(*golex).stmts,$1.ast)} 
;

Statement: Expression Operator Expression
    {

        $$.ast = NewNode("Statement", nil).AddKid($2.ast).AddKid($1.ast).AddKid($3.ast)

        __yyfmt__.Printf("%s + %s",$1.token.Lexeme,$3.token.Lexeme)

        }
;
Expression: INTEGER_LITERAL {$$.ast = NewNode("INTEGER_LITERAL",$1.token)}
        ;
Operator: PLUS {$$.ast = NewNode("PLUS", $1.token)}
        | MINUS {$$.ast = NewNode("MINUS", $1.token)}
        ;
%%