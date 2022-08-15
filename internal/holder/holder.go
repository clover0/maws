package holder

type Holder interface {
	Add(key, elem string)
	OutAll() string
}
