package pkg

import "fmt"

const kb int64 = 1024
const mb int64 = 1024 * kb
const gb int64 = 1024 * mb
const tb int64 = 1024 * gb

func ByteLabel(i int64) string {
	if i < 0 {
		return "Error"
	}

	var unit string
	var divisor int64

	if i >= tb {
		divisor = tb
		unit = "TB"
	} else if i >= gb {
		divisor = gb
		unit = "GB"

	} else if i >= mb {
		divisor = mb
		unit = "MB"

	} else if i >= kb {
		divisor = kb
		unit = "KB"
	} else {
		divisor = 1
		unit = "B"
	}

	num := float64(i) / float64(divisor)
	return fmt.Sprintf("%.2f %s", num, unit)
}
