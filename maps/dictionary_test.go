package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := WordNotFoundError
		assertNotNil(t, err)
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dicionary := Dictionary{}
		errAdd := dicionary.Add("test", "this is just a test")
		assertNil(t, errAdd)

		want := "this is just a test"
		got, err := dicionary.Search("test")

		assertNil(t, err)
		assertString(t, got, want)

	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dicionary := Dictionary{word: definition}
		err := dicionary.Add(word, definition)
		want := AlreadyExistsError

		assertNotNil(t, err)
		assertError(t, err, want)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	originalDefinition := "original definition"
	updatedDefinition := "updated definition"
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: originalDefinition}
		errUpdt := dictionary.Update(word, updatedDefinition)
		assertNil(t, errUpdt)

		got, err := dictionary.Search(word)
		assertNil(t, err)
		assertString(t, got, updatedDefinition)
	})
	t.Run("non-existent word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update(word, updatedDefinition)

		assertNotNil(t, err)
		assertError(t, err, WordNotFoundUpdateError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is just a test"}
		err := dictionary.Delete(word)

		assertNil(t, err)
	})

	t.Run("non-existent word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)

		assertNotNil(t, err)
		assertError(t, err, WordNotFoundDeleteError)
	})
}

func assertString(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given: %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given: %q", got, want, "test")
	}
}

func assertNotNil(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("should find added word: ", err)
	}

}

func assertNil(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("did not exptected to get an error but got %q", err)
	}
}
