package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		query := "test"
		got, _ := dictionary.Search(query)
		want := "this is just a test"

		assertSearch(t, got, want, query)
	})

	t.Run("unknown word", func(t *testing.T) {
		query := "unknown"
		_, got := dictionary.Search(query)
		assertError(t, got, ErrNotFound, query)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertAddError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		// dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertAddError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		assertAddError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update word not found", func(t *testing.T) {
		// dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertAddError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)
		_, err := dictionary.Search(word)

		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})

}
func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}

}

func assertSearch(t *testing.T, got, want, query string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, query)
	}
}

func assertError(t *testing.T, got, want error, query string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, query)
	}
}

func assertAddError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}
