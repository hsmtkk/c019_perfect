package perfect

type Type int

const (
	Undef Type = iota
	Perfect
	Nearly
	Neither
)

func Judge(n int) Type {
	sum := 0
	for _, d := range Divisor(n) {
		sum += d
	}
	if n == sum {
		return Perfect
	} else if (n-sum) == 1 || (n-sum) == -1 {
		return Nearly
	}
	return Neither
}

func Divisor(n int) []int {
	ds := []int{}
	for m := 1; m < n; m++ {
		if n%m == 0 {
			ds = append(ds, m)
		}
	}
	return ds
}
