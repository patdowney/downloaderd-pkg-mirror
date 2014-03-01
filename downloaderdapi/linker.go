package downloaderdapi

type Linked interface {
	GetLink(relation string) string
}
