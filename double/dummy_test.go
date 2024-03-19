package double

import "testing"

type DummySearcher struct{}

func (ds DummySearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{}
}
func (ds DummySearcher) Update(people []*Person, firstName, lastName, phone string) *Person {
	return &Person{}
}

func TestFindItShouldReturnErrorWhenFirstNameOrLastNameIsEmpty(t *testing.T) {
	phoneBook := &PhoneBook{}
	want := ErrMissingArgs

	dd := DummySearcher{}

	_, got := phoneBook.Find(dd, "", "")
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
