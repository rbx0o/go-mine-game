package domain

import (
	"testing"

	"github.com/google/uuid"
)

/*
Test_NewID
тестирует функцию NewID на наличие возвращаемых ошибок
*/
func Test_NewID(t *testing.T) {
	id, err := NewID()
	if err != nil {
		t.Errorf("NewID return error: %v", err)
	}

	if id == ID(uuid.Nil) {
		t.Errorf("NewID return nil ID")
	}
}

/*
Test_ParseID
тестирует функцию ParseID на корректный парсинг из строки
*/
func Test_ParseID(t *testing.T) {
	has, err := ParseID("123")
	want := ID(uuid.Nil)
	if err == nil {
		t.Errorf("has: %v, want: %v", has, want)
	}

	has, err = ParseID("b6bf1713-d757-4234-a56f-0d0d715d3804")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

/*
Test_String
тестирует функцию String на корректность возвращаемого значения
*/
func Test_String(t *testing.T) {
	want := "b6bf1713-d757-4234-a56f-0d0d715d3804"
	id, _ := ParseID(want)
	has := id.String()

	if want != has {
		t.Errorf("has: %v, want: %v", has, want)
	}
}

/*
Test_IsZero
тестирует функцию IsZero
*/
func Test_IsZero(t *testing.T) {
	nilID := ID(uuid.Nil)
	if !nilID.IsZero() {
		t.Errorf("nil id is not zero value, nilID: %v", nilID)
	}

	id, _ := NewID()
	if id.IsZero() {
		t.Errorf("id is zero value, id: %v", id)
	}
}
