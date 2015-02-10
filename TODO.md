
# Someday, maybe

+ weekday name relative current date or YYYY-MM-DD dates
    + reldate Tuesday # would return the YYYY-MM-DD date for Tuesday in the current week
    + reldate --from=2015-02-01 Monday # this would return 2015-02-02 reflecting the Monday of the week (Sun-Sat) containing 2015-02-01
+ weekday integer (e.g. Sunday is 0 Saturday is 6)
    + reldate weekday # this would return 0 though 6 corresponding to the current day in the week (Sunday is 0, Saturday is 6)
    + reldate --from=2015-02-02 weekday # this would return 1 for Monday since 2015-02-02 is a Monday (where Sunday is 0 and Saturday is 6)
+ weekday names relative to month based on current date or YYYY-MM-DD dates
    + reldate --from=2015-02-10 first Monday # this would yield 2015-02-02
    + reldate --from=2015-02-10 current Monday # this would yield 2015-02-09
    + reldate --form=2015-02-10 next Monday # this would yield 2015-02-16
    + reldate --form=2015-02-10 previous Monday # this would yield 2015-02-02
    + reldate --from=2015-02-10 last Monday # this would yield 2015-02-23
    + reldate --from=2015-02-10 second Tuesday # this would yield 2015-02-10
    + reldate --from=2015-02-10 third Thursday # this would yield 2015-02-16
    + reldate --from=2015-02-10 fourth Sunday # this would yield 2015-02-22
    + reldate --from=2015-02-10 fifth Monday # this would yield an empty string and exit with an error
    + reldate --from=2015-06-19 fifth Monday # this would yield 2015-06-29


