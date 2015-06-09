/*
Package primes implements a library for generating prime numbers.
*/
package primes

type signal struct{}
type chanSignal chan signal

func NewChanSignal() chanSignal {
	return make(chan signal, 1)
}

func NewSignal() signal {
	return signal{}
}

// Generate a slices of prime numbers
// max prime number will be less than maxN
// out will be returning prime numbers
// you can stop generating by send signal via enough <- primes.NewSignal()
func Generate(maxN int, out chan<- int, enough chanSignal) {

	if maxN < 2 {
		return
	}

	res_mask := make([]bool, ceilToEven(maxN)/2-1)

	for i := range res_mask {
		res_mask[i] = true
	}

	out <- 2

	defer close(out)

	for i := 0; i < len(res_mask); i++ {
		select {
		case <-enough:
			return
		default:
			if res_mask[i] == true {
				out <- idx2value(i)
				markCompositeNumbers(res_mask, i)
			}
		}
	}
}

func markCompositeNumbers(mask []bool, idx int) {
	cur_val := idx2value(idx)
	for sum := cur_val * 2; sum <= idx2value(len(mask)-1); sum += cur_val {
		new_idx, ok := value2idx(sum)
		if !ok {
			continue
		}
		mask[new_idx] = false
	}
}

func idx2value(idx int) int {
	return idx*2 + 3
}

func ceilToEven(v int) int {
	if v%2 != 0 {
		return v + 1
	}
	return v
}

func value2idx(value int) (int, bool) {
	if value%2 == 0 {
		return 0, false
	}
	return (value - 3) / 2, true
}
