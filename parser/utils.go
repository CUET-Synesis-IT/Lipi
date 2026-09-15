package parser

func parseBengaliDigits(num string) int64 {
	out := int64(0)
	for _, r := range num {
		out = out*10 + int64(r-'০')
	}
	return out
}
