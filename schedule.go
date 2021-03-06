package konturtransferbot

import "time"

// Schedule contains all information on transfer departure times
type Schedule struct {
	WorkDayRouteToOffice    Route
	WorkDayRouteFromOffice  Route
	SaturdayRouteToOffice   Route
	SaturdayRouteFromOffice Route
}

// GetToOfficeText returns text representation of full schedule to office
func (s Schedule) GetToOfficeText(now time.Time) (string, string) {
	prefix := "*Геологическая → Офис*\n\n"
	suffix := "\nСубботний рейс в " + s.SaturdayRouteToOffice.String()

	timeAgnosticRoute := prefix + s.WorkDayRouteToOffice.String() + suffix
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return timeAgnosticRoute, ""
	}

	timeSensitiveRoute := prefix + s.WorkDayRouteToOffice.StringWithDivider(now) + suffix
	if timeAgnosticRoute == timeSensitiveRoute {
		return timeAgnosticRoute, ""
	}

	return timeSensitiveRoute, timeAgnosticRoute
}

// GetFromOfficeText returns text representation of full schedule from office
func (s Schedule) GetFromOfficeText(now time.Time) (string, string) {
	prefix := "*Рейсы из офиса*\n\n"
	suffix := "\nСубботний дежурный в " + s.SaturdayRouteFromOffice.String()

	timeAgnosticRoute := prefix + s.WorkDayRouteFromOffice.String() + suffix
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return timeAgnosticRoute, ""
	}

	timeSensitiveRoute := prefix + s.WorkDayRouteFromOffice.StringWithDivider(now) + suffix
	if timeAgnosticRoute == timeSensitiveRoute {
		return timeAgnosticRoute, ""
	}

	return timeSensitiveRoute, timeAgnosticRoute
}
