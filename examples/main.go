package main

import (
	"fmt"
	"math"
	"time"

	"github.com/kixorz/suncalc"
)

// This is a runnable example to demonstrate the SunCalculator usage
func main() {
	// Example coordinates (Zastavka)
	latitude := 49.1908
	longitude := 16.3614

	// Create a new SunCalculator for today
	today := time.Now()
	sc := suncalc.NewSunCalculator(latitude, longitude, today)

	// Print all sun times for today
	fmt.Println("Sun times for today:")
	for name, t := range sc.GetAllTimes() {
		fmt.Printf("  %s: %s\n", name, t.Value.Format("15:04:05"))
	}

	// Get and print current phase
	currentPhase := sc.GetCurrentPhase()
	fmt.Printf("\nCurrent phase: %s\n", currentPhase)

	// Test all binary getters
	fmt.Println("\nBinary phase checks:")
	fmt.Printf("  IsNight: %t\n", sc.IsNight())
	fmt.Printf("  IsAstroDawn: %t\n", sc.IsAstroDawn())
	fmt.Printf("  IsNauticalDawn: %t\n", sc.IsNauticalDawn())
	fmt.Printf("  IsCivilDawn: %t\n", sc.IsCivilDawn())
	fmt.Printf("  IsSunrise: %t\n", sc.IsSunrise())
	fmt.Printf("  IsGoldenMorning: %t\n", sc.IsGoldenMorning())
	fmt.Printf("  IsDay: %t\n", sc.IsDay())
	fmt.Printf("  IsGoldenEvening: %t\n", sc.IsGoldenEvening())
	fmt.Printf("  IsSunset: %t\n", sc.IsSunset())
	fmt.Printf("  IsCivilDusk: %t\n", sc.IsCivilDusk())
	fmt.Printf("  IsNauticalDusk: %t\n", sc.IsNauticalDusk())
	fmt.Printf("  IsAstroDusk: %t\n", sc.IsAstroDusk())

	// Get and print sun position
	fmt.Println("\nSun position:")
	sunPos := sc.GetSunPosition()
	fmt.Printf("  Azimuth: %.2f° (%.2f rad)\n", sunPos.Azimuth*180/math.Pi, sunPos.Azimuth)
	fmt.Printf("  Altitude: %.2f° (%.2f rad)\n", sunPos.Altitude*180/math.Pi, sunPos.Altitude)

	// Get and print moon position
	fmt.Println("\nMoon position:")
	moonPos := sc.GetMoonPosition()
	fmt.Printf("  Azimuth: %.2f° (%.2f rad)\n", moonPos.Azimuth*180/math.Pi, moonPos.Azimuth)
	fmt.Printf("  Altitude: %.2f° (%.2f rad)\n", moonPos.Altitude*180/math.Pi, moonPos.Altitude)
	fmt.Printf("  Distance: %.2f km\n", moonPos.Distance)
	fmt.Printf("  Parallactic Angle: %.2f° (%.2f rad)\n", moonPos.ParallacticAngle*180/math.Pi, moonPos.ParallacticAngle)

	// Get and print moon illumination
	fmt.Println("\nMoon illumination:")
	moonIllum := sc.GetMoonIllumination()
	fmt.Printf("  Fraction: %.2f\n", moonIllum.Fraction)
	fmt.Printf("  Phase: %.2f\n", moonIllum.Phase)
	fmt.Printf("  Angle: %.2f° (%.2f rad)\n", moonIllum.Angle*180/math.Pi, moonIllum.Angle)

	// Get and print moon times
	fmt.Println("\nMoon times:")
	moonTimes := sc.GetMoonTimes(today)
	if moonTimes.Rise.IsZero() {
		fmt.Println("  Rise: No moonrise today")
	} else {
		fmt.Printf("  Rise: %s\n", moonTimes.Rise.Format("15:04:05"))
	}
	if moonTimes.Set.IsZero() {
		fmt.Println("  Set: No moonset today")
	} else {
		fmt.Printf("  Set: %s\n", moonTimes.Set.Format("15:04:05"))
	}
	fmt.Printf("  Always Up: %t\n", moonTimes.AlwaysUp)
	fmt.Printf("  Always Down: %t\n", moonTimes.AlwaysDown)
}
