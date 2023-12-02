package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var linkMap = make(map[string]string)

const letters = "abacdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQURSTUVWXYZ0123456789"

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

func CheckShort(link string) string {
	short, ok := linkMap[link]
	if ok {
		fmt.Println("Found a short link :", short)
		return short
	} else {
		shortlink := generateRandomString(8)
		linkMap[link] = shortlink
		fmt.Println("Couldn't find a short link, generated a new one :", shortlink)
		return shortlink
	}
}

func getLongLink(shortlink string, linklist map[string]string) (string, error) {

	for long, short := range linklist {
		if short == shortlink {
			fmt.Println("Got the original Link for the shortlink  :", long)
			return long, nil
		}
	}
	return "", fmt.Errorf("Could not find long link for the respective %s url", shortlink)
}

func prettyPrintMap(m map[string]string) {
	fmt.Println("Updated hash map :")
	for key, value := range m {
		fmt.Printf("%s : %s \n", key, value)
	}
}

var name = generateRandomString(10)

const port string = ":8080"

func main() {

	fmt.Println("Initialzing Router ..")

	router := mux.NewRouter()

	router.HandleFunc("/", handleHome)
	router.HandleFunc("/about", handleAbout)

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
		fmt.Fprintf(w, "Your shortened link : http://localhost:8080/%s \n", shortLink)
		prettyPrintMap(linkMap)
		fmt.Println("")
	}
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
