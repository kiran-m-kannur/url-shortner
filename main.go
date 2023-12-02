package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Hash map that maps all links to a randomly generated string.
// Eg : https://youtube.com/kirankannur -> Ab63s8Uy
var linkMap = make(map[string]string)

// pool of characters where the random string is generated from
const letters = "abacdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQURSTUVWXYZ0123456789"

// Generates Random string of input length
func generateRandomString(len int) string {

	randomBytes := randASCIIBytes(len)
	return string(randomBytes)
}

func randASCIIBytes(n int) []byte {
	output := make([]byte, n)
	randomness := make([]byte, n)
	_, err := rand.Read(randomness)
	if err != nil {
		fmt.Printf("error while performing Rand : %v", err)
	}
	l := len(letters)
	for pos := range output {
		random := uint8(randomness[pos])
		randomPos := random % uint8(l)
		output[pos] = letters[randomPos]
	}
	return output
}

// CheckShort func is to check for already existing map.
// if a map already does not exist, it creates a new map.
// 5 int has 62^5 possibilities.Meaning, 916132832 distinct links can be shortened.
func CheckShort(link string) string {
	short, ok := linkMap[link]
	if ok {
		fmt.Println("Found a short link :", short)
		return short
	} else {
		shortlink := generateRandomString(5)
		linkMap[link] = shortlink
		fmt.Println("Couldn't find a short link, generated a new one :", shortlink)
		return shortlink
	}
}

//getLonglink gets the original link from the Map.

func getLongLink(shortlink string, linklist map[string]string) (string, error) {

	for long, short := range linklist {
		if short == shortlink {
			fmt.Println("Got the original Link for the shortlink  :", long)
			return long, nil
		}
	}
	return "", fmt.Errorf("Could not find long link for the respective %s url", shortlink)
}

// Something redundant, used to analyze the hashmap
func prettyPrintMap(m map[string]string) {
	fmt.Println("Updated hash map :")
	for key, value := range m {
		fmt.Printf("%s : %s \n", key, value)
	}
}

// The program works at port 8080
const port string = ":8080"

// "/" handles the task of giving short link
func main() {

	fmt.Println("Initialzing Router ..")

	router := mux.NewRouter()

	router.HandleFunc("/", handleHome)
	router.HandleFunc("/about", handleAbout)
	router.HandleFunc("/short/{id}", handleRedirect)

	fmt.Printf("Listening at http://localhost:%s \n", port[1:])
	http.ListenAndServe(port, router)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, "Kiran")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	link := r.FormValue("url")
	if len(link) < 5 {
		fmt.Fprintf(w, "Please enter a valid Link")
	} else {
		fmt.Printf("User entered link : %s \n", link)
		shortLink := CheckShort(link)
		fmt.Fprintf(w, "Your shortened link : localhost:8080/short/%s \n", shortLink)
		prettyPrintMap(linkMap)
		fmt.Println("")
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	value := strings.TrimPrefix(r.URL.Path, "/short/")
	redirectLink, err := getLongLink(value, linkMap)
	if err != nil {
		fmt.Fprintf(w, "The url you are trying to access does not exist")
	}
	fmt.Println("redirectlink :", redirectLink)
	http.Redirect(w, r, redirectLink, 301)

}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/template/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, "Kiran Kannur")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
