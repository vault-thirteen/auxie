package time

import "time"

// ToSecondStart returns time rounded to seconds, i.e. all the sub-second data
// is rounded to a second's start.
func ToSecondStart(t time.Time) time.Time {
	return t.Add(-time.Duration(t.Nanosecond()) * time.Nanosecond)
}

// ToMinuteStart returns time rounded to minutes, i.e. all the sub-minute data
// is rounded to a minute's start.
func ToMinuteStart(t time.Time) time.Time {
	return ToSecondStart(t).Add(-time.Duration(t.Second()) * time.Second)
}

// ToHourStart returns time rounded to hours, i.e. all the sub-hour data
// is rounded to an hour's start.
func ToHourStart(t time.Time) time.Time {
	return ToMinuteStart(t).Add(-time.Duration(t.Minute()) * time.Minute)
}

// ToDayStart returns time rounded to days, i.e. all the sub-day data
// is rounded to a day's start.
func ToDayStart(t time.Time) time.Time {
	return ToHourStart(t).Add(-time.Duration(t.Hour()) * time.Hour)
}

// ToMonthStart returns time rounded to months, i.e. all the sub-month data
// is rounded to a month's start.
func ToMonthStart(t time.Time) time.Time {
	return ToDayStart(t).Add(-time.Duration(t.Day()-1) * dayDuration)
}

// ToNextMonthStart shifts the time to the start of the next month.
func ToNextMonthStart(t time.Time) time.Time {
	return ToMonthStart(ToMonthStart(t).Add(dayDuration * 33))
}

// ToPreviousMonthStart shifts the time to the start of the previous month.
func ToPreviousMonthStart(t time.Time) time.Time {
	return ToMonthStart(ToMonthStart(t).Add(-dayDuration))
}
