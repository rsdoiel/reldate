/**
 * Generates a date in YYYY-MM-DD format based on a relative time
 * description (e.g. -1 week, +3 years)
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2014 all rights reserved.
 * Released under the Simplified BSD License
 * See: http://opensource.org/licenses/bsd-license.php
 */
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	help       bool
	relativeTo string
	relativeT  time.Time
	timeInc    int
	timeUnit   int
)

var usage = func(exit_code int, msg string) {
	var fh = os.Stderr
	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `%s
 USAGE %s TIME_INCREMENT TIME_UNIT

 EXAMPLES
 
 Two days from today: %s 2 days
 Three weeks ago: %s -- -3 weeks
 Three weeks from 2014-01-01: %s --from=2014-01-01 3 weeks
 Three days before 2014-01-01: %s --from=2014-01-01 -- -3 days

 Time increments are a positive or negative integer. Time unit can be
 either day(s), week(s), month(s), or year(s).

 OPTIONS

`, msg, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(fh, "\t-%s\t(defaults to %s) %s\n", f.Name, f.Value, f.Usage)
	})

	fmt.Fprintf(fh, `

 copyright (c) 2014 all rights reserved.
 Released under the Simplified BSD License
 See: http://opensource.org/licenses/bsd-license.php

`)
	os.Exit(exit_code)
}

func init() {
	const (
		relativeToUsage = "Date the relative time is calculated from."
		helpUsage       = "Display this help document."
	)

	flag.StringVar(&relativeTo, "from", relativeTo, relativeToUsage)
	flag.StringVar(&relativeTo, "f", relativeTo, relativeToUsage)
	flag.BoolVar(&help, "help", help, helpUsage)
	flag.BoolVar(&help, "h", help, helpUsage)
}

func assertOk(e error, failMsg string) {
	if e != nil {
		usage(1, fmt.Sprintf(" %s\n %s\n", failMsg, e))
	}
}

func relativeTime(t time.Time, i int, u string) (time.Time, error) {
	switch {
	case strings.HasPrefix(u, "day"):
		return t.AddDate(0, 0, i), nil
	case strings.HasPrefix(u, "week"):
		return t.AddDate(0, 0, 7*i), nil
	case strings.HasPrefix(u, "month"):
		return t.AddDate(0, i, 0), nil
	case strings.HasPrefix(u, "year"):
		return t.AddDate(i, 0, 0), nil

	}
	return t, errors.New("Time unit must be day(s), week(s), month(s) or year(s).")
}

func main() {
	const yyyymmdd = "2006-01-02"
	var err error

	//FIXME: How do I ignore -# so negative values are easy to process
	// as time increments.
	flag.Parse()
	if help == true {
		usage(0, "")
	}

	argc := flag.NArg()
	argv := flag.Args()

	if argc < 2 {
		usage(1, "Missing time increment or units.")
	} else if argc > 2 {
		usage(1, "Too many command line arguments.")
	}

	relativeT = time.Now()
	if relativeTo != "" {
		relativeT, err = time.Parse(yyyymmdd, relativeTo)
		assertOk(err, "Cannot parse the from date.")
	}
	timeInc, err := strconv.Atoi(argv[0])
	assertOk(err, "Time increment should be a positive or negative integer")
	t, err := relativeTime(relativeT, timeInc, argv[1])
	assertOk(err, "Time unit should be either day(s), week(s), month(s) or year(s)")

	fmt.Println(t.Format(yyyymmdd))
}
