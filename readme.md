# GoLexThatJava
This library is a Lexical Analysis for a minified java syntax.
You can find more info about the target syntax in this [Document](./TP_2022.doc.pdf)
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
oussema_zouaghi@pop-os:~/go/src/github.com/Mellywins/GoLexThatJava$ go run *.go
---------- START ----------
classTABIDENTIFIER{TAB publicTABstaticTABvoidTABmain(String[]TABIDENTIFIER){TAB System.out.println(newTABIDENTIFIER().IDENTIFIER(INTEGER_LITERAL));TABTAB }TAB}TABclassTABIDENTIFIERTAB{TAB intTABIDENTIFIER;TAB publicTABvoidTABIDENTIFIER(){TAB System.out.println(IDENTIFIER);TAB }TAB}TABclassTABIDENTIFIERTABextendsTABIDENTIFIERTAB{TAB publicTABintTABIDENTIFIER(intTABIDENTIFIER){TAB intTABIDENTIFIERTAB;TAB SINGLE_LINE_COMMENT TAB for(intTABIDENTIFIER=INTEGER_LITERAL;TABIDENTIFIER<=IDENTIFIER;TABIDENTIFIER++){TAB IDENTIFIERTAB=TABIDENTIFIER;TAB System.out.println(IDENTIFIER);TAB }TAB ifTAB(IDENTIFIERTAB<=TABINTEGER_LITERAL)TAB IDENTIFIERTAB=TABINTEGER_LITERALTAB;TAB MULTI_LINE_COMMENTTAB elseTAB IDENTIFIERTAB=TABIDENTIFIERTAB*TAB(this.IDENTIFIER(IDENTIFIERTAB-TABINTEGER_LITERAL))TAB;TAB returnTABIDENTIFIERTAB;TAB }TAB}
---------- END ----------
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