package doc

import(
	"fmt"
	"log"
	"ioutil"
)

type doc struct {}

func (d *doc) repl_load(){
	var input string
	for input != "exit"{
		if input == "exit"{
			log.Println("Exiting DOC REPL")
			break;
		}
		else if input "std"{
			content, err := ioutil.ReadFile("std.txt")

			if err != nil{
				log.Fatal(err)
			}
			str := string(content)
			fmt.Println(str)
		}
		else{
			fmt.Println("Invalid Argument Entered!")
		}

	}
}

func (d *doc) doc_handler(file string){
	// reads file and returns related documentation
	// used for the case of multiple arguments
}
