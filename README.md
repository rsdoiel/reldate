reldate
=======

A small command line utility which returns the relative date in YYYY-MM-DD format. This is helpful
when scripting various time relationships.

## Example

If today was 2014-08-03 and you wanted the date three days in the past try--

```
    reldate 3 days
```

The output would be 

```
    2014-08-06
```

## Time units

Supported time units are

+ day(s)
+ week(s)
+ year(s)

## Specifying a date to calucate from

_reldate_ handles dates in the YYYY-MM-DD format (e.g. March 1, 2014 would be 2014-03-01).  By default _reldate_ uses today as
the date to calculate relative time from.  If you use the *--from* option you can it will calculate the relative date from that 
specific date. 

```
    reldate --from=2014-08-03 3 days
```

Will yield

```
    2014-08-06
```

## Negative increments

Command line arguments traditionally start with a dash which we also use to denote a nagative number. To tell the command line
process that to not treat negative numbers as an "option" preceed your time increment and time unit with a double dash.

```
   reldate --from=2014-08-03 -- -3 days 
```

Will yield

```
    2014-07-31
```


