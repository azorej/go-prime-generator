package primes

import "testing"
import "time"
import "errors"

const TIMEOUT = time.Second * 2 // 2 seconds

func TestReverse(t *testing.T) {
	gage_arr := [][]int{{2}, {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149}}

	for _, gage := range gage_arr {
		testGage(t, gage)
	}
}

func testGage(t *testing.T, gage []int) {

	maxN := gage[len(gage)-1]

	res, err := generatePrimes(maxN)
	if err != nil {
		t.Error(err)
		return
	}

	if !compareSlices(res, gage) {
		t.Errorf("Wanted simple numbers for %v, got %v", maxN, res)
	}

}

func generatePrimes(maxN int) ([]int, error) {

	res := make([]int, 0)
	ch := make(chan int, 20)
	enough := NewChanSignal()

	go Generate(maxN, ch, enough)

	timer := time.NewTimer(TIMEOUT) //Timeout
	defer timer.Stop()

	for {
		select {
		case v, more := <-ch:
			if !more {
				return res, nil
			}
			res = append(res, v)

		case <-timer.C:
			enough <- NewSignal()
			return res, errors.New("Timeout :(")
		}
	}

}

func compareSlices(s1, s2 []int) bool {

	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}

	return true
}
