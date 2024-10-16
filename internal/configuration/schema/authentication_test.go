package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticationBackendExtraAttribute(t *testing.T) {
	testCases := []struct {
		name  string
		have  AuthenticationBackendExtraAttribute
		vtype string
		mv    bool
	}{
		{
			"ShouldReturnDefaults",
			AuthenticationBackendExtraAttribute{},
			"",
			false,
		},
		{
			"ShouldReturnDefaults",
			AuthenticationBackendExtraAttribute{
				ValueType:   "string",
				MultiValued: true,
			},
			"string",
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.vtype, tc.have.GetValueType())
			assert.Equal(t, tc.mv, tc.have.IsMultiValued())
		})
	}
}

func TestAuthenticationBackendLDAPAttributesAttribute(t *testing.T) {
	testCases := []struct {
		name  string
		have  AuthenticationBackendLDAPAttributesAttribute
		vtype string
		mv    bool
	}{
		{
			"ShouldReturnDefaults",
			AuthenticationBackendLDAPAttributesAttribute{},
			"",
			false,
		},
		{
			"ShouldReturnDefaults",
			AuthenticationBackendLDAPAttributesAttribute{
				AuthenticationBackendExtraAttribute: AuthenticationBackendExtraAttribute{
					MultiValued: true,
					ValueType:   "integer",
				},
			},
			"integer",
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.vtype, tc.have.GetValueType())
			assert.Equal(t, tc.mv, tc.have.IsMultiValued())
		})
	}
}
