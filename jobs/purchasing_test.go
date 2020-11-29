package jobs

import (
	"fmt"
	"github.com/seamusv/fm-integration/encoding"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGeneratePurchaseOrderNumber(t *testing.T) {
	input := []byte(`{
		"correlationKey": "corr-123",
		"organisation": "YUKON",
		"orderNumberPrefix": "C",
		"billingAddress": "BHPWICTW10",
		"shippingAddress": "SHPWICTW10",
		"vendorCode": "CDVENASSESYS"
	}`)

	output, err := GeneratePurchaseOrderNumber(&MockProcessor{}, input)
	assert.NoError(t, err)
	fmt.Printf("ORDER: %s\n", string(output))
}

type MockExecutor struct {
	err error
}

func (m *MockExecutor) Login(profile, organisation string, businessDate time.Time) {
}

func (m *MockExecutor) Logout() {
}

func (m *MockExecutor) Execute(command string, messageCodes ...string) *encoding.Response {
	r, _ := encoding.Parse([]byte(`<trans ok="Y"><screendata><return-fields><f n="IDORDR" v="C00006942"/></return-fields></screendata></trans>`))
	return r
}

func (m *MockExecutor) ExecuteFields(command string, v interface{}, messageCodes ...string) *encoding.Response {
	return nil
}

func (m *MockExecutor) Err() error {
	return m.err
}

type MockProcessor struct {
}

func (p *MockProcessor) Process(f func(executor Executor)) {
	f(&MockExecutor{})
}
