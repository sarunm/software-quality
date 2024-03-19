package double

import "errors"

var (
	ErrMissingArgs   = errors.New("FirstName and LastName are required")
	ErrNoPersonFound = errors.New("no person found")
)

type Queryer interface {
	Search(people []*Person, firstName, lastName string) *Person
	// Update(people []*Person, firstName, lastName string, phone string) *Person
}

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}
type PhoneBook struct {
	People []*Person
}

func (p *PhoneBook) Find(queryer Queryer, firstName, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArgs
	}

	person := queryer.Search(p.People, firstName, lastName)
	if person == nil {
		return "", ErrNoPersonFound

	}

	// IsUpdate := queryer.Update(p.People, firstName, lastName, "1234567890")
	// if IsUpdate == nil {
	// 	return "", ErrNoPersonFound
	// }
	return person.Phone, nil
}
