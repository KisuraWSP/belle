package lexer

import(
	"fmt"
)

/*
	ENUMS
*/
const (
	// reserved keywords
	Declare int = iota 
	Implement
	Proc
	Return
	Var
	If
	Then 
	Else 
	For 
	Call
	
			
	// types
	Int // signed integer (64 bit)
	Flt // signed float
	String // string
	Bool // Boolean
	Chr // char
	Any // Any type
	Error // Error type
	

	// Operators
	PLUS // +
	MINUS // -
	STAR // *
	SLASH_L // /
	EQUAL // =
	NOT // !
	OR // |
	AND // &
	TRIANGLE_BRACKET_L // <
	TRIANGLE_BRACKET_R  // >
	
	// Symbols
	SLASH_R // \ //
	SINGLEQUOTE // '
	DOUBLEQUOTE // "
	
	// Seperator
	BRACKET_L // (
	BRACKET_R // )
	SEMICOLON // ;
	COLON // :
	COMMA // ,
	DOT // .
)

// lexical analysis (generate tokens for each keyword or expression evaluated)
type token struct{
	Type int // checks the type of the token 
		     /*
		     	Possible types include:
		     		identifier - defined by the programmer
		     		keyword - reserved by the langauge
		     		seperator - ; , : , " , " , { , } , ( , )
		     		operator - + , > , =
		     		literal - "test", true, 2.232
		     		comment - discarded by the program
		     		whitespace - discarded by the program
		     */
	Value string // represents what the value that the token is evaluating 
}

func Create_token(value string, type int) (string, int){
	return value, type
}

func isAlpha(str: string){
    //return str.toUpperCase() != str.toLowerCase();
}

func isInt(str: string){
	//const c = str.charCodeAt(0);
    //const bounds = ['0'.charCodeAt(0), '9'.charCodeAt(0)];
    return (c >= bounds[0] && c<= bounds[1]);
}

func isSkippable(str: string) string {
    return str == " " || str == "\n" || str == "\t"
}

func isSpecial(str: string) string{
    return str == "->"
}

func Generate_expression(source_code string){
	//generate a valid expression and make sure it works with the programming languages logic
	/*
	we would want to express it like this as below
	      	eg:- x = 1;
	      	 [(identifier,x),(operator,=),(literal,1)]
	*/ 	
	/*
	
	const tokens = new Array<Token>();
    const src = source_code.split("");

    // build each token till end of file
    while(src.length > 0){
        // single character tokens
        if(src[0] == '('){
            tokens.push(create_token(src.shift(), TokenType.OpenParen));
        }else if (src[0] == ')'){
            tokens.push(create_token(src.shift(), TokenType.CloseParen));
        }else if (src[0] == '{'){
            tokens.push(create_token(src.shift(), TokenType.OpenCurly));
        }else if (src[0] == '}'){
            tokens.push(create_token(src.shift(), TokenType.CloseCurly));
        }else if (src[0] == '+' || src[0] == '-' ||  src[0] == '*' || src[0] == '/'){
            tokens.push(create_token(src.shift(), TokenType.BinaryOp));
        }else if (src[0] == ';'){
            tokens.push(create_token(src.shift(), TokenType.SemiColon));
        }else if (src[0] == ':'){
            tokens.push(create_token(src.shift(), TokenType.Colon));
        }else if (src[0] == '='){
            tokens.push(create_token(src.shift(), TokenType.Equals));
        }else if (src[0] == '|'){
            tokens.push(create_token(src.shift(), TokenType.Or));
        }else if (src[0] == '>'){
            tokens.push(create_token(src.shift(), TokenType.GreaterThan));
        }else if (src[0] == '<'){
            tokens.push(create_token(src.shift(), TokenType.LessThan));
        }else if (src[0] == '&'){
            tokens.push(create_token(src.shift(), TokenType.And));
        }else if (src[0] == '!'){
            tokens.push(create_token(src.shift(), TokenType.Not));
        }else if (src[0] == '#'){
            tokens.push(create_token(src.shift(), TokenType.Hash));
        }else if (src[0] == '@'){
            tokens.push(create_token(src.shift(), TokenType.At));
        }else if (src[0] == '.'){
            tokens.push(create_token(src.shift(), TokenType.Dot));
        }else{
            // multi-character tokens
            if(isInt(src[0])){
                let num = "";
                while(src.length > 0 && isInt(src[0])){
                    num += src.shift();
                }
                tokens.push(create_token(num, TokenType.Number))
            }else if(isAlpha(src[0])){
                let id = "";
                while(src.length > 0 && isAlpha(src[0])){
                    id += src.shift();
                }
                // check for reserved keywords
                const reserved = reserved_keywords[id];
                const spec_op = special_operators[id];
                const types = primitives[id];
                if(reserved){
                    tokens.push(create_token(id, reserved));
                }else if(spec_op){
                    tokens.push(create_token(id, spec_op));   
                }else if(id in primitives){
                    tokens.push(create_token(id, types));       
                }else{
                    tokens.push(create_token(id, TokenType.Identifier));
                }
            }else if(isSkippable(src[0])){
                src.shift();
            }else {
                console.error(
                    "Unrecognized Character found in Source",
                    src[0].charCodeAt(0),
                    src[0]
                );
                Deno.exit(1);
            }
        }
    }
    return tokens;
	
	*/
}
