package win32

import "strconv"

func HexToDecimal(hexStr string) int {
	// 将十六进制字符串转换为十进制整数
	decimal64, err := strconv.ParseInt(hexStr, 16, 32)
	if err != nil {
		return 0
	}
	decimal := int(decimal64) // 将 int64 转换为 int
	return decimal
}
