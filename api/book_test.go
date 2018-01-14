package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookSerializatioToJson(t *testing.T) {
	book := Book{Title: "Foo", Author: "Bar", ISBN: "888"}
	json := book.toJSON()

	assert.Equal(t, `{"title":"Foo","author":"Bar","isbn":"888"}`, string(json), "Assertion failed")
}

func TestBookDeserializatioToJson(t *testing.T) {
	expected := Book{Title: "Foo", Author: "Bar", ISBN: "888"}
	json := []byte(`{"title":"Foo","author":"Bar","isbn":"888"}`)

	actual := FromJson(json)

	assert.Equal(t, actual, expected, "Assertion failed")
}
