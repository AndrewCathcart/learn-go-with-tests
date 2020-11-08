package maps

// Dictionary represents a map of string to string
type Dictionary map[string]string

const (
	// ErrNotFound is returned if the key does not exist in the map
	ErrNotFound = DictionaryErr("word not found")
	// ErrWordExists is returned if the key already exists in the map
	ErrWordExists = DictionaryErr("this word already exists")
	// ErrWordDoesNotExist is an explicit error to indicate an update failed
	ErrWordDoesNotExist = DictionaryErr("word does not exist and could not be updated")
)

// DictionaryErr represents errors that the dictionary can return
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search tries to find the definition of the given word in the given dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add attemps to add a word and a definition to the dictionary
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

// Update attempts to update the definition of a word in the dictionary
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

// Delete removes a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
