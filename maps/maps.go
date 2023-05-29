package main

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound     = DictionaryErr("could not find the word you are looking for")
	ErrWordExists   = DictionaryErr("value exists in key")
	ErrDoesNotExist = DictionaryErr("value does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = value
	case ErrNotFound:
		return ErrDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
