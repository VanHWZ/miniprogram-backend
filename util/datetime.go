package util

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Datetime time.Time

func (dt *Datetime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*dt = Datetime(value)
		return nil
	}
	return errors.New("wrong time type")
}

func (dt Datetime) Value() (driver.Value, error) {
	datetime := time.Time(dt)
	var zeroTime time.Time
	if datetime.Unix() == zeroTime.Unix() {
		return nil, nil
	}
	return datetime, nil
}

const (
	DatetimeFormat = "2006-01-02 15:04:05"
	RFC3339Nano    = "2006-01-02T15:04:05.999999999Z07:00"
)

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (dt Datetime) MarshalJSON() ([]byte, error) {
	t := time.Time(dt)
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DatetimeFormat)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, DatetimeFormat)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON string to time in json
func (dt *Datetime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	var t *time.Time
	t = new(time.Time)
	*t, err = time.ParseInLocation(`"`+DatetimeFormat+`"`, string(data), time.Local)
	*dt = Datetime(*t)
	return err
}