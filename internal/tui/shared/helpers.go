package shared

import "fmt"

// FormatNumberLong formats an integer with comma separators.
// e.g. 1234567890 -> 1,234,567,890
func FormatNumberLong(n int) string {
	if n == 0 {
		return "0"
	}
	negative := n < 0
	if negative {
		n = -n
	}
	s := fmt.Sprintf("%d", n)
	var result []byte
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, byte(c))
	}
	if negative {
		return "-" + string(result)
	}
	return string(result)
}

// FormatLargeNumber formats a number with K or M suffixes.
// e.g. 1234567890 -> 1.2M
func FormatLargeNumber(n int) string {
	if n >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(n)/1000000.0)
	}
	if n >= 1000 {
		return fmt.Sprintf("%.1fK", float64(n)/1000.0)
	}
	return fmt.Sprintf("%d", n)
}

// FormatPercentageBonus formats the percentage input as `+X%`
// e.g. 1.25 -> +125%
func FormatPercentageBonus(bonus float64) string {
	return fmt.Sprintf("+%.0f%%", bonus*100)
}
