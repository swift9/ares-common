package converters

import "strconv"

func Int642Int(val int64) int {
	return int(val)
}
func Int2Int64(val int) int64 {
	return int64(val)
}

func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Int2String(val int) string {
	return strconv.Itoa(val)
}
