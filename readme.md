# GoLexThatJava
This library is a Lexical Analysis for a minified java syntax.
You can find more info about the target syntax in this [Document](./TP_2022.doc.pdf)

This tool uses a DFA (Deterministic Finite Automata) as a backend to efficiently tokenize the Code text.
## The minjava BNF:
```BNF
BNF de MiniJAVA
Program ::= MainClass ( ClassDeclaration )* <EOF>
MainClass ::= "class" Identifier "{" "public" "static" "void" "main" "(" "String" "[""]" Identifier ")" "{" Statement "}" "}"
ClassDeclaration ::= "class" Identifier ( "extends" Identifier )? "{" ( VarDeclaration )* ( MethodDeclaration )* "}"
VarDeclaration ::= Type Identifier ";"
MethodDeclaration ::= "public" Type Identifier "(" ( Type Identifier ( "," Type Identifier )* )? ")""{" ( VarDeclaration )* ( Statement )* "return" Expression ";" "}"
Type ::= "int" "[" "]"
        | "boolean"
        | "int"
        | Identifier
Statement ::= "{" ( Statement )* "}"
            | "if" "(" Expression ")" Statement "else" Statement
            | "while" "(" Expression ")" Statement
            | "System.out.println" "(" Expression ")" ";"
            | Identifier "=" Expression ";"
            | Identifier "[" Expression "]" "=" Expression ";"
Expression ::= Expression ( "&&" | "<" | "+" | "-" | "*" ) Expression
            |  Expression "[" Expression "]"
            |  Expression "." "length"
            |  Expression "." Identifier "(" ( Expression ( "," Expression )* )? ")"
            |  <INTEGER_LITERAL>
            |  <BOOLEAN_LITERAL>
            |  Identifier
            | "this"
            | "new" "int" "[" Expression "]"
            | "new" Identifier "(" ")"
            | "!" Expression
            | "(" Expression ")"
Identifier ::= <IDENTIFIER>
```
## Ressources:
* Articles: https://hackthology.com/writing-a-lexer-in-go-with-lexmachine.html
## Execution example
Running this program on [example.java](./example.java) yields: 
```JAVA
> go run *.go
---------- START ----------
class IDENTIFIER{
    public static void main(String[] IDENTIFIER){
        System.out.println(new IDENTIFIER().IDENTIFIER(INTEGER_LITERAL)); 
    }
}
class IDENTIFIER {
    int IDENTIFIER;
    public void IDENTIFIER(){
        System.out.println(IDENTIFIER);
    }
}
class IDENTIFIER extends IDENTIFIER {
    public int IDENTIFIER(int IDENTIFIER){
        int IDENTIFIER ;
        SINGLE_LINE_COMMENT        
        for(int IDENTIFIER=INTEGER_LITERAL; IDENTIFIER<=IDENTIFIER; IDENTIFIER++){
            IDENTIFIER = IDENTIFIER;
            System.out.println(IDENTIFIER);
        }
        if (IDENTIFIER <= INTEGER_LITERAL)
            IDENTIFIER = INTEGER_LITERAL ;
            MULTI_LINE_COMMENT
        else
            IDENTIFIER = IDENTIFIER * (this.IDENTIFIER(IDENTIFIER - INTEGER_LITERAL)) ;
        return IDENTIFIER ;
    }
}
---------- END ----------
```