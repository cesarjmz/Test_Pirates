package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	value, foundValue := os.LookupEnv(key)
	if foundValue {
		return value
	}
	return fallback
}

func getNewGifURL() string {
	res, err := http.Get("http://giphy-proxy.solutions.656.mba/")

	// check for response error
	if err != nil {
		log.Fatal(err)
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// close response body
	res.Body.Close()
	return string(data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// myGifURL := "https://media.giphy.com/media/MBaeLRcqYNyk5p1zu0/giphy.gif"
	myGifURL := getNewGifURL()
	myFunHTML := `
	    <body style ="background-color: #ex7063">
	    <html>
	    <h1>Hello World</h1>
            <div>
                
                <img src=" ` + myGifURL + `">
                <form action="/action_page.php">
                      <label for="fname">First name:</label><br>
                      <input type="text" id="fname" name="fname" value="John"><br>
                      <label for="lname">Last name:</label><br>
                      <input type="text" id="lname" name="lname" value="Doe"><br><br>
                      <input type="submit" value="Submit">
                </form>
                
                <form action="">
                    <button type="button">display gif</button>
                </form>
            </div>
        </html>   
	`
	fmt.Fprintf(w, myFunHTML)
	// fmt.Fprintf(w, "Hello World")

}

func main() {
	// Say that when we receive a request for the '/' (or "root") URL
	// we want the function `indexHandler` to handle it.
	http.HandleFunc("/", indexHandler)
	// Start listening for HTTP requests.
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}
