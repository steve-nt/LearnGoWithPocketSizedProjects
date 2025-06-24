package bookworm

import (
	"encoding/json"
	"os"
)
// Define a Book type first
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// Updated Bookworm struct to match your JSON
type Bookworm struct {
	Name  string  `json:"name"`
	Books []Book  `json:"books"` // Now using []Book instead of []string
}

//loadBookworms reads the file and returns the list of bookworms,
//and their beloved books, found therein
func LoadBookworms(filePath string) ([]Bookworm, error ) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    //Initialize the type in which the file will be decoded.
    var bookworms []Bookworm

    //Decode the file and store the content in the variable bookworms.
    err = json.NewDecoder(f).Decode(&bookworms)
    if err != nil {
        return nil, err
    }
    return bookworms, nil
}