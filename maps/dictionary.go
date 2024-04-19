package main

var (
	AlreadyExistsError      = DictionaryError("word already exists in dictionary")
	WordNotFoundError       = DictionaryError("cound not find the word you were looking for")
	WordNotFoundUpdateError = DictionaryError("cannot update word because it does not exist")
	WordNotFoundDeleteError = DictionaryError("cannot delete word because it does not exist")
)

type Dictionary map[string]string

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", WordNotFoundError
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	if d[word] != "" {
		return AlreadyExistsError
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	if d[word] != "" {
		d[word] = definition
		return nil
	}

	return WordNotFoundUpdateError
}

func (d Dictionary) Delete(word string) error {
	if d[word] != "" {
		d[word] = ""
		return nil
	}
	return WordNotFoundDeleteError
}
