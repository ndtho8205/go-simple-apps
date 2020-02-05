package maps

const (
	NotFoundError         = DictionaryError("could not find the word you were looking for")
	WordExistsError       = DictionaryError("cannot add word because it already exists")
	WordDoesNotExistError = DictionaryError("cannot update word because it does not exist")
)

type DictionaryError string

func (error DictionaryError) Error() string {
	return string(error)
}

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", NotFoundError
	}

	return definition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case NotFoundError:
		dictionary[word] = definition
	case nil:
		return WordExistsError
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Update(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case NotFoundError:
		return WordDoesNotExistError
	case nil:
		dictionary[word] = definition
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)
}
