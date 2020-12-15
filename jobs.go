package fm_integration

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"time"
)

const (
	NoLoginProfile = ""
)

type (
	Processor interface {
		Process(func(Executor))
	}

	Executor interface {
		Login(profile, organisation string, businessDate time.Time)
		Logout()
		Execute(command string, messageCodes ...string) *Response
		ExecuteFields(command string, v interface{}, messageCodes ...string) *Response
		Err() error
	}
)

var (
	validate = validator.New()
)

func UnmarshalAndValidate(input []byte, v interface{}) error {
	if err := json.Unmarshal(input, v); err != nil {
		return err
	}
	return validate.Struct(v)
}

//go:generate mockery --name Executor
