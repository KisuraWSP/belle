package repl

import(
	"fmt"
	"log"
)

type method interface{
	Pop()
	Push(value string)
	Peek()
}

type stack struct{
	region []string
	top int	
}

func (s *stack) Pop() {
	// Pop the stack	
	fmt.Print("TO BE IMPLEMENTED\n")
}

func (s *stack) Push(value string) {
	// Push the stack
	fmt.Print("TO BE IMPLEMENTED->")
	fmt.Printf("%s\n",value)
}

func (s *stack) Peek() {
	// Peek the stack
	fmt.Print("TO BE IMPLEMENTED\n")
}

// REPL (READ-EVALUATE-PRINT-LOOP)
func Load() {
	var input string
	s := &stack{}
	for {
		fmt.Print("RUN=> ")
		fmt.Scanln(&input)
		if input == "exit" {
			log.Println("Exiting REPL")
			break
		}
		if input == "peek" {
			s.Peek()
		}
		if input == "pop" {
			s.Pop()
		}
		s.Push(input)
	}
}
