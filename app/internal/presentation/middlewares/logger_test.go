package middlewares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldMaskField(t *testing.T) {
	tests := []struct {
		name   string
		field  string
		expect bool
	}{
		{
			name:   "password",
			field:  "password",
			expect: true,
		},
		{
			name:   "email",
			field:  "email",
			expect: true,
		},
		{
			name:   "name",
			field:  "name",
			expect: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expect, shouldMaskField(tt.field))
		})
	}
}

func TestMaskField(t *testing.T) {
	tests := []struct {
		name           string
		fields         map[string]interface{}
		expectedFields map[string]interface{}
	}{
		{
			name: "mask email and password",
			fields: map[string]interface{}{
				"name":     "test",
				"email":    "test@test.co.jp",
				"password": "password",
			},
			expectedFields: map[string]interface{}{
				"name":     "test",
				"email":    "*****",
				"password": "*****",
			},
		},
		{
			name: "mask nested fields",
			fields: map[string]interface{}{
				"customer": map[string]interface{}{
					"name":     "test",
					"email":    "test@test.co.jp",
					"password": "password",
				},
			},
			expectedFields: map[string]interface{}{
				"customer": map[string]interface{}{
					"name":     "test",
					"email":    "*****",
					"password": "*****",
				},
			},
		},
		{
			name:           "body is nil",
			fields:         map[string]interface{}{},
			expectedFields: map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expectedFields, maskFields(tt.fields))
		})
	}
}
