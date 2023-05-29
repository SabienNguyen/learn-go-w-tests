package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("search", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("word not in dict", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		d := Dictionary{key: value}
		err := d.Add(key, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update an existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		d := Dictionary{key: value}
		newValue := "new value"
		d.Update(key, newValue)

		assertDefinition(t, d, key, newValue)
	})
	t.Run("update a key that is not yet added", func(t *testing.T) {
		d := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := d.Update(key, value)

		assertError(t, err, ErrDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	value := "this is just a test"
	d := Dictionary{key: value}
	d.Delete(key)

	_, err := d.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", key)
	}

}

func assertDefinition(t testing.TB, d Dictionary, key, value string) {
	t.Helper()
	got, err := d.Search(key)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, value)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
