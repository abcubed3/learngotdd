package main

import "fmt"

// Error constants for Error when key is not found
const (
	ErrNotFound         = DictionaryErr("could not find the word you are looking for")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exists")
)

// Dictionary struct
type Dictionary map[string]string

// DictionaryErr type
type DictionaryErr string

// Search Function to search within a dictionary stuct
func (d Dictionary) Search(query string) (string, error) {
	definition, ok := d[query]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add Method to add a word to a dictionary stuct
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update Method to change the definition a word to a dictionary stuct
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

// Delete Method to add a word to a dictionary stuct
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

// Dictionary helper error function that implements the Error() function
func (e DictionaryErr) Error() string {
	return string(e) //Cast the error into a string
}
func main() {
	dictionary := Dictionary{"test": "this is just a test"}
	query := "test"
	result, err := dictionary.Search(query)
	if err == nil {
		fmt.Printf("Searched %q, Found %q", query, result)
	} else {
		fmt.Printf("Searched %q, but %v", query, err)
	}

}
