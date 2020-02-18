package status

import (
	"fmt"

	"gopkg.in/guregu/null.v3"
)

type Status struct {
	Name     string      `db:"name" json:"name" valid:"utfletternumspace"`
	Code     string      `db:"code" json:"code" valid:"alphanum"`  // кодовое имя
	Priority int         `db:"priority" json:"priority" valid:"-"` // приоритет , статусы с более высоким значеием не могут заменены с более низки значение
	Color    string      `db:"color" json:"color" valid:"hexcolor"`
	Comment  null.String `db:"comment" json:"comment" valid:"null_utfletternumspace"`
	Deleted  bool        `db:"deleted" json:"deleted" valid:"-"`
}
type Statuses []Status

func (stats *Statuses) GetStatusesFromStatusCode(code string) (Statuses, error) {
	st, err := stats.GetStatusByCode(code)
	if err != nil {
		return Statuses{}, err
	}
	return stats.GetStatusesFromPriority(st.Priority), nil
}
func (stats *Statuses) GetStatusesFromPriority(pr int) Statuses {
	result := Statuses{}
	for _, st := range *stats {
		if st.Priority >= pr {
			result = append(result, st)
		}
	}

	return result
}
func (stats *Statuses) GetStatusByCode(code string) (Status, error) {
	for _, st := range *stats {
		if st.Code == code {
			return st, nil
		}
	}
	return Status{}, fmt.Errorf("status %v not found", code)
}
func (stats *Statuses) GetCodes() []string {
	result := []string{}
	for _, s := range *stats {
		result = append(result, s.Code)
	}
	return result
}

// Получить статысы в промежутке между приорететами
func (stats *Statuses) GetBeetwenByPriority(from, to int) Statuses {
	result := Statuses{}
	for _, st := range *stats {
		if st.Priority >= from && st.Priority <= to {
			result = append(result, st)
		}
	}
	return result
}
