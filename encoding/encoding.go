package encoding

import "encoding/xml"

type (
	XMLRequest struct {
		XMLName xml.Name `xml:"trans"`
		Gui     string   `xml:"gui,attr"`
		Command Command  `xml:"command"`
		Fields  []Field  `xml:"screendata>put-fields>f"`
	}

	XMLResponse struct {
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
		Number      string `xml:"no,attr"`
		Description string `xml:"v,attr"`
	}
)
