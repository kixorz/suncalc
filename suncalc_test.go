package suncalc

import (
	"fmt"
	"time"
)

func main() {
	// Example coordinates (San Francisco)
	latitude := 37.7749
	longitude := -122.4194

	// Create a new SunCalculator for today
	today := time.Now()
	sc := NewSunCalculator(latitude, longitude, today)

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
}
