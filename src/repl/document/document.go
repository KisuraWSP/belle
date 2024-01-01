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
	var state int
	var input string 
	state = 1
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
		} else {
			fmt.Println("Invalid Argument Entered!")
			state = 0
			break
		}
	}
	if state != 1 && input != "exit"{
		fmt.Print("Do You want to enter again (Y/n): ")
		fmt.Scan(&input)
		if input == "y" || input == "Y"{
			Load()
		}
	}
}

func  Handler(file string) {
	// reads file and returns related documentation
	// used for the case of multiple arguments
}
