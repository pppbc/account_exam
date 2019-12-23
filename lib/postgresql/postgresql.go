package util

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"

	json "account_exam/lib/pgencoding/json2"
)

// PgJsonScanWrap 中的 struct tag 使用的是 `db:"name"`
// time should be type: TIMESTAMP WITH TIME ZONE
type PgJsonScanWrap struct {
	Value interface{}
}

func (p *PgJsonScanWrap) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	data, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("JsonScanWrap must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(data, p.Value)
}

func PgMapQuery(sql string, maps map[string]interface{}) (string, []interface{}) {
	var (
		args = make([]interface{}, 0, len(maps))
		n    = 0
	)
	for name, val := range maps {
		placeholder := "$" + name
		if strings.Contains(sql, placeholder) {
			rval := reflect.ValueOf(val)
			_, ok := val.(driver.Valuer)
			if !ok && rval.Kind() == reflect.Slice {
				num := rval.Len() - 1
				params := ""
				for i := 0; i <= num; i++ {
					n++
					if i < num {
						params += fmt.Sprintf("$%d, ", n)
					}
					args = append(args, rval.Index(i).Interface())
				}
				if num >= 0 {
					params += fmt.Sprintf("$%d", n)
				}
				sql = strings.Replace(sql, placeholder, "("+params+")", -1)
				continue
			}
			n++
			sql = strings.Replace(sql, placeholder, fmt.Sprintf("$%d", n), -1)
			args = append(args, val)
		}
	}
	return sql, args
}

var likeReplacer = strings.NewReplacer(
	`\`, `\\`,
	`%`, `\%`,
	`_`, `\_`,
)

func PgEscapeLike(value string) string {
	return likeReplacer.Replace(value)
}
