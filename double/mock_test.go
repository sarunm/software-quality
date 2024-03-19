package double

import "testing"

type MockSearcher struct {
	phone        string
	methodToCall map[string]bool
}

func (ms *MockSearcher) Search(people []*Person, firstName, lastName string) *Person {
	ms.methodToCall["Search"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) Create(people []*Person, firstName, lastName, phone string) *Person {
	ms.methodToCall["Create"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
	}
}

func (ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodToCall == nil {
		ms.methodToCall = make(map[string]bool)
	}
	ms.methodToCall[methodName] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {
	for methodName, called := range ms.methodToCall {
		if !called {
			t.Errorf("expected %v to be called, but it wasn't", methodName)
		}
	}
}

func TestFindCallSearchAndReturnPersonUsingMock(t *testing.T) {
	fakePhone := "1234567890"
	PhoneBook := &PhoneBook{}
	mock := &MockSearcher{phone: fakePhone}

	mock.ExpectToCall("Search")
	got, _ := PhoneBook.Find(mock, "John", "Doe")

	if got != fakePhone {
		t.Errorf("got %v want %v", got, fakePhone)
	}

	mock.Verify(t)
}
