package main

import (
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"html/template"
)

func main() {
	http.HandleFunc("/libraryList", homepage)
	http.HandleFunc("/libraryList/searchBooks", searchBooks)
	// Set listen port and serve it.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}

func homepage(res http.ResponseWriter, req *http.Request) {

	// Sets json file data to @LibraryJSON struct.
	jsonResult := readJSONFile()

	// Setting home page html info.
	// Home page html renderer.
	if req.Method == "GET" {
		t, _ := template.ParseFiles("./templates/bookList.gtpl")
		t.Execute(res, jsonResult.Books)
	} else {
		// Form POST action for submit.
		// Parsing requested arguments.
		req.ParseForm()
		redirect_URL := "http://www.google.com/search?q=" + req.Form.Get("bookName")
		http.Redirect(res, req, redirect_URL, 301)
	}
}

/**
 * Redirecting page for Google search.
 */
func searchBooks(res http.ResponseWriter, req *http.Request) {
	// Parsing requested arguments.
	req.ParseForm()
	redirect_URL := "http://www.google.com/search?q=" + req.Form.Get("bookName")
	http.Redirect(res, req, redirect_URL, 301)
}

/**
 * Library struct for matching up with JSON data.
 */
type LibraryJSON struct {
	Books []struct {
		Author    string   `json:"author"`
		GenreS    string   `json:"genre_s"`
		ID        string   `json:"id"`
		InStock   bool     `json:"inStock"`
		Name      string   `json:"name"`
		PagesI    int      `json:"pages_i"`
		Price     int      `json:"price"`
		SequenceI int      `json:"sequence_i"`
		SeriesT   string   `json:"series_t"`
	} `json:"books"`
}

/**
 * Returns LibraryJSON data for defined input.
 */
func readJSONFile() (libraryJSON LibraryJSON) {

	// Reading library json as mock data.
	file, err := ioutil.ReadFile("./json-data/library.json")
	if err != nil {
		log.Fatal("File error : ", err)
	}

	json.Unmarshal(file, &libraryJSON)
	return
}