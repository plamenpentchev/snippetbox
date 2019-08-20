package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field can not be blank")
		}
	}
}

func (f *Form) MaxLength(field string, d int) {
	v := f.Get(field)
	l := utf8.RuneCountInString(v)
	if l > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long [%d], maximum is %d characters", l, d))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	v := f.Get(field)
	for _, opt := range opts {
		if v == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
