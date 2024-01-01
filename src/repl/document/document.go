package document

import(
	"fmt"
	"log"
	"os"
)

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
			content, err := os.ReadFile("examples/test.txt")
			if err != nil {
				log.Fatal(err)
			}
			str := string(content)
			fmt.Println(str)
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
