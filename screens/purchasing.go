package screens

import "time"

type (
	PO401 struct {
		IDORDR   *string    `fm:"IDORDR"`
		IDVEND   *string    `fm:"IDVEND"`
		LINEBILL *string    `fm:"LINEBILL"`
		LINESCHD *time.Time `fm:"LINESCHD"`
		LINESHPT *string    `fm:"LINESHPT"`
		IDOP01   *string    `fm:"IDOP01"`
		IDOP02   *string    `fm:"IDOP02"`
		LINEMTCH *int       `fm:"LINEMTCH"`
		LINESHTY *int       `fm:"LINESHTY"`
		PARTYPE  *int       `fm:"PARTYPE"`
		SWREL    *bool      `fm:"SWREL"`
		LINETOL  *string    `fm:"LINETOL"`
	}
)
