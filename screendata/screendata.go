package screendata

// https://itnext.io/creating-your-own-struct-field-tags-in-go-c6c86727eff

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func Marshal(operation string, v interface{}) ([]byte, error) {
	fields := make([]Field, 0)

	rv := reflect.ValueOf(v)
	t := rv.Type()
	for i := 0; i < t.NumField(); i++ {
		if name, ok := t.Field(i).Tag.Lookup("fm"); ok {
			fieldValue := rv.Field(i).Interface()
			var strValue string

			switch v := fieldValue.(type) {
			case string:
				strValue = v
				break
			case int64:
				strValue = strconv.FormatInt(v, 10)
				break
			case int:
				strValue = strconv.FormatInt(int64(v), 10)
				break
			case bool:
				if v {
					strValue = "Y"
				} else {
					strValue = "N"
				}
				break
			case time.Time:
				strValue = v.Format("2006/01/02")
				break
			default:
				return nil, fmt.Errorf("unknown '%T' type provided", v)
			}
			fields = append(fields, Field{Name: name, Value: strValue})
		}
	}

	transaction := &Transaction{
		Gui: "N",
		Command: Command{
			Operation: operation,
			Process:   "FM",
			Fields:    fields,
		},
	}

	return xml.MarshalIndent(transaction, "", "  ")
}

type (
	Transaction struct {
		XMLName xml.Name `xml:"trans"`
		Gui     string   `xml:"gui,attr"`
		Command Command
	}

	Command struct {
		XMLName   xml.Name `xml:"command"`
		Operation string   `xml:"cmd,attr"`
		Process   string   `xml:"proc,attr"`
		Fields    []Field  `xml:"screendata>put-fields>f,omitempty"`
	}

	Field struct {
		Name  string `xml:"n,attr"`
		Value string `xml:"v,attr"`
	}
)
