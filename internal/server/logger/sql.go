package logger

import (
	"log"
	"strings"
)

func Sql(query string, params []string) {
	log.Printf("SQL: %s, [%s]", query, strings.Join(params, ","))
}
