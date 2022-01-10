package iteration

func Repeat(arg string, times ...int) string {
	var (
		repeated string
		t        = 5
	)
	if len(times) > 0 {
		t = times[0]
	}
	for i := 0; i < t; i++ {
		repeated += arg
	}
	return repeated
}
