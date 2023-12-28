package doc

import(
	"fmt"
	"log"
	"os"
)

// Documentation & File Loading REPL
func Load() {
	var state int
	var input string 
	fmt.Print("DOC=>")
	state = 1
	fmt.Scan(&input)
	for input != "exit" {
		if input == "exit" {
			log.Println("Exiting DOC REPL")
		} else if input == "test" {
			content, err := os.ReadFile("examples/test.txt")
			if err != nil {
				log.Fatal(err)
			}
			str := string(content)
			fmt.Println(str)
			break
		} else {
			fmt.Println("Invalid Argument Entered!")
			state = 0
			break
		}
	}
	if state != 1 {
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
