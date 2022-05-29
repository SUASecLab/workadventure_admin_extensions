package extensions

import (
	"testing"
)

var testdata = []struct {
	name   string
	uuid   string
	result bool
}{
	{"UUID 1", "", false},
	{"UUID 2", "4f436b45-ae03-4dbf-95a2-44e77ce54847", true},
	{"UUID 3", "4f436b45-ae03-4dbf-95a2-44e77ce5484", false},
	{"UUID 4", "4f436b45-ae03-4dbf-95a2-44e77ce548477", false},
	{"UUID 5", "abc", false},
}

func TestIsUUIDValid(t *testing.T) {
	for _, data := range testdata {
		t.Run(data.name, func(t *testing.T) {
			data := data
			t.Parallel()
			result := IsUUIDValid(data.uuid)
			if result != data.result {
				t.Errorf("Error while validating uuid %s: expected %v but received %v", data.uuid, result, data.result)
			}
		})
	}
}
