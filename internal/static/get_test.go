package static_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/ecosystem/forma/internal/static"
)

func TestLoadTemplate(t *testing.T) {
	tests := []struct {
		name      string
		base, tpl string
		golden    string
	}{
		{"error template", "./templates", "error.html", "./templates/error.html"},
		{"error template, bindata", "static/templates", "error.html", "./templates/error.html"},
		{"redirect template", "./templates", "redirect.html", "./templates/redirect.html"},
		{"redirect template, bindata", "static/templates", "redirect.html", "./templates/redirect.html"},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			obtained, err := LoadTemplate(tc.base, tc.tpl)
			assert.NoError(t, err)
			assert.Equal(t, expected, obtained)
		})
	}
}
