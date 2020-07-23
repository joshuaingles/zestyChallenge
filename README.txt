-- Instructions -- 
- Download web service from GitHub
-- URL:

- Insure that port 9000 is clear or change port number in server.go -> main() -> http.ListenAndServe()

- Open a terminal and navigate to the zestyChallenge_Ingles folder

- run server.go

- Open a browser and go to http://localhost:{PORT NUMBER}/autocomplete?term={TERM} or use below cURL example
-- Replace {PORT NUMBER} if changed in server.go and replace {TERM} with desired search term

- Results will be displayed in the browser window and output to the server.go terminal

-- cURL Example ---

curl http://localhost:9000/autocomplete?term=pl

-- References --

// Structure and baseline functionality of web server :
// https://blog.logrocket.com/creating-a-web-server-with-golang/

// Extraction of URL parameters:
// https://stackoverflow.com/questions/15407719/in-gos-http-package-how-do-i-get-the-query-string-on-a-post-request

// Opening and reading a file:
// https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file-using-go
