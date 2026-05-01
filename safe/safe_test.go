package safe

import "testing"

func TestMustAt(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("Ожидали панику, но её не было")
		} else if err != "index out of range" {
			t.Errorf("Неожиданное сообщение паники: %v", err)
		}
	}()

	_ = MustAt([]int{1, 2, 3}, 55)
}
