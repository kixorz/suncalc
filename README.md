# SunCalc 🌞

A lightweight Go library for calculating sun/moon positions and phases.

**Current Version: 1.0.0**

## Overview

SunCalc is a BSD-licensed Go library that provides precise calculations for:

- 🌅 Sun position and sunlight phases (sunrise, sunset, golden hour, etc.)
- 🌙 Moon position and lunar phase
- ⏱️ Time-based calculations for astronomical events

Originally created in JavaScript by [Vladimir Agafonkin](http://agafonkin.com/en) ([@mourner](https://github.com/mourner)) as part of the [SunCalc.net project](http://suncalc.net), and translated to Go by Douglas Six.

## Installation

```bash
go get github.com/kixorz/suncalc@v1.0.0
```

Or for the latest version:

```bash
go get github.com/kixorz/suncalc
```

## Usage Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "time"
    "github.com/kixorz/suncalc"
)

func main() {
    // Example coordinates (latitude, longitude)
    lat, lng := 51.5, -0.1  // London
    
    // Get today's sunlight times
    now := time.Now()
    times := suncalc.GetTimes(now, lat, lng, 0, time.UTC)
    
    // Print sunrise and sunset times
    fmt.Printf("Sunrise: %s\n", times[suncalc.Sunrise].Value.Format("15:04:05"))
    fmt.Printf("Sunset: %s\n", times[suncalc.Sunset].Value.Format("15:04:05"))
    
    // Get current sun position
    pos := suncalc.GetPosition(now, lat, lng)
    fmt.Printf("Sun altitude: %.2f°, azimuth: %.2f°\n", 
        pos.Altitude * 180 / math.Pi, 
        pos.Azimuth * 180 / math.Pi)
}
```

### Using SunCalculator

The SunCalculator provides a simpler API for day phase calculations:

```go
// Create a calculator for a specific location and time
sc := suncalc.NewSunCalculator(latitude, longitude, time.Now())

// Get the current phase of the day
phase := sc.GetCurrentPhase()
fmt.Printf("Current phase: %s\n", phase)

// Check if it's currently day or night
if sc.IsDay() {
    fmt.Println("It's daytime!")
} else if sc.IsNight() {
    fmt.Println("It's nighttime!")
}

// Get moon illumination
moonIllum := sc.GetMoonIllumination()
fmt.Printf("Moon illumination: %.1f%%\n", moonIllum.Fraction * 100)
```

## Features

### Sun Calculations

- **Sun Position**: Get altitude and azimuth for any location and time
- **Sunlight Phases**: Calculate precise times for:
  - 🌄 Sunrise and sunset
  - ✨ Golden hour (soft light ideal for photography)
  - 🌆 Civil, nautical, and astronomical twilight
  - ☀️ Solar noon and nadir

### Moon Calculations

- **Moon Position**: Get altitude, azimuth, and distance
- **Moon Illumination**: Calculate phase and illuminated fraction
- **Moon Times**: Calculate moonrise and moonset times

## API Reference

### Sun Times

```go
suncalc.GetTimes(date time.Time, latitude float64, longitude float64, height float64, location *time.Location)
```

Returns a map with the following properties (each is a `DayTime` object):

| Property | Description |
|----------|-------------|
| `Sunrise` | Sunrise (top edge of the sun appears on the horizon) |
| `SunriseEnd` | Sunrise ends (bottom edge of the sun touches the horizon) |
| `GoldenHourEnd` | Morning golden hour ends |
| `SolarNoon` | Solar noon (sun is in the highest position) |
| `GoldenHour` | Evening golden hour starts |
| `SunsetStart` | Sunset starts (bottom edge of the sun touches the horizon) |
| `Sunset` | Sunset (sun disappears below the horizon) |
| `Dusk` | Dusk (evening nautical twilight starts) |
| `NauticalDusk` | Nautical dusk (evening astronomical twilight starts) |
| `Night` | Night starts (dark enough for astronomical observations) |
| `Nadir` | Nadir (darkest moment of the night) |
| `NightEnd` | Night ends (morning astronomical twilight starts) |
| `NauticalDawn` | Nautical dawn (morning nautical twilight starts) |
| `Dawn` | Dawn (morning civil twilight starts) |

### Binary Phase Methods

The `SunCalculator` provides binary methods to easily check if the current time is in a specific phase:

| Method | Description |
|--------|-------------|
| `IsNight()` | Returns true if current time is during night (sun is more than 18° below horizon) |
| `IsAstroDawn()` | Returns true if current time is during astronomical dawn (sun is between 18° and 12° below horizon) |
| `IsNauticalDawn()` | Returns true if current time is during nautical dawn (sun is between 12° and 6° below horizon) |
| `IsCivilDawn()` | Returns true if current time is during civil dawn (sun is between 6° below horizon and sunrise) |
| `IsSunrise()` | Returns true if current time is during sunrise (sun is rising above the horizon) |
| `IsGoldenMorning()` | Returns true if current time is during morning golden hour |
| `IsDay()` | Returns true if current time is during full daylight |
| `IsGoldenEvening()` | Returns true if current time is during evening golden hour |
| `IsSunset()` | Returns true if current time is during sunset (sun is setting below the horizon) |
| `IsCivilDusk()` | Returns true if current time is during civil dusk (sun is between sunset and 6° below horizon) |
| `IsNauticalDusk()` | Returns true if current time is during nautical dusk (sun is between 6° and 12° below horizon) |
| `IsAstroDusk()` | Returns true if current time is during astronomical dusk (sun is between 12° and 18° below horizon) |

## License

BSD License

## Credits

- Original JavaScript implementation: [Vladimir Agafonkin](http://agafonkin.com/en)
- Go port: Douglas Six
- Astronomical formulas: [Astronomy Answers](http://aa.quae.nl/)