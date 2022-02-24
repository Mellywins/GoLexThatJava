# GoLexThatJava
This library is a Lexical Analysis for a minified java syntax.
You can find more info about the target syntax in this [Document](./TP_2022.doc.pdf)
## Execution example
Running this program on [example.java](./example.java) yields: 
```
> go run *.go

---------- START ----------
class IDENTIFIER{
 public static void main(String[] IDENTIFIER){
 System.out.println(new IDENTIFIER().IDENTIFIER(INTEGER_LITERAL)); 
 }
}
class IDENTIFIER {
 public void IDENTIFIER(){
 System.out.println(panic: Lexer error: could not match text starting at 8:28 failing at 8:29.
        unmatched text: "\""

goroutine 1 [running]:
main.must(...)
        /home/oussema_zouaghi/go/src/github.com/Mellywins/GoLexThatJava/lexemes.go:32
main.main()
        /home/oussema_zouaghi/go/src/github.com/Mellywins/GoLexThatJava/main.go:34 +0x3ec
exit status 2
oussema_zouaghi@pop-os:~/go/src/github.com/Mellywins/GoLexThatJava$ go run *.go
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
 SINGLE_LINE_COMMENT IDENTIFIER(int IDENTIFIER=INTEGER_LITERAL; IDENTIFIER<=IDENTIFIER; IDENTIFIER++){
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
```