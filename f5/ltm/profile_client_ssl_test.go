package ltm

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TMOptions_MarshallJSON(t *testing.T) {
	opts := TMOptions{"a", "b", "c"}
	bytes, err := json.Marshal(opts)
	assert.NoError(t, err)
	str := string(bytes)
	assert.Equal(t, "[\"a\",\"b\",\"c\"]", str)
}

func Test_TMOptions_UnmarshallJSON(t *testing.T) {
	t.Parallel()

	var jsonTests = []struct {
		name  string
		input string
	}{
		{"space inside braces", "{ dont-insert-empty-fragments no-tlsv1.3 }"},
		{"no spaces", "{dont-insert-empty-fragments no-tlsv1.3}"},
		{"space outside and inside braces", " { dont-insert-empty-fragments no-tlsv1.3 } "},
		{"space outside brace", " {dont-insert-empty-fragments no-tlsv1.3} "},
		{"\"quotation marks", "\"{dont-insert-empty-fragments no-tlsv1.3}\""},
	}

	for _, tc := range jsonTests {
		t.Run(tc.name, func(t *testing.T) {
			opts := &TMOptions{}
			err := opts.UnmarshalJSON([]byte(tc.input))
			assert.NoError(t, err)
			assert.Equal(t, &TMOptions{"dont-insert-empty-fragments", "no-tlsv1.3"}, opts)
		})
	}

}
