
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
	AddKid(NewNode("class",$1.token)).
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
					$$.ast= NewNode("<NewClassDeclaration>:", nil).
					AddKid(NewNode("",$1.token)).
					AddKid(NewNode("",$2.token)).
					AddKid($3.ast).
					AddKid(NewNode("{",$4.token)).
					AddKid($5.ast).
					AddKid($6.ast).
					AddKid(NewNode("}",$7.token))

			}
		| ClassDeclaration ClassDeclaration
			{
				$$.ast=NewNode("<ExtraClassDeclarations>:", nil).
					AddKid($1.ast).
					AddKid($2.ast)
			}
			;
Extension: EXTENDS IDENTIFIER
		{
		$$.ast=NewNode("<Extension> ",nil).
			AddKid(NewNode("",$1.token)).
			AddKid(NewNode("",$2.token))

		}
	| ;

VarDeclaration : VarDeclaration VarDeclaration
 			{
 			$$.ast=NewNode("<Variable definitions>:",nil).
 				AddKid($1.ast).
 				AddKid($2.ast)
 			 }

		| Type IDENTIFIER SEMICOLON
			{
				$$.ast=NewNode("",nil).
				AddKid($1.ast).
				AddKid(NewNode("",$2.token)).
				AddKid(NewNode("",$3.token))
			}
		| { $$.ast=nil }
		;

Statement : LEFTANGLEBRACKET  Statement  RIGHTANGLEBRACKET
		{
		$$.ast=NewNode("bracketed statement",nil).
			AddKid(NewNode("{",$1.token)).
			AddKid($2.ast).
			AddKid(NewNode("}",$3.token))
		}
            | IF LEFTPARENTHESIS Expression RIGHTPARENTHESIS Statement ELSE Statement
            	{
            	$$.ast=NewNode("<If block>",nil).
            		AddKid(NewNode("",$1.token)).
            		AddKid(NewNode("(",$2.token)).
            		AddKid($3.ast).
            		AddKid(NewNode(")",$4.token)).
            		AddKid($5.ast).
            		AddKid(NewNode("",$6.token)).
            		AddKid($7.ast)
            	}
            | WHILE LEFTPARENTHESIS Expression RIGHTPARENTHESIS Statement
            	{
            	$$.ast=NewNode("<While block>",nil).
            		AddKid(NewNode("",$1.token)).
            		AddKid(NewNode("",$2.token)).
            		AddKid($3.ast).
            		AddKid(NewNode(")",$4.token)).
            		AddKid($4.ast)

            	}
            | SYSTEMOUTPRINTLN LEFTPARENTHESIS Expression RIGHTPARENTHESIS SEMICOLON
            	{
            		$$.ast=NewNode("<Print Statement>",nil).
            			AddKid(NewNode("",$1.token)).
            			AddKid(NewNode("",$2.token)).
            			AddKid($3.ast).
            			AddKid(NewNode(")",$4.token)).
            			AddKid(NewNode("",$5.token))
            	}
            | IDENTIFIER EQUAL Expression SEMICOLON
            	{
            		$$.ast=NewNode("<Affectation>",nil).
            			AddKid(NewNode("",$1.token)).
            			AddKid(NewNode("=",$2.token)).
            			AddKid($3.ast).
            			AddKid(NewNode("",$4.token))
            	}
            | IDENTIFIER LEFTBRACKET Expression RIGHTBRACKET EQUAL Expression SEMICOLON
            	{
            	$$.ast=NewNode("<Array Affectation>",nil).
            		AddKid(NewNode("",$1.token)).
            		AddKid(NewNode("[",$2.token)).
            		AddKid($3.ast).
            		AddKid(NewNode("]",$4.token)).
            		AddKid(NewNode("=",$5.token)).
            		AddKid($6.ast).
            		AddKid(NewNode("",$7.token))
            	}
            | Statement Statement
            {
		$$.ast=NewNode("<BlockOfStatements>:",nil).
			AddKid($1.ast).
			AddKid($2.ast)
            }
            | {$$.ast=NewNode("empty statement content",nil) } ;
Type: INT LEFTBRACKET RIGHTBRACKET { $$.ast=NewNode("int[]",$1.token)}
	| BOOLEAN {$$.ast=NewNode("",$1.token)}
	| INT {$$.ast=NewNode("",$1.token)}
	| IDENTIFIER {$$.ast=NewNode("",$1.token)} ;
MethodTypeDeclaration:
		| Type IDENTIFIER
		{
			$$.ast=NewNode("",nil).
			AddKid($1.ast).
			AddKid(NewNode("",$2.token))
		}
		|  Type IDENTIFIER COMMA MethodTypeDeclaration
		{
		$$.ast = NewNode("<Method input type>:",nil).
			AddKid($1.ast).
			AddKid(NewNode("",$2.token)).
			AddKid(NewNode("",$3.token)).
			AddKid($4.ast)
		}
		 ;

MethodDeclaration : PUBLIC Type IDENTIFIER LEFTPARENTHESIS MethodTypeDeclaration RIGHTPARENTHESIS LEFTANGLEBRACKET VarDeclaration Statement RETURN Expression SEMICOLON RIGHTANGLEBRACKET
		  {
		  $$.ast=NewNode("<MethodDeclaration>:",nil).
			AddKid(NewNode("",$1.token)).
			AddKid($2.ast).
			AddKid(NewNode("",$3.token)).
			AddKid(NewNode("(",$4.token)).
			AddKid($5.ast).
			AddKid(NewNode(")",$6.token)).
			AddKid(NewNode("{",$7.token)).
			AddKid($8.ast).
			AddKid($9.ast).
			AddKid(NewNode("",$10.token)).
			AddKid($11.ast).
			AddKid(NewNode(";",$12.token)).
			AddKid(NewNode("}",$13.token))
		  }
		  | MethodDeclaration MethodDeclaration
		  {
		  	$$.ast=NewNode("<MethodDeclarations>:",nil).
				AddKid($1.ast).
				AddKid($2.ast)

		  }
		   ;

Expression : IDENTIFIER
		{
		$$.ast=NewNode("",$1.token)
		}
	    | Expression LOGICALAND Expression
	    	{
	    	$$.ast=NewNode("&&",$2.token).
	    		AddKid($1.ast).
	    		AddKid($3.ast)
	    	}
	    | Expression LESS Expression
	    	{
	    	$$.ast=NewNode("<",$2.token).
	    		AddKid($1.ast).
	    		AddKid($3.ast)
	    	}
	    | Expression PLUS Expression
	    	{
	    	$$.ast=NewNode("+",$2.token).
	    		AddKid($1.ast).
	    		AddKid($3.ast)
	    	}
	    | Expression MINUS Expression
	    	{
	    	$$.ast=NewNode("+",$2.token).
	    		AddKid($1.ast).
	    		AddKid($3.ast)
	    	}
	    | Expression ASTERIX Expression
	    	{
	    	$$.ast=NewNode("*",$2.token).
	    		AddKid($1.ast).
	    		AddKid($3.ast)
	    	}
            | INTEGER_LITERAL
            	{
            	$$.ast=NewNode("",$1.token)
            	}
            | BOOLEAN_LITERAL
            	{
            	$$.ast=NewNode("",$1.token)
            	}
            |  Expression LEFTBRACKET Expression RIGHTBRACKET
            	{
            	$$.ast=NewNode("",nil).
            		AddKid($1.ast).
            		AddKid(NewNode("[",$2.token)).
            		AddKid($3.ast).
            		AddKid(NewNode("]",$4.token))
            	}
            |  Expression PERIOD  LENGTH
            	{
            	$$.ast=NewNode("<PeriodAccess>:",nil).
            		AddKid($1.ast).
            		AddKid(NewNode(".",$2.token)).
            		AddKid(NewNode("",$3.token))
            	}
            |  Expression PERIOD IDENTIFIER LEFTPARENTHESIS MethodExpressionSignature RIGHTPARENTHESIS
            	{
            		$$.ast=NewNode("<AccessedFunction>", nil).
            			AddKid($1.ast).
            			AddKid(NewNode(".",$2.token)).
            			AddKid(NewNode("",$3.token)).
            			AddKid(NewNode("(",$4.token)).
            			AddKid($5.ast).
            			AddKid(NewNode(")",$6.token))
            	}
            | THIS
            	{
            	$$.ast=NewNode("",$1.token)
            	}
            | NEW INT LEFTBRACKET Expression RIGHTBRACKET
            	{
            		$$.ast=NewNode("",nil).
            			AddKid(NewNode("",$1.token)).
            			AddKid(NewNode("",$2.token)).
            			AddKid(NewNode("[",$3.token)).
            			AddKid($4.ast).
            			AddKid(NewNode("]",$5.token))
            	}
            | NEW IDENTIFIER LEFTPARENTHESIS RIGHTPARENTHESIS
            	{
            		$$.ast=NewNode("<Instantiation>:",nil).
            			AddKid(NewNode("",$1.token)).
            			AddKid(NewNode("",$2.token)).
            			AddKid(NewNode("(",$3.token)).
            			AddKid(NewNode(")",$4.token))
            	}
            | BANG Expression
            	{
            		$$.ast=NewNode("<NOT!>:",nil).
            			AddKid(NewNode("!",$1.token)).
            			AddKid($2.ast)
            	}
            | LEFTPARENTHESIS Expression RIGHTPARENTHESIS
            	{
            	$$.ast=NewNode("<ParenthesisedExpression>",nil).
            		AddKid(NewNode("(",$1.token)).
            		AddKid($2.ast).
            		AddKid(NewNode(")",$3.token))
            	}
       	    ;
MethodExpressionSignature : Expression
				{
				$$.ast=$1.ast
				}
			  | MethodExpressionSignature COMMA Expression
			  	{
			  	$$.ast=NewNode("<MultiParametersMethodSignature>",nil).
			  		AddKid($1.ast).
			  		AddKid(NewNode(".",$2.token)).
			  		AddKid($3.ast)
			  	}
			  ;
%%