package models

import (
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestBaseName(t *testing.T) {
	image := Image{Path: "/foo/bar.jpg"}
	assertEqual(t, "bar.jpg", image.BaseName())
}

func TestUrlPath(t *testing.T) {
	image := Image{Path: "/foo/bar?.jpg"}
	assertEqual(t, "bar%3F.jpg", image.UrlPath())
}
