package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {

	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error(("got invalid when should have been valid"))
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}

	// // 異常系
	form := New(postedData)

	has := form.Has("hoge")
	if has {
		t.Error("form shows has field when it does not")
	}

	// 正常系
	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it shold")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}

	// 異常系
	form := New(postedData)

	form.MinLength("x", 3)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	// errors.go Get function
	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	// 異常系
	postedData = url.Values{}
	postedData.Add("some_field", "aaaaa")
	form = New(postedData)

	form.MinLength("some_field", 100000)
	if form.Valid() {
		t.Error("shows minlength of 100000 met when data is shorter")
	}

	// 正常系
	postedData = url.Values{}
	postedData.Add("another_field", "aaaaa")
	form = New(postedData)
	form.MinLength("another_field", 3)

	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when it is")
	}

	// errors.go Get function
	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}

	// 異常系
	form := New(postedData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	// 正常系
	postedData = url.Values{}
	postedData.Add("email", "example@gmail.com")
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() == false {
		t.Error("got an invalid email when we shoud not have")
	}
	// 異常系
	postedData = url.Values{}
	postedData.Add("email", "example@")
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address")
	}

}
