package jobs

import (
	"fmt"
	"github.com/seamusv/fm-integration/encoding"
	"strings"
	"time"
)

type MockProcessor struct {
	Executor Executor
}

func (p *MockProcessor) Process(f func(executor Executor)) {
	f(p.Executor)
}

func buildFieldMap(screen interface{}) (map[string]string, error) {
	f, err := encoding.Fields(screen)
	if err != nil {
		return nil, err
	}
	fields := make(map[string]string)
	for _, field := range f {
		fields[field.Name] = field.Value
	}
	return fields, nil
}

func buildDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func buildMockXMLResponse(fields map[string]string) []byte {
	elements := make([]string, 0)
	for n, v := range fields {
		elements = append(elements, fmt.Sprintf(`<f n="%s" v="%s" />`, n, v))
	}
	return []byte(fmt.Sprintf(`<trans ok="Y"><screendata><return-fields>%s</return-fields></screendata></trans>`, strings.Join(elements, "")))
}
