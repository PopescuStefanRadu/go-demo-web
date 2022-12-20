package types_test

import (
	"encoding/json"
	"faceit/pkg/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestOptional(t *testing.T) {
	type jsonObj struct {
		BoolVal types.Option[bool] `json:"bool_val"`
	}

	testData := []struct {
		description    string
		jsonInput      string
		expectedOption types.Option[bool]
	}{
		{
			description:    "missing data",
			jsonInput:      `{}`,
			expectedOption: types.Option[bool]{},
		},
		{
			description: "null data",
			jsonInput:   `{"bool_val":null}`,
			expectedOption: types.Option[bool]{
				IsPresent: true,
				Val:       nil,
			},
		},
		{
			description: "true",
			jsonInput:   `{"bool_val":true}`,
			expectedOption: types.Option[bool]{
				IsPresent: true,
				Val:       types.ToPtr(true),
			},
		},
	}
	for _, td := range testData {
		t.Run(td.description, func(t *testing.T) {
			g := gomega.NewWithT(t)
			parsed := jsonObj{}
			err := json.Unmarshal([]byte(td.jsonInput), &parsed)
			g.Expect(err).To(gomega.BeNil())
			g.Expect(parsed.BoolVal).To(gomega.Equal(td.expectedOption))
		})
	}
}
