package converters

import "strconv"

func String2Bytes(val string) []byte {
	return []byte(val)
}

func String2Int64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

func String2Int(val string) (int, error) {
	return strconv.Atoi(val)
}
