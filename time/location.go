package time

import "time"

// GetLocationOffset gets the time offset of a location as duration.
func GetLocationOffset(location *time.Location) (timeOffset time.Duration, err error) {
	// Time in location.
	var timeInLocation time.Time
	timeInLocation, err = time.ParseInLocation(TimeFormatFor_GetLocationOffset, sampleTimeString, location)
	if err != nil {
		return timeOffset, err
	}

	// Time in UTC time zone.
	var timeInUtcTimezone time.Time
	timeInUtcTimezone, err = time.Parse(TimeFormatFor_GetLocationOffset, sampleTimeString)
	if err != nil {
		return timeOffset, err
	}

	// Delta.
	return timeInUtcTimezone.Sub(timeInLocation), nil
}

// ParseDayTimeStringInLocation parses time string in the location.
func ParseDayTimeStringInLocation(dayTimeString string, location *time.Location) (dayStartTime time.Time, err error) {
	// Unfortunately, the built-in 'ParseInLocation' function works not as it
	// could be understood from its name. So, we are implementing a true
	// 'ParseInLocation' method here ...

	// Get location's time zone offset.
	var locationOffset time.Duration
	locationOffset, err = GetLocationOffset(location)
	if err != nil {
		return dayStartTime, err
	}

	// Parse the time and correct it.
	dayStartTime, err = time.Parse(FormatDayTimeString, dayTimeString)
	if err != nil {
		return dayStartTime, err
	}
	dayStartTime = dayStartTime.In(location).Add(-locationOffset)

	return dayStartTime, nil
}
