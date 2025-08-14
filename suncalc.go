package suncalc

import (
	"time"
)

// Usage:
//
//	// Create a new SunCalculator for a specific location and date
//	sc := NewSunCalculator(latitude, longitude, time.Now())
//
//	// Get the current day phase
//	phase := sc.GetCurrentPhase()
//
//	// Check if it's currently a specific phase
//	if sc.IsGoldenHour() {
//	    // Take golden hour photos
//	}
//
//	// Get all calculated sun times for the day
//	times := sc.GetAllTimes()

// DayPhase represents different phases of the day
// Each phase corresponds to a specific period of the day based on the sun's position
type DayPhase string

const (
	// PhaseNight represents full darkness when the sun is far below the horizon
	PhaseNight DayPhase = "night" // Full night time

	// PhaseAstroDawn represents the period when the sun is between 18° and 12° below the horizon (morning)
	PhaseAstroDawn DayPhase = "astroDawn" // Astronomical dawn (night ends)

	// PhaseNauticalDawn represents the period when the sun is between 12° and 6° below the horizon (morning)
	PhaseNauticalDawn DayPhase = "nauticalDawn" // Nautical dawn

	// PhaseCivilDawn represents the period when the sun is between 6° below the horizon and sunrise (morning twilight)
	PhaseCivilDawn DayPhase = "civilDawn" // Civil dawn (morning twilight)

	// PhaseSunrise represents the period when the sun is rising above the horizon
	PhaseSunrise DayPhase = "sunrise" // Sunrise period

	// PhaseGoldenMorning represents the period shortly after sunrise with soft, warm light (morning golden hour)
	PhaseGoldenMorning DayPhase = "goldenMorning" // Morning golden hour

	// PhaseDay represents the period of full daylight between the golden hours
	PhaseDay DayPhase = "day" // Full daylight

	// PhaseGoldenEvening represents the period shortly before sunset with soft, warm light (evening golden hour)
	PhaseGoldenEvening DayPhase = "goldenEvening" // Evening golden hour

	// PhaseSunset represents the period when the sun is setting below the horizon
	PhaseSunset DayPhase = "sunset" // Sunset period

	// PhaseCivilDusk represents the period when the sun is between sunset and 6° below the horizon (evening twilight)
	PhaseCivilDusk DayPhase = "civilDusk" // Civil dusk (evening twilight)

	// PhaseNauticalDusk represents the period when the sun is between 6° and 12° below the horizon (evening)
	PhaseNauticalDusk DayPhase = "nauticalDusk" // Nautical dusk

	// PhaseAstroDusk represents the period when the sun is between 12° and 18° below the horizon (evening)
	PhaseAstroDusk DayPhase = "astroDusk" // Astronomical dusk (night starts)
)

// SunCalculator provides methods to determine sun positions and day phases
// It wraps the suncalc library to provide a simpler API for day phase calculations
// and binary getter methods for each phase.
type SunCalculator struct {
	latitude  float64                 // Latitude in decimal degrees (positive for north, negative for south)
	longitude float64                 // Longitude in decimal degrees (positive for east, negative for west)
	date      time.Time               // Date for which sun calculations are performed
	times     map[DayTimeName]DayTime // Cached sun times for the day
}

// NewSunCalculator creates a new SunCalculator for the given coordinates and date
// Parameters:
//   - latitude: Latitude in decimal degrees (positive for north, negative for south)
//   - longitude: Longitude in decimal degrees (positive for east, negative for west)
//   - date: Date for which to calculate sun times (the time component is used for the day)
//
// Returns a pointer to a new SunCalculator with pre-calculated sun times for the specified day
func NewSunCalculator(latitude, longitude float64, date time.Time) *SunCalculator {
	sc := &SunCalculator{
		latitude:  latitude,
		longitude: longitude,
		date:      date,
	}

	// Calculate all sun times for the day
	sc.times = GetTimes(date, latitude, longitude)

	return sc
}

// GetCurrentPhase determines the current day phase based on the current time
// This method compares the current time with the pre-calculated sun times
// to determine which phase of the day it currently is.
//
// Returns a DayPhase constant representing the current phase of the day
func (sc *SunCalculator) GetCurrentPhase() DayPhase {
	now := time.Now()

	// Check if we're before sunrise (night or dawn phases)
	if now.Before(sc.times[NightEnd].Value) {
		return PhaseNight
	} else if now.Before(sc.times[NauticalDawn].Value) {
		return PhaseAstroDawn
	} else if now.Before(sc.times[Dawn].Value) {
		return PhaseNauticalDawn
	} else if now.Before(sc.times[Sunrise].Value) {
		return PhaseCivilDawn
	} else if now.Before(sc.times[SunriseEnd].Value) {
		return PhaseSunrise
	} else if now.Before(sc.times[GoldenHourEnd].Value) {
		return PhaseGoldenMorning
	}

	// Check if we're after sunset (dusk or night phases)
	if now.After(sc.times[Night].Value) {
		return PhaseNight
	} else if now.After(sc.times[NauticalDusk].Value) {
		return PhaseAstroDusk
	} else if now.After(sc.times[Dusk].Value) {
		return PhaseNauticalDusk
	} else if now.After(sc.times[Sunset].Value) {
		return PhaseCivilDusk
	} else if now.After(sc.times[SunsetStart].Value) {
		return PhaseSunset
	} else if now.After(sc.times[GoldenHour].Value) {
		return PhaseGoldenEvening
	}

	// If none of the above, it's full daylight
	return PhaseDay
}

// Binary getter methods for each phase

// IsNight returns true if current time is during night
// Night is the period when the sun is far below the horizon (more than 18°)
func (sc *SunCalculator) IsNight() bool {
	return sc.GetCurrentPhase() == PhaseNight
}

// IsAstroDawn returns true if current time is during astronomical dawn
// Astronomical dawn is when the sun is between 18° and 12° below the horizon in the morning
func (sc *SunCalculator) IsAstroDawn() bool {
	return sc.GetCurrentPhase() == PhaseAstroDawn
}

// IsNauticalDawn returns true if current time is during nautical dawn
// Nautical dawn is when the sun is between 12° and 6° below the horizon in the morning
func (sc *SunCalculator) IsNauticalDawn() bool {
	return sc.GetCurrentPhase() == PhaseNauticalDawn
}

// IsCivilDawn returns true if current time is during civil dawn
// Civil dawn is when the sun is between 6° below the horizon and sunrise
func (sc *SunCalculator) IsCivilDawn() bool {
	return sc.GetCurrentPhase() == PhaseCivilDawn
}

// IsSunrise returns true if current time is during sunrise
// Sunrise is the period when the sun is rising above the horizon
func (sc *SunCalculator) IsSunrise() bool {
	return sc.GetCurrentPhase() == PhaseSunrise
}

// IsGoldenMorning returns true if current time is during morning golden hour
// Morning golden hour is the period shortly after sunrise with soft, warm light
func (sc *SunCalculator) IsGoldenMorning() bool {
	return sc.GetCurrentPhase() == PhaseGoldenMorning
}

// IsDay returns true if current time is during full daylight
// Day is the period of full daylight between the golden hours
func (sc *SunCalculator) IsDay() bool {
	return sc.GetCurrentPhase() == PhaseDay
}

// IsGoldenEvening returns true if current time is during evening golden hour
// Evening golden hour is the period shortly before sunset with soft, warm light
func (sc *SunCalculator) IsGoldenEvening() bool {
	return sc.GetCurrentPhase() == PhaseGoldenEvening
}

// IsSunset returns true if current time is during sunset
// Sunset is the period when the sun is setting below the horizon
func (sc *SunCalculator) IsSunset() bool {
	return sc.GetCurrentPhase() == PhaseSunset
}

// IsCivilDusk returns true if current time is during civil dusk
// Civil dusk is when the sun is between sunset and 6° below the horizon
func (sc *SunCalculator) IsCivilDusk() bool {
	return sc.GetCurrentPhase() == PhaseCivilDusk
}

// IsNauticalDusk returns true if current time is during nautical dusk
// Nautical dusk is when the sun is between 6° and 12° below the horizon in the evening
func (sc *SunCalculator) IsNauticalDusk() bool {
	return sc.GetCurrentPhase() == PhaseNauticalDusk
}

// IsAstroDusk returns true if current time is during astronomical dusk
// Astronomical dusk is when the sun is between 12° and 18° below the horizon in the evening
func (sc *SunCalculator) IsAstroDusk() bool {
	return sc.GetCurrentPhase() == PhaseAstroDusk
}

// GetAllTimes returns all calculated sun times for the day
// This provides access to the raw sun times calculated by the suncalc library
func (sc *SunCalculator) GetAllTimes() map[DayTimeName]DayTime {
	return sc.times
}

// GetSunPosition returns the position of the sun (azimuth and altitude) for the given time
// If no time is provided, it uses the current time
// Returns an object with Azimuth and Altitude properties in radians
func (sc *SunCalculator) GetSunPosition(t ...time.Time) SunPosition {
	timeToUse := time.Now()
	if len(t) > 0 {
		timeToUse = t[0]
	}
	return GetPosition(timeToUse, sc.latitude, sc.longitude)
}

// GetMoonPosition returns the position of the moon for the given time
// If no time is provided, it uses the current time
// Returns an object with Altitude, Azimuth, Distance and ParallacticAngle properties
func (sc *SunCalculator) GetMoonPosition(t ...time.Time) MoonPosition {
	timeToUse := time.Now()
	if len(t) > 0 {
		timeToUse = t[0]
	}
	return GetMoonPosition(timeToUse, sc.latitude, sc.longitude)
}

// GetMoonIllumination returns information about the moon's illumination for the given time
// If no time is provided, it uses the current time
// Returns an object with Fraction, Phase, and Angle properties
func (sc *SunCalculator) GetMoonIllumination(t ...time.Time) MoonIllumination {
	timeToUse := time.Now()
	if len(t) > 0 {
		timeToUse = t[0]
	}
	return GetMoonIllumination(timeToUse)
}

// GetMoonTimes returns the moon rise and set times for the given date
// If no date is provided, it uses the current date
// If inUTC is true, it will search the specified date from 0 to 24 UTC hours
// Returns an object with Rise, Set, AlwaysUp, and AlwaysDown properties
func (sc *SunCalculator) GetMoonTimes(date time.Time) MoonTimes {
	return GetMoonTimes(date, sc.latitude, sc.longitude)
}
