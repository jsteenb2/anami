package anami_test

import (
	"io/ioutil"
	"testing"
)

func goldenHelper(t *testing.T, filepath string, b []byte) []byte {
	if *update {
		if err := ioutil.WriteFile(filepath, b, 0644); err != nil {
			t.Fatal(err)
		}
	}

	expected, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatal(err)
	}

	return expected
}
