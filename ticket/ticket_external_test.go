package ticket_test

import (
	"testing"

	ticket "github.com/sarunm/software-quality"
)

func TestTicketPrice(t *testing.T) {
	t.Run("should return 0 when age is 0", func(t *testing.T) {
		want := 0.0
		age := 0

		got := ticket.Price(age)

		if got != want {
			t.Errorf("Price(%d) = %f; want %f", age, got, want)
		}
	})
}
