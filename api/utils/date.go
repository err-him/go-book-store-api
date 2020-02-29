package utils

import (
	"encoding/json"
	"strings"
	"time"
)

type MysqlFormatDate time.Time

func (f *MysqlFormatDate) UnmarshalJSON(b []byte) (err error) {
	str := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return
	}
	*f = MysqlFormatDate(t)
	return nil
}

func (f MysqlFormatDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(f)
}

func (f MysqlFormatDate) Format(s string) string {
	t := time.Time(f)
	return t.Format(s)
}
