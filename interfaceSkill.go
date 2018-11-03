package main

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20

	var w welcome = "hao"

	Each(persons, w)
}

type welcome string

func (w welcome) Do(k, v interface{}) {

}

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}
