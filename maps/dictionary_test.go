package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	existingWord := "word"
	definition := "That's a String!"
	dictionary := Dictionary{existingWord: definition}

	t.Run("Known word", func(t *testing.T) {
		want := definition

		got, _ := dictionary.Search(existingWord)

		assertEquals(t, got, want)
	})

	t.Run("Known word", func(t *testing.T) {
		notExistingWord := "unknown_word"
		want := ErrNotFound

		_, got := dictionary.Search(notExistingWord)

		assertError(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "word"
		definition := "definition"

		err := dictionary.Add(word, definition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing word", func(t *testing.T) {
		word := "word"
		definition := "definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordAlreadyExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "word"
		definition := "definition"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("New word", func(t *testing.T) {
		word := "word"
		definition := "definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	word := "word"
	definition := "definition"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertEquals(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, expected %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, want string) {
	t.Helper()
	got, err := dictionary.Search(word)
	assertNoError(t, err)
	assertEquals(t, got, want)
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("Expected to get an error %q", want)
	}
	if got != want {
		t.Errorf("Got %q, expected %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("No error expected, got %q", err)
	}
}
