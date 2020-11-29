package jobs

import (
	"fmt"
	"github.com/seamusv/fm-integration/encoding"
	"github.com/seamusv/fm-integration/jobs/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"
	"time"
)

func TestGeneratePurchaseOrderNumber(t *testing.T) {
	executor := &mocks.Executor{}
	executor.On("Login", mock.Anything, mock.Anything, mock.Anything).Return()
	executor.On("Logout").Return()
	executor.On("Err").Return(nil)
	executor.On("Execute", "PO401", "Z00007").Return(nil)
	executor.On("ExecuteFields", "ADD", mock.AnythingOfType("screens.PO401"), "Z00062").Return(
		func(command string, screen interface{}, expectedCodes ...string) *encoding.Response {
			fields, err := buildFieldMap(screen)
			assert.NoError(t, err)
			assert.Equal(t, "2018/03/31", fields["LINESCHD"])
			return nil
		})
	executor.On("Execute", "PROCESS", "P40163").Return(
		func(command string, expectedCodes ...string) *encoding.Response {
			r, err := encoding.Parse(buildMockXMLResponse(map[string]string{"IDORDR": "C00006942"}))
			assert.NoError(t, err)
			return r
		})

	processor := &MockProcessor{Executor: executor}
	input := []byte(`{
		"correlationKey": "corr-123",
		"organisation": "YUKON",
		"orderNumberPrefix": "C",
		"billingAddress": "BHPWICTW10",
		"shippingAddress": "SHPWICTW10",
		"vendorCode": "CDVENASSESYS"
	}`)

	output, err := GeneratePurchaseOrderNumber(processor, buildDate(2017, time.December, 24), input)
	assert.NoError(t, err)
	assert.Regexp(t, `"orderNumber"\s*:\s*"C00006942"`, string(output))
}

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
