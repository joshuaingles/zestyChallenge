package main

import ( 
	"fmt"
	"log"
	"net/http"
	"bufio"
	"os"
	"strings"
	"regexp"
)

func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
    	if r.URL.Path != "/autocomplete" {
        	http.Error(w, "404 not found.", http.StatusNotFound)
        	return
    	}

    	if r.Method != "GET" {
        	http.Error(w, "Method is not supported.", http.StatusNotFound)
        	return
    	}
	log.Println("- autocompleteHandler()")

	term := r.URL.Query().Get("term")

	if len(term) == 0 {
        	log.Println("!- No term parameter")
        	return
    	}

    	log.Println("-- term: " + term)

	results := parseTerm(term)
	
	for i := 1 ; i <= 25 ; i++ {
		highest := getHighest(results)

		for k, v := range highest {
			if v != 0 {
				value := fmt.Sprintf("%d", v)
				output := "word : " + k + " , count : " + value + "\n"
				fmt.Fprintf(w, output)	
			}	
		}	
	}

	log.Print("- autocompleteHandler() : done")
}

func parseTerm(term string)map[string]int {
	log.Println("- parseTerm()")
	var results map[string]int = make(map[string]int)

	// Readying file for use
	file, err := os.Open("shakespeare-complete.txt")
    	if err != nil {
        	log.Fatal(err)
    	}
    	defer file.Close()

	// Uses scanner to integrate over each line of the text
    	scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
		// Processing of line, removes numbers and special characters, toLower
        	line := scanner.Text()
		reg, _ := regexp.Compile("[^a-zA-Z]")
		line = reg.ReplaceAllString(line, " ")
		line = strings.ToLower(line)

		// Tokenizing of line
		tokens := strings.Split(line, " ")
		
		// Checks if a token's begining characters match term
		// if yes, either adds it to the map or updates count
		for _, element := range tokens { 
			if element != term && len(element) >= len(term) && element[0:len(term)] == term {  
        			_, present := results[element]
				
				if present { 
					results[element] = results[element] + 1
				} else { 
					results[element] = 1
				}
    			}
    		}
    	}

    	if err := scanner.Err(); err != nil {
        	log.Fatal(err)
    	}

	//log.Println(results)

	return results 

}

func getHighest(parsedMap map[string]int)map[string]int {
	log.Println("- getHighest()")
	highestWord := ""
	highestCount := 0
	
	for k, v := range parsedMap {
		if v > highestCount && v != 0 {
			highestWord = k
			highestCount = v	
		}
	}

	highest := map[string]int{highestWord: highestCount}
	log.Print("-- Highest: ")
	log.Println(highest)
	
	delete(parsedMap, highestWord)
	log.Println("-- highest mapping removed")

	return highest
}

func main() {
	http.HandleFunc("/autocomplete", autocompleteHandler)

  	fmt.Printf("- Starting server at port 9000\n")
	if err := http.ListenAndServe(":9000", nil); err != nil {
        	log.Fatal(err)
    	}
}
