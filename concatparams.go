package student

func ConcatParams(args []string) string {
	var o string
	for i, ss := range args {
		if i == 0 {
			o = ss
			continue
		}
		o += "\n" + ss
	}
	return o
}
