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
	help        bool
	relative_to string
	relative_t  time.Time
	time_inc    int
	time_unit   int
)

var Usage = func(exit_code int, msg string) {
	fmt.Fprintf(os.Stderr, `%s
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

	flag.PrintDefaults()

	fmt.Fprintf(os.Stderr, `

 copyright (c) 2014 all rights reserved.
 Released under the Simplified BSD License
 See: http://opensource.org/licenses/bsd-license.php

`)
	os.Exit(exit_code)
}

func init() {
	const (
		relative_to_usage = "Date the relative time is calculated from."
		help_usage        = "Display this help document."
	)

	flag.StringVar(&relative_to, "from", relative_to, relative_to_usage)
	flag.StringVar(&relative_to, "f", relative_to, relative_to_usage)
	flag.BoolVar(&help, "help", help, help_usage)
	flag.BoolVar(&help, "h", help, help_usage)
}

func assertOk(e error, fail_msg string) {
	if e != nil {
		Usage(1, fmt.Sprintf(" %s\n %s\n", fail_msg, e))
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
		Usage(0, "")
	}

	argc := flag.NArg()
	argv := flag.Args()

	if argc < 2 {
		Usage(1, "Missing time increment or units.")
	} else if argc > 2 {
		Usage(1, "Too many command line arguments.")
	}

	relative_t = time.Now()
	if relative_to != "" {
		relative_t, err = time.Parse(yyyymmdd, relative_to)
		assertOk(err, "Cannot parse the from date.")
	}
	time_inc, err := strconv.Atoi(argv[0])
	assertOk(err, "Time increment should be a positive or negative integer")
	t, err := relativeTime(relative_t, time_inc, argv[1])
	assertOk(err, "Time unit should be either day(s), week(s), month(s) or year(s)")

	fmt.Println(t.Format(yyyymmdd))
}
