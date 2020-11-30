package encoding

import (
	"encoding/xml"
	"fmt"
	"github.com/seamusv/fm-integration/fm"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type (
	Response struct {
		internalResponse *XMLResponse
		data             map[string]string
		messages         map[string]string
	}
)

func Parse(b []byte) (*Response, error) {
	ir := &XMLResponse{}
	if err := xml.Unmarshal(b, ir); err != nil {
		return nil, err
	}
	data := make(map[string]string)
	for _, field := range ir.Fields {
		data[field.Name] = field.Value
	}
	messages := make(map[string]string)
	for _, message := range ir.Messages {
		messages[message.Number] = message.Description
	}
	return &Response{
		internalResponse: ir,
		data:             data,
		messages:         messages,
	}, nil
}

func (r *Response) Unmarshal(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("struct must be a pointer value")
	}
	t := reflect.TypeOf(v).Elem()
	for i := 0; i < t.NumField(); i++ {
		if name, ok := t.Field(i).Tag.Lookup("fm"); ok {
			if dataValue, ok := r.data[name]; ok {
				field := rv.Elem().Field(i)
				if field.Kind() != reflect.Ptr {
					return fmt.Errorf("no support for non-pointer field '%s %T'", t.Field(i).Name, field.Type().String())
				}
				switch v := field.Interface().(type) {
				case *string:
					field.Set(reflect.ValueOf(fm.String(r.data[name])))
					break
				case *int64:
					if len(dataValue) > 0 {
						i, err := strconv.ParseInt(r.data[name], 10, 64)
						if err != nil {
							return fmt.Errorf("error parsing '%s' as %T: %s", name, v, dataValue)
						}
						field.Set(reflect.ValueOf(fm.Int64(i)))
					}
					break
				case *int:
					if len(dataValue) > 0 {
						i, err := strconv.Atoi(r.data[name])
						if err != nil {
							return fmt.Errorf("error parsing '%s' as %T: %s", name, v, dataValue)
						}
						field.Set(reflect.ValueOf(fm.Int(i)))
					}
					break
				case *bool:
					if len(dataValue) > 0 {
						field.Set(reflect.ValueOf(fm.Bool(dataValue == "Y")))
					}
					break
				case *time.Time:
					if len(dataValue) > 0 {
						t, err := time.Parse("2006/01/02", dataValue)
						if err != nil {
							return fmt.Errorf("error parsing '%s' as %T: %s", name, v, dataValue)
						}
						field.Set(reflect.ValueOf(fm.Time(t)))
					}
				default:
					return fmt.Errorf("no support for '%s %T'", t.Field(i).Name, v)
				}
			}
		}
	}
	return nil
}

func (r *Response) FieldValue(field string) (string, bool) {
	data, ok := r.data[field]
	return data, ok
}

func (r *Response) MessageContainsOneOf(messages ...string) error {
	errors := make([]string, 0)
	for k, v := range r.messages {
		for _, m := range messages {
			if k == m {
				return nil
			}
		}
		errors = append(errors, fmt.Sprintf("%s: %s", k, v))
	}
	return fmt.Errorf("expecting one of [%s], received: %s", strings.Join(messages, ", "), strings.Join(errors, "; "))
}
