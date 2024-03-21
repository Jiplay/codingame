package temperature

import "testing"

type Test struct {
	in  []string
	out string
}

var tests = []Test{
	{[]string{"8", "22", "-33", "12", "-25", "-3", "-24", "-1"}, "-1"},
	{[]string{"-8", "-22", "-33", "-12", "-25", "-3", "-24", "-1"}, "-1"},
	{[]string{"-5", "40", "-2", "-12", "44", "-3", "-24", "-1"}, "-1"},
	{[]string{"-5", "40", "-2", "-12", "44", "-3", "-24", "-1", "1"}, "1"},
	{[]string{"-5", "40", "-4", "-12", "44", "-3", "-24", "-3", "3"}, "3"},
	{[]string{}, "0"},
}

func TestFindClosestTo0(t *testing.T) {
	for i, test := range tests {
		closestTo0 := FindClosestTo0(test.in)
		if closestTo0 != test.out {
			t.Errorf("#%d: Closest(%s)=%s; expecting %s", i, test.in, closestTo0, test.out)
		}
	}
}

/*

* go test; permet de lancer les tests
* go test -coverprofile=coverage.out; permet de créer coverage.out
* go tool cover -html=coverage.out; permet de visualiser le coverage

 */
