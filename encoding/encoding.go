package encoding

type (
	Field struct {
		Name  string `xml:"n,attr"`
		Value string `xml:"v,attr"`
	}

	Message struct {
		Number      string `xml:"no"`
		Description string `xml:"v"`
	}
)
