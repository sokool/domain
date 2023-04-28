package domain_test

import (
	"testing"
)

func TestURL_UnmarshalJSON(t *testing.T) {
	//var u domain.URL
	//assert.Error(t, u.UnmarshalJSON([]byte(`""`)))
	//assert.Error(t, u.UnmarshalJSON([]byte(`"       asdf r"`)))
	//assert.Error(t, u.UnmarshalJSON([]byte(`"rollee"`)))
	//
	//assert.NoError(t, u.UnmarshalJSON(nil))
	//assert.NoError(t, u.UnmarshalJSON([]byte(``)))
	//assert.NoError(t, u.UnmarshalJSON([]byte(`null`)))
	//assert.NoError(t, u.UnmarshalJSON([]byte(`"https://greg:j6G1Df7@rollee.io:8888/documents/january-invoice.pdf?version=3"`)))
	//
	//assert.Equal(t, u.Schema, "https")
	//assert.Equal(t, u.Username, "greg")
	//assert.Equal(t, u.Password, "j6G1Df7")
	//assert.Equal(t, u.Host, "rollee.io")
	//assert.Equal(t, u.Port, "8888")
	//assert.Equal(t, u.Path, []string{"documents", "january-invoice.pdf"})
	//assert.Equal(t, u.Query, url.Values{"version": {"3"}})
}

func TestURL_IsZero(t *testing.T) {
	//var u domain.URL
	//assert.Equal(t, u.IsZero(), true)
}
