package double

import (
	"testing"
)

type FakeSearcher struct{}

func (fs FakeSearcher) Search(people []*Person, firstName, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}

	return people[0]
}

func TestFindCallSearchAndReturnEmptyStringForNoPerson(t *testing.T) {

	phoneBook := &PhoneBook{}
	fake := FakeSearcher{}

	got, _ := phoneBook.Find(fake, "John", "Doe")

	if got != "" {
		t.Errorf("got %v want %v", got, "")
	}
}
