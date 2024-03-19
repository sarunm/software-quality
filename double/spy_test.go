package double

import "testing"

type SpySearcher struct {
	phone           string
	searchWasCalled bool
	whatIsInputName string
}

func (ss *SpySearcher) Search(people []*Person, firstName, lastName string) *Person {
	ss.searchWasCalled = true
	ss.whatIsInputName = firstName
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}
func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "1234567890"
	phoneBook := &PhoneBook{}
	spy := &SpySearcher{phone: fakePhone}

	got, _ := phoneBook.Find(spy, "John", "Doe")

	if !spy.searchWasCalled {
		t.Errorf("expected search to be called")
	}

	if spy.whatIsInputName != "John" {
		t.Errorf("got %v want %v", spy.whatIsInputName, "John")
	}

	if got != fakePhone {
		t.Errorf("got %v want %v", got, fakePhone)
	}
}
