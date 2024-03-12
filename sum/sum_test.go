package sum

import "testing"

func TestSum(t *testing.T) {

	t.Run("should return 3 when input 1 and 2", func(t *testing.T) {

		want := 3

		got := sum(1, 2)
		if got != want {
			t.Errorf("sum(1, 2) = %d; want %d", got, want)
		}
	})

	t.Run("should return 1 when input 1 and 0", func(t *testing.T) {
		want := 1

		got := sum(1, 0)
		if got != want {
			t.Errorf("sum(1, 0) = %d; want %d", got, want)
		}
	})

	t.Run("should return -2 when input -1 and -1", func(t *testing.T) {
		want := -2

		got := sum(-1, -1)

		if got != want {
			t.Errorf("sum(-1, -1) = %d; want %d", got, want)
		}
	})

}

// func TestSumShouldReturn1WhenInput1and0(t *testing.T) {
// 	got := sum(1, 0)
// 	if got != 1 {
// 		t.Errorf("sum(1, 0) = %d; want 1", got)
// 	}
// }

// func TestSumShouldReturnMinus2WhenInputMinus1andMinus1(t *testing.T) {
// 	got := sum(-1, -1)
// 	if got != -2 {
// 		t.Errorf("sum(-1, -1) = %d; want -2", got)
// 	}
// }
