package internal

import "encoding/xml"

type (
	Request struct {
		XMLName xml.Name `xml:"trans"`
		Gui     string   `xml:"gui,attr"`
		Command Command  `xml:"command"`
		Fields  []Field  `xml:"screendata>put-fields>f"`
	}

	Response struct {
		XMLName  xml.Name  `xml:"trans"`
		Fields   []Field   `xml:"screendata>return-fields>f"`
		Messages []Message `xml:"msgs>msg"`
	}

	Command struct {
		Operation string `xml:"cmd,attr"`
	}

	Field struct {
		Name  string `xml:"n,attr"`
		Value string `xml:"v,attr"`
	}

	Message struct {
		Number      string `xml:"no"`
		Description string `xml:"v"`
	}
)
