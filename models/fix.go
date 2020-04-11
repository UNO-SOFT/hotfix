package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gobuffalo/pop/v5"
)

// Fix
type Fix struct {
	Name                 string `json:"name" db:"name"`
	CreatedAt, UpdatedAt time.Time
	Where                []string
	Events               []Event
	Deployed             bool
	State                FixState
	Todo                 []string
}
type FixState int8

const (
	FixAllowed = FixState(+2)
	FixOK      = FixState(+1)
	FixUnknown = FixState(0)
	FixNOK     = FixState(-1)
	FixBanned  = FixState(-2)
)

func (fs FixState) String() string {
	switch fs {
	case -2:
		return "Banned"
	case -1:
		return "NOK"
	case 1:
		return "OK"
	case 2:
		return "Allowed"
	default:
		return "Unknown"
	}
}

// String is not required by pop and may be deleted
func (f Fix) String() string {
	je, _ := json.Marshal(f)
	return string(je)
}

func GetFixes(tx *pop.Connection) ([]Fix, error) {
	var events []Event
	if err := tx.Order("f_with ASC, f_when ASC").All(&events); err != nil {
		return nil, err
	}
	log.Println("events:", events)
	var f Fix
	var fixes []Fix
	where := make(map[string]time.Time)
	A := func(f Fix) {
		if f.Name != "" {
			f.Where = make([]string, 0, len(where))
			for k := range where {
				f.Where = append(f.Where, k)
			}
			if !f.Deployed && f.State != FixBanned {
				f.Todo = append(f.Todo, "-2")
				if f.State != FixAllowed {
					f.Todo = append(f.Todo, "-1", "+1", "+2")
				}
			}
			fixes = append(fixes, f)
		}
		for k := range where {
			delete(where, k)
		}
	}
	for _, e := range events {
		if f.Name != e.With {
			A(f)
			f = Fix{Name: e.With, CreatedAt: e.When}
		}
		where[e.Where] = e.When
		f.UpdatedAt = e.When
		f.Events = append(f.Events, e)
		f.Deployed = f.Deployed || e.Where == "+PRD" || e.Where == "PRD"
		if !f.Deployed {
			switch e.What {
			case "-2":
				f.State = FixBanned
			case "-1":
				if f.State != FixBanned && f.State != FixAllowed {
					f.State = FixNOK
				}
			case "+1":
				if f.State != FixBanned && f.State != FixAllowed {
					f.State = FixOK
				}
			case "+2":
				if f.State != FixBanned {
					f.State = FixAllowed
				}
			default:
			}
		}
	}
	A(f)
	return fixes, nil
}
