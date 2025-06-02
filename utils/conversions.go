package utils

import (
	"database/sql"
	"fmt"
)

func NullableToString(ns sql.NullString, fallback string) string {
    if ns.Valid {
        return ns.String
    }
    return fallback
}

func NullableIntToString(ni sql.NullInt64, fallback string) string {
    if ni.Valid {
        return fmt.Sprint(ni.Int64)
    }
    return fallback
}