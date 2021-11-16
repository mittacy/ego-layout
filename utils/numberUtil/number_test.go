package numberUtil

import "testing"

func TestRoundFloat(t *testing.T) {
	cases := []struct {
		Name   string
		Num    float64
		Digits int
		Expect float64
	}{
		{"", 1.1, 0, 1},
		{"", 1.1, 1, 1.1},
		{"", 1.15, 1, 1.1},
		{"", 1.16, 1, 1.2},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := RoundFloat(c.Num, c.Digits); ans != c.Expect {
				t.Fatalf("%f keep %d digits, expected %f, but %f got",
					c.Num, c.Digits, c.Expect, ans)
			}
		})
	}
}
