package geneticAlgorithm

import (
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//洗牌算法
func ShuffleStrings(strings []string) ([]string, error) {
	if len(strings) <= 0 {
		return nil, errors.New("the length of the parameter strings should not be less than 0")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}
	return strings, nil
}

func ShuffleIntegers(ints []int) ([]int, error) {
	if len(ints) <= 0 {
		return nil, errors.New("the length of the parameter strings should not be less than 0")
	}

	for i := len(ints) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		ints[i], ints[num] = ints[num], ints[i]
	}
	return ints, nil
}

func Rotate(t []int, n int) []int {
	tmp := make([]int, n)
	copy(tmp, t[:n])
	copy(t[:len(t)-n], t[n:])
	copy(t[len(t)-n:], tmp[:])
	return t
}
