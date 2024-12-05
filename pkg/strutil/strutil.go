package strutil

import (
	"math"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

// LeftPad leftpads str with char to length n
// if str is gte n, bail out
func LeftPad(str string, n int, char rune) string {
	if n < 1 {
		return str
	}

	numCharToAppend := n - len(str)
	if numCharToAppend <= 0 {
		return str
	}

	var builder strings.Builder

	for i := 0; i < numCharToAppend; i++ {
		builder.WriteRune(char)
	}

	builder.WriteString(str)

	return builder.String()
}

func NumberFormat(n float64, precision int, isEmpty bool) string {
	if n == 0 && isEmpty {
		return ""
	}

	coefficient := math.Pow(10, float64(precision))
	n = math.Round(n*coefficient) / coefficient

	p := message.NewPrinter(language.English)
	withCommaThousandSep := p.Sprintf("%f", number.Decimal(n, number.MinFractionDigits(precision), number.MaxFractionDigits(precision)))

	return withCommaThousandSep
}
