package document

import(
	"fmt"
	"log"
	"os"
)

func GetDoc_txt(file string){
	content, err := os.ReadFile("examples/"+file+".txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(content)
	fmt.Println(str)
}	


func GetDoc_belle(file string){
	content, err := os.ReadFile("examples/"+file+".belle")
	if err != nil {
		log.Fatal(err)
	}
	str := string(content)
	fmt.Println(str)
}	

// Documentation & File Loading REPL
func Load() {
	var input string 
	for {
		fmt.Print("DOC=>")
		fmt.Scanln(&input)
		if input == "exit" {
			log.Println("Exiting DOC REPL")
			break
		} else if input == "test" {
			GetDoc_txt("test")
		} else if input == "the_array" {
			GetDoc_belle("the_array")
		} else if input == "functions" {
			GetDoc_belle("functions")
		} else if input == "procedures" {
			GetDoc_belle("procedures")
		} else if input == "types" {
			GetDoc_belle("types")
		} else if input == "variables" {
			GetDoc_belle("variables")
		} else if input == "condition_handling" {
			GetDoc_belle("condition_handling")
		} else if input == "simple_declare" {
			GetDoc_belle("simple_declare")
		} else {
			fmt.Println("Invalid Argument Entered!")
		}
	}
	
}
