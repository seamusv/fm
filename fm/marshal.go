package fm

// https://itnext.io/creating-your-own-struct-field-tags-in-go-c6c86727eff

import (
	"encoding/xml"
	"fmt"
	"github.com/seamusv/fm-integration/fm/internal"
	"reflect"
	"strconv"
	"time"
)

func Marshal(operation string, v interface{}) ([]byte, error) {
	fields := make([]internal.Field, 0)

	rv := reflect.ValueOf(v)
	t := rv.Type()
	for i := 0; i < t.NumField(); i++ {
		if name, ok := t.Field(i).Tag.Lookup("fm"); ok {
			field := rv.Field(i)
			if field.Kind() != reflect.Ptr {
				return nil, fmt.Errorf("no support for non-pointer field '%s %T'", t.Field(i).Name, field.Type().String())
			}
			if !field.IsNil() {
				fieldValue := field.Interface()
				var strValue string

				switch v := fieldValue.(type) {
				case *string:
					strValue = *v
					break
				case *int64:
					strValue = strconv.FormatInt(*v, 10)
					break
				case *int:
					strValue = strconv.FormatInt(int64(*v), 10)
					break
				case *bool:
					if *v {
						strValue = "Y"
					} else {
						strValue = "N"
					}
					break
				case *time.Time:
					strValue = v.Format("2006/01/02")
					break
				default:
					return nil, fmt.Errorf("no support for '%s %T'", t.Field(i).Name, v)
				}
				fields = append(fields, internal.Field{Name: name, Value: strValue})
			}
		}
	}

	transaction := &internal.Request{
		Gui:     "N",
		Command: internal.Command{Operation: operation},
		Fields:  fields,
	}

	return xml.MarshalIndent(transaction, "", "  ")
}
