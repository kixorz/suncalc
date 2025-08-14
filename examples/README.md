# SunCalc Example

This directory contains an example application that demonstrates how to use the SunCalc library.

## Running the Example

To run the example, navigate to the examples directory and use the `go run` command:

```bash
cd examples
go run main.go
```

## What the Example Does

The example demonstrates the following features of the SunCalc library:

1. Creating a new SunCalculator for a specific location (San Francisco)
2. Getting all sun times for the day
3. Determining the current day phase
4. Checking specific day phases (night, dawn, day, etc.)
5. Getting the sun position (azimuth and altitude)
6. Getting the moon position (azimuth, altitude, distance, and parallactic angle)
7. Getting the moon illumination (fraction, phase, and angle)
8. Getting the moon rise and set times

## Example Output

The example will output information similar to the following:

```
Sun times for today:
  dawn: 12:51:15
  nauticalDusk: 04:16:53
  night: 04:54:04
  goldenHourEnd: 13:56:20
  solarNoon: 20:16:43
  sunsetStart: 03:10:44
  dusk: 03:42:11
  nauticalDawn: 12:16:33
  nightEnd: 11:39:22
  goldenHour: 02:37:06
  nadir: 08:16:43
  sunrise: 13:19:49
  sunset: 03:13:37
  sunriseEnd: 13:22:43

Current phase: day

Binary phase checks:
  IsNight: false
  IsAstroDawn: false
  IsNauticalDawn: false
  IsCivilDawn: false
  IsSunrise: false
  IsGoldenMorning: false
  IsDay: true
  IsGoldenEvening: false
  IsSunset: false
  IsCivilDusk: false
  IsNauticalDusk: false
  IsAstroDusk: false

Sun position:
  Azimuth: -81.65° (-1.43 rad)
  Altitude: 37.45° (0.65 rad)

Moon position:
  Azimuth: 103.56° (1.81 rad)
  Altitude: -57.96° (-1.01 rad)
  Distance: 390358.06 km
  Parallactic Angle: 58.09° (1.01 rad)

Moon illumination:
  Fraction: 0.97
  Phase: 0.44
  Angle: -110.45° (-1.93 rad)

Moon times:
  Rise: 03:51:48
  Set: 13:14:38
  Always Up: false
  Always Down: false
```

Note that the actual values will vary depending on the current date and time.