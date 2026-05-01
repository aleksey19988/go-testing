package validate

import (
	"errors"
	"testing"
)

func TestValidateName(t *testing.T) {
	err := ValidateName("")
	if !errors.Is(err, ErrEmptyName) {
		t.Errorf("ValidateName() error = %v, wantErr %v", err, ErrEmptyName)
	}

	err = ValidateName("test")
	if err != nil {
		t.Errorf("ValidateName() error = %v, wantErr %v", err, nil)
	}
}
