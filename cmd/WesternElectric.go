package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	movingAverage "github.com/davecb/WesternElectric/pkg/MovingAverage"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

/*
 * Western Electric Rules -- a classic set of decision rules in statistical process control
 * for detecting out-of-control or non-random conditions.
 */

func usage() {
	//nolint
	fmt.Fprint(os.Stderr, "Usage: westernelectric --samples N {file|-}\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var nSamples int
	var fp *os.File
	var err error

	flag.IntVar(&nSamples, "nSamples", 5, "number of samples in the moving average")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprint(os.Stderr, "You must supply an input file, or '-' and a stream on stdin\n\n") //nolint
		usage()
	}
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime) // show file:line in logs

	filename := flag.Arg(0)
	if filename == "-" {
		// if the filename is "-", read stdin
		fp = os.Stdin
	} else {
		fp, err = os.Open(filename) //nolint
		if err != nil {
			log.Fatalf("error opening %s: %q, halting.", filename, err)
		}
		defer func() {
			err := fp.Close()
			if err != nil {
				log.Printf("Close of input file %q failed, ignored. %v\n",
					filename, err)
			}
		}()
	}
	worker(fp, nSamples)
}

// worker reads the input and applies the rules, comparing the data
// to a rolling average.
func worker(fp *os.File, nSamples int) {
	var nr int
	var average float64
	var sd float64

	// set up csv reader to read fields out of a file
	r := csv.NewReader(fp)
	r.Comma = '\t'
	r.Comment = '#'
	r.FieldsPerRecord = -1 // ignore differences
	r.LazyQuotes = true    // allow bad quoting

	// set up moving average
	add := movingAverage.New(nSamples)

	// read lines containing a datestamp and a (usually floating-point) value
	fmt.Printf("%s\t%s\t%s\t%s %s\n", "#date", "datum", "average", "stddev", "flags")
	for nr = 0; ; nr++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error %q, in %q, line %d\n", err, record, nr)
		}
		if len(record) < 2 {
			// skip it, but complain
			log.Printf("error in line %d, %q. Ignored.\n", nr, record)
			continue
		}

		date := record[0]
		// parse the value field
		datum, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatalf("Encountered invalid float64 in %q on line %d, halting.\n", record, nr)
		}

		//log.Printf("at time %q, got %g, average = %g, sd = %g\n", record[0], datum, average, sd)
		if nr > nSamples {
			// see if we break any of the rules once we have an average to use
			three := threeSigma(datum, average, sd)
			two := ""
			one := ""

			// 	print stats
			fmt.Printf("%s\t%g\t%0.4g\t%0.4g%s\n", date, datum, average, sd,
				three+two+one)
		}
		average, sd = add(datum)
	}
}

// threeSigma does the classic single-sample test and returns a string
func threeSigma(datum, average, sd float64) string {
	if math.Abs(datum) > average+(3*sd) {
		if datum > 0 {
			return " 3σ"
		} else {
			return " -3σ"
		}
	}
	return ""
}

// twoSigma detects 2 out of 3 points at +/- 2 sigma
func twoSigma(datum, average, sd float64) string {
	return ""
}

// oneSigma detects  4/5 at 1 +/- sigma

// noSigma detects  9/9 on the same side of 0
