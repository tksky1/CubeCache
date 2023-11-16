package cache

import (
	"testing"
)

func TestMapper(t *testing.T) {
	hash := NewMapper(3)

	hash.AddNode("6fd", "4dsa", "22")

	testCases := map[string]string{
		"23":   "4dsa",
		"11":   "22",
		"2dd3": "6fd",
		"27":   "6fd",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, hash.Get(k))
		}
	}

	hash.AddNode("814fds")

	testCases["27"] = "814fds"
	testCases["2dd3"] = "814fds"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, hash.Get(k))
		}
	}

	hash.RemoveNode("814fds")
	if hash.Get("27") == "814fds" && hash.Get("27") != "6fd" {
		t.Errorf("test removeNode fail")
	}

}
