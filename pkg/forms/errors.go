package forms


type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	res := e[field]
	if len(res) == 0 {
		return ""
	}
	return res[0]
}
