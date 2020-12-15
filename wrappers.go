package fm_integration

import "time"

func Bool(v bool) *bool {
	return &v
}

func Float64(v float64) *float64 {
	return &v
}

func Int(v int) *int {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func String(v string) *string {
	return &v
}

func Time(v time.Time) *time.Time {
	return &v
}
