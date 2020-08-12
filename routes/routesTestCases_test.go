package routes

import (
	"testing"
)

type Browsers struct {
	Chrome  string
	FireFox string
	IE      string
}

var browsers Browsers

func TestHomeSuite(t *testing.T) {
	testCases := []struct {
		Browsers string
	}{
		{Browsers: "browsers.FireFox"}, //as for demo fr showig in logs 'Failed on Browser.Type
		{Browsers: "browsers.Chrome"},
		{Browsers: "browsers.IE"},
	}

	for _, tC := range testCases {
		t.Run(tC.Browsers, func(t *testing.T) {
			TestHome200(t)
			TestHome404(t)

		})
	}
}

func TestAllSuites(t *testing.T) {
	testCases := []struct {
		Suites string
	}{
		{Suites: "Descript"},
	}

	for _, tC := range testCases {

		t.Run(tC.Suites, func(t *testing.T) {
			TestHomeSuite(t)

		})
	}
}
