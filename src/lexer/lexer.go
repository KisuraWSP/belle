package lexer

// lexical analysis (generate tokens for each keyword or expression evaluated)
type token struct{
	order int // used for the expression generation to determine sequence order
	content *tree // not sure need to do more research
	type_ int // used to identify whether if it is a reserved keyword or not
}

func (t *token) Check_if_reserved(){
	// if the keyword is reserved then do some operation with it
	// if not its a value by default to that operations type 
}


func Generate_expression(tokenlist []*token){
	//generate a valid expression and make sure it works with the programming languages logic 	
}
