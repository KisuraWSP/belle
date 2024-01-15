package document

/*
	IMPORTS
*/
import(
	"fmt"
	"log"
	"os"
)

/*
	HELPER FUNCTIONS	
*/
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


/*
	Documentation & File Loading REPL
*/
func Load() {
	var input string 
	for {
		fmt.Print("DOC=>")
		fmt.Scanln(&input)
		if input == "exit" {
			log.Println("Exiting DOC REPL")
			break
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
		} else if input == "declare" {
			GetDoc_belle("declare")
		} else {
			fmt.Println("Invalid Argument Entered!")
		}
	}	
}
