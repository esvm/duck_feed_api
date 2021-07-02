package validation

import (
	"testing"

	"github.com/esvm/duck_feed_api/src/exceptions"
	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Empty     string `validate:"in=[]"`
	OneItem   string `validate:"omitempty,in=[abc]"`
	Multiple  string `validate:"omitempty,in=[abc;def;ghi]"`
	Malformed string `validate:"omitempty,in=["`
}

func TestEmptyInValidator(t *testing.T) {
	f := Foo{
		Empty: "Back In Black",
	}

	err := Validate(&f)
	assert.NotNil(t, err, "The field value is not one of the possible values defined on the tag.")
	assert.IsType(t, exceptions.ValidationError{}, err, "Err should be a ValidationError.")
}

func TestOneItemInValidator(t *testing.T) {
	f := Foo{
		OneItem: "Back In Black",
	}

	err := Validate(&f)
	assert.NotNil(t, err, "The field value is not one of the possible values defined on the tag.")
	assert.IsType(t, exceptions.ValidationError{}, err, "Err should be a ValidationError.")
}

func TestOnePresentItemInValidator(t *testing.T) {
	f := Foo{
		OneItem: "abc",
	}

	err := Validate(&f)
	assert.Nil(t, err, "The field value is one of possible values defined on the tag.")
}

func TestMultipleInValidator(t *testing.T) {
	f := Foo{
		Multiple: "Ramble on",
	}

	err := Validate(&f)
	assert.NotNil(t, err, "The field value is not one of possible values defined on the tag.")
	assert.IsType(t, exceptions.ValidationError{}, err, "Err should be a ValidationError.")
}

func TestMultiplePresentInValidator(t *testing.T) {
	f := Foo{
		Multiple: "abc",
	}

	err := Validate(&f)
	assert.Nil(t, err, "The field value is one of possible values defined on the tag.")
}

func TestPanicInMalformed(t *testing.T) {
	f := Foo{
		Malformed: "abc",
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Validator should panic on malformed tags.")
		}
	}()

	Validate(&f)
}
