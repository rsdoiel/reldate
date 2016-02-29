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
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rsdoiel/reldate"
)

var (
	help          bool
	endOfMonthFor bool
	relativeTo    string
	relativeT     time.Time
)

var usage = func(exit_code int, msg string) {
	var fh = os.Stderr
	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `%s
 USAGE %s [TIME_INCREMENT TIME_UNIT|WEEKDAY_NAME]
    

 EXAMPLES
 
 Two days from today: %s 2 days
 Three weeks ago: %s -- -3 weeks
 Three weeks from 2014-01-01: %s --from=2014-01-01 3 weeks
 Three days before 2014-01-01: %s --from=2014-01-01 -- -3 days
 The Friday of this week: %s Friday
 The Monday in week containing 2015-02-06: %s --from=2015-02-06 Monday

 Time increments are a positive or negative integer. Time unit can be
 either day(s), week(s), month(s), or year(s). Weekday names are
 case insentive (e.g. Monday and monday). They can be abbreviated
 to the first three letters of the name, e.g. Sunday can be Sun, Monday
 can be Mon, Tuesday can be Tue, Wednesday can be Wed, Thursday can
 be Thu, Friday can be Fri or Saturday can be Sat.

 OPTIONS

`, msg, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])

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
		endOfMonthUsage = "Display the end of month day. E.g. 2012-02-29"
	)

	flag.StringVar(&relativeTo, "from", relativeTo, relativeToUsage)
	flag.StringVar(&relativeTo, "f", relativeTo, relativeToUsage)
	flag.BoolVar(&endOfMonthFor, "end-of-month", endOfMonthFor, endOfMonthUsage)
	flag.BoolVar(&help, "help", help, helpUsage)
	flag.BoolVar(&help, "h", help, helpUsage)
}

func assertOk(e error, failMsg string) {
	if e != nil {
		usage(1, fmt.Sprintf(" %s\n %s\n", failMsg, e))
	}
}

func main() {
	var (
		err        error
		unitString string
	)

	flag.Parse()
	if help == true {
		usage(0, "")
	}

	argc := flag.NArg()
	argv := flag.Args()

	if argc < 1 && endOfMonthFor == false {
		usage(1, "Missing time increment and units (e.g. +2 days) or weekday name (e.g. Monday, Mon).\n")
	} else if argc > 2 {
		usage(1, "Too many command line arguments.\n")
	}

	relativeT = time.Now()
	if relativeTo != "" {
		relativeT, err = time.Parse(reldate.YYYYMMDD, relativeTo)
		assertOk(err, "Cannot parse the from date.\n")
	}

	if endOfMonthFor == true {
		fmt.Println(reldate.EndOfMonth(relativeT))
		os.Exit(0)
	}

	timeInc := 0
	if argc == 2 {
		unitString = strings.ToLower(argv[1])
		timeInc, err = strconv.Atoi(argv[0])
		assertOk(err, "Time increment should be a positive or negative integer.\n")
	} else {
		// We may have a weekday string
		unitString = strings.ToLower(argv[0])
	}
	t, err := reldate.RelativeTime(relativeT, timeInc, unitString)
	assertOk(err, "Did not understand command.")
	fmt.Println(t.Format(reldate.YYYYMMDD))
}
