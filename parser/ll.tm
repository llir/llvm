language llvm(go);

lang = "llvm"
package = "github.com/mewmew/l-tm/parser"

# Lexer

:: lexer

IntegerConstant : /0|[1-9][0-9]+/

# Parser

:: parser

input : IntegerConstant ;
