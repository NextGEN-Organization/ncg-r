package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"runtime/debug"
	"strconv"
	"time"
)

var (
// par = parallelizer.NewGroup()
)

const (
	// maxCPUs  = 8
	// filename = "test.txt"
	NumLns                 = 50 * 1000000
	LettersPerLine         = 16
	letters                = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lettersLength          = int64(len(letters))
	newLine                = byte('\n')
	lettersPerLinePlusOne  = int(LettersPerLine + 1)
	lettersPerLineMinusOne = LettersPerLine - 1
	letterIdxBits          = int64(6)
	letterIdxMask          = 1<<letterIdxBits - 1
	letterIdxMask8         = int8(letterIdxMask)
	letterIdxMax           = int8(63 / letterIdxBits)
)

func init() {
	debug.SetGCPercent(-1)
}

func NLinesInBytes(n int, b *[]byte) {

	var (
		idx    int64
		j      int
		cache  int64
		remain int8
	)

	for i := 0; i < n; i++ {
		// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
		for j, cache, remain = lettersPerLineMinusOne, rand.Int63(), letterIdxMax; j >= 0; {
			if remain == 0 {
				cache, remain = rand.Int63(), letterIdxMax
			}
			if idx = cache & letterIdxMask; idx < lettersLength {
				(*b)[j+i*lettersPerLinePlusOne] = letters[idx]
				j--
			}
			cache >>= letterIdxBits
			remain--
		}
		(*b)[(i+1)*lettersPerLinePlusOne-1] = newLine
	}
}

func SaveToFile(filename string, b *[]byte) {
	if err := ioutil.WriteFile(filename, *b, 0644); err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]
	d, err := strconv.Atoi(args[0])

	if err != nil {
		panic(err)
	}

	val := make([]byte, d*(LettersPerLine+1))
	startTime := time.Now()
	NLinesInBytes(d, &val)
	SaveToFile(args[1], &val)
	diff := time.Since(startTime)
	println("Ran in", diff.Milliseconds(), "ms.\nGenerated ", d/1000000, "million codes.")
}
