package util

var DEFAULT_LIMIT, DEFAULT_OFFSET int64 = 20, 0

func DefaultOrderSqlParam(order, asc string) string {
	if len(order) == 0 {
		order = "id"
	}
	switch asc {
	case "true":
		asc = "asc"
	case "false":
		asc = "desc"
	case "asc":
		asc = "asc"
	case "desc":
		asc = "desc"
	default:
		asc = "desc"
	}
	return order + " " + asc
}
func DefaultLimitSqlParam(pageSize, pageNum int64) (limit, offset int) {
	if pageSize > 0 && pageNum > 0 {
		return int(pageSize), int((pageNum - 1) * pageSize)
	}
	return int(DEFAULT_LIMIT), int(DEFAULT_OFFSET)
}
