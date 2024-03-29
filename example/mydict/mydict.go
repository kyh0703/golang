package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("not Found")
	errCantUpdate = errors.New("Cant update non-exiting word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

<<<<<<< HEAD
=======
// Update a word to the dictionary
>>>>>>> 753cd406a8b860cf1d1ec9777313b5132652b062
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
<<<<<<< HEAD
		return err
	}
	return nil
}
=======
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return errNotFound
	}
	delete(d, word)
	return nil
}
>>>>>>> 753cd406a8b860cf1d1ec9777313b5132652b062
