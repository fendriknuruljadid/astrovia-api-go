package utils

import "strings"

func ParseDBError(err error) string {
    e := err.Error()

    switch {
    case strings.Contains(e, "duplicate key"):
        return "duplicate data"
    case strings.Contains(e, "foreign key constraint"):
        return "invalid relation"
    case strings.Contains(e, "not-null constraint"):
        return "missing required field"
    default:
        return e
    }
}
