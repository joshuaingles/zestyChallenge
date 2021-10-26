-- Project Info --
- Author: Joshua Ingles
- Built: 23Jul20
- Stack:
-- Go

-- About --
This is my work for the ZestyChallenge ( https://github.com/zesty-io/gists/tree/master/code-challenge-autocomplete ).
The challenge was to create an autocomplete web service in Go that used the entire works of Shakespeare as a reference.
The service takes in a term ( i.e. "th" or "fr" ) and finds the top 25 matches based on word frequency.
If less than 25 words match the term, then it shows however many words match the term ordered from highest to lowest occurance of the word.


-- Instructions -- 
- Download web service from GitHub
-- URL: https://github.com/joshuaingles/zestyChallenge

- Insure that port 9000 is clear or change port number in server.go -> main() -> http.ListenAndServe()

- Open a terminal and navigate to the zestyChallenge directory

- run server.go

- Open a browser and go to http://localhost:{PORT NUMBER}/autocomplete?term={TERM} or use below cURL example
-- Replace {PORT NUMBER} if changed in server.go and replace {TERM} with desired search term

- Results will be displayed in the browser window and output to the server.go terminal.


-- cURL Example ---

curl http://localhost:9000/autocomplete?term=pl


-- References --

// Structure and baseline functionality of web server :
// https://blog.logrocket.com/creating-a-web-server-with-golang/

// Extraction of URL parameters:
// https://stackoverflow.com/questions/15407719/in-gos-http-package-how-do-i-get-the-query-string-on-a-post-request

// Opening and reading a file:
// https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file-using-go
