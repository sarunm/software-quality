package double

import "testing"

type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone}
}

func TestFindReturnsPerson(t *testing.T) {
	fakePhone := "1234567890"
	phoneBook := &PhoneBook{}
	ss := StubSearcher{phone: fakePhone}

	got, _ := phoneBook.Find(ss, "John", "Doe")
	if got != fakePhone {
		t.Errorf("got %v want %v", got, fakePhone)
	}
}
