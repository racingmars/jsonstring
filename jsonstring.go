package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	inputstr := string(input)
	output, err := json.Marshal(inputstr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
