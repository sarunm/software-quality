package ticket

import "testing"

func TestTicketPrice(t *testing.T) {

	tests := []struct {
		name string
		age  int
		want float64
	}{
		{name: "should return 0 when age is 0", age: 0, want: 0.0},
		{name: "should return 0 when age under 3", age: 3, want: 0.0},
		{name: "should return 15 when age is 4", age: 4, want: 15.0},
		{name: "should return 15 when age under 15", age: 15, want: 15.0},
		{name: "should return 30 when age is 16", age: 16, want: 30.0},
		{name: "should return 30 when age under 50", age: 50, want: 30.0},
		{name: "should return 5 when age more than 50", age: 51, want: 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)
			if got != tt.want {
				t.Errorf("Price(%d) = %f; want %f", tt.age, got, tt.want)
			}
		})
	}

	// t.Run("should return 0 when age is 0", func(t *testing.T) {
	// 	want := 0.0
	// 	age := 0

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(%d) = %f; want %f", age, got, want)
	// 	}
	// })

	// t.Run("should return 0 when age under 3", func(t *testing.T) {
	// 	want := 0.0
	// 	age := 3

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(%d) = %f; want %f", age, got, want)
	// 	}
	// })

	// t.Run("should return 15 when age is 4", func(t *testing.T) {
	// 	want := 15.0
	// 	age := 4

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(%d) = %f; want %f", age, got, want)
	// 	}
	// })

	// t.Run("should return 15 when age under 15", func(t *testing.T) {
	// 	want := 15.0
	// 	age := 15

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(%d) = %f; want %f", age, got, want)
	// 	}
	// })

	// t.Run("should return 30 when age is 16", func(t *testing.T) {
	// 	want := 30.0
	// 	age := 16

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(%d) = %f; want %f", age, got, want)
	// 	}
	// })

	// t.Run("should return 30 when age under 50", func(t *testing.T) {
	// 	want := 30.0
	// 	age := 50

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(%d) = %f; want %f", age, got, want)
	// 	}
	// })

	// t.Run("should return 5 when age more than 50", func(t *testing.T) {
	// 	want := 5.0
	// 	age := 51

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(%d) = %f; want %f", age, got, want)
	// 	}
	// })
}
