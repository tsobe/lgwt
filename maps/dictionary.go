package maps

const (
	ErrNotFound          = DictionaryErr("not found")
	ErrWordAlreadyExists = DictionaryErr("word already exists")
)

type DictionaryErr string
type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = newDefinition
	}
	return err
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (err DictionaryErr) Error() string {
	return string(err)
}
