package numberUtil

import (
	"fmt"
	"strconv"
)

// Round float小数点保留位数
// @param f
// @param n 保留几位小数点后数字
// @return float64
func RoundFloat(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}
