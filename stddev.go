// Command stddev reads float64 values on its stdin and outputs number of
// samples, population standard deviation and mean.
//
// Values that cannot be parsed as float64 are silently ignored.
//
//	printf "1\n2\n3\n4\n5\njunk\n" | stddev
//	5 samples: stddev 1.414214 mean 3.000000
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/antonlindstrom/gostddev"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		fmt.Fprintln(os.Stderr, "Expects float64 values on stdin, one per line")
		os.Exit(2)
	}
	cnt, sd, mean, err := do(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%d samples: stddev %f mean %f\n", cnt, sd, mean)
}

func do(rd io.Reader) (cnt int, sd, mean float64, err error) {
	sc := bufio.NewScanner(rd)
	var nums []float64
	for sc.Scan() {
		f, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			continue
			// return 0, 0, err
		}
		nums = append(nums, f)
	}
	return len(nums), gostddev.StdDev(nums), gostddev.Mean(nums), sc.Err()
}
