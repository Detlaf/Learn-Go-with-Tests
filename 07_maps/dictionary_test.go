package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("something")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("add a new word", func(t *testing.T) {
		word := "example"
		definition := "this is a definition"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "example"
		definition := "this is a definition"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "example"
	definition := "some definition"
	dictionary := Dictionary{word: definition}

	t.Run("update a definition for an existing word", func(t *testing.T) {
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		if err != nil {
			t.Fatal("didn't expect an error")
		}
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update a definition for a word not in the dictionary", func(t *testing.T) {
		err := dictionary.Update("novelty", "new")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "example"
	definition := "some definition"
	dictionary := Dictionary{word: definition}

	t.Run("Delete an existing word from the dictionary", func(t *testing.T) {
		err := dictionary.Delete(word)

		if err != nil {
			t.Fatal("didn't expect to fail")
		}

		_, searchError := dictionary.Search(word)
		assertError(t, searchError, ErrNotFound)
	})

	t.Run("Delete a word not in the dictionary", func(t *testing.T) {
		err := dictionary.Delete("who")

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}
