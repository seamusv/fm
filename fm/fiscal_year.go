package fm

import "time"

type FiscalYear struct {
	t time.Time
}

func CurrentFiscalYear() *FiscalYear {
	return &FiscalYear{t: time.Now()}
}

func (f FiscalYear) Begin() *FiscalYear {
	var year = f.t.Year()
	if f.t.Month() < time.April {
		year -= 1
	}
	return &FiscalYear{t: time.Date(year, time.April, 1, 0, 0, 0, 0, f.t.Location())}
}

func (f FiscalYear) End() *FiscalYear {
	var year = f.t.Year()
	if f.t.Month() > time.March {
		year += 1
	}
	return &FiscalYear{t: time.Date(year, time.March, 31, 0, 0, 0, 0, f.t.Location())}
}

func (f FiscalYear) Time() time.Time {
	return f.t
}
