package jobs

import (
	"encoding/json"
	"github.com/seamusv/fm-integration/fm"
	"github.com/seamusv/fm-integration/screens"
	"time"
)

func GeneratePurchaseOrderNumber(processor Processor, input []byte) ([]byte, error) {
	request := struct {
		CorrelationKey    string `json:"correlationKey" validate:"required"`
		Organisation      string `json:"organisation" validate:"required,max=6"`
		BillingAddress    string `json:"billingAddress" validate:"required,max=10"`
		OrderNumberPrefix string `json:"orderNumberPrefix" validate:"required,alpha"`
		ShippingAddress   string `json:"shippingAddress" validate:"required,max=10"`
		VendorCode        string `json:"vendorCode" validate:"required,max=12"`
	}{}
	if err := UnmarshalAndValidate(input, &request); err != nil {
		return nil, err
	}

	response := struct {
		CorrelationKey string `json:"correlationKey"`
		OrderNumber    string `json:"orderNumber"`
	}{
		CorrelationKey: request.CorrelationKey,
	}

	resultCh := make(chan error, 1)
	processor.Process(func(e Executor) {
		e.Login(NoLoginProfile, request.Organisation, time.Now())
		defer e.Logout()

		e.Execute("PO401", "Z00007")
		po401 := &screens.PO401{
			IDORDR:   fm.String(request.OrderNumberPrefix),
			IDVEND:   fm.String(request.VendorCode),
			LINEBILL: fm.String(request.BillingAddress),
			LINESCHD: fm.Time(fm.CurrentFiscalYear().End().Time()),
			LINESHPT: fm.String(request.ShippingAddress),
		}
		e.ExecuteFields("ADD", po401, "Z00062")
		if res := e.Execute("PROCESS", "P40163"); res != nil {
			orderNumber, _ := res.FieldValue("IDORDR")
			response.OrderNumber = orderNumber
		}

		resultCh <- e.Err()
	})

	err := <-resultCh
	if err != nil {
		return nil, err
	}

	return json.Marshal(response)
}
