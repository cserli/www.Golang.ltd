package util

import (
	"LollipopGo/LollipopGo/log"
	"strconv"
)

func Str2int_LollipopGo(data string) int {
	v, err := strconv.Atoi(data)
	if err != nil {
		log.Debug(err.Error())
		return -1
	}
	return v
}

func Int2str_LollipopGo(data int) string {
	return strconv.Itoa(data)
}
