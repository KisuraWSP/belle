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
	content []string
	top int	
}

func (s *stack) Pop() {
	if s.top > 0 {
		s.top--
		s.content = s.content[:s.top]
	} else {
		fmt.Println("Stack is empty, cannot pop.")
	}
}

func (s *stack) Push(value string) {
	s.content = append(s.content[:s.top], value)
	s.top++
}

func (s *stack) Peek() {
	if s.top > 0 {
		fmt.Println("Top element:", s.content[s.top-1])
	} else {
		fmt.Println("Stack is empty, cannot peek.")
	}
}



// REPL (READ-EVALUATE-PRINT-LOOP)
func Load() {
	var input string
	s := &stack{}
	for {
		fmt.Print("RUN=> ")
		fmt.Scan(&input)
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
