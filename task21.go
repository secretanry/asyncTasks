package main

import "fmt"

type Printer interface {
	Print(msg string) string
}

type LegacyPrinter struct{}

func (l *LegacyPrinter) PrintLegacy(msg string) string {
	result := "Legacy: " + msg
	fmt.Println(result)
	return result
}

type LegacyPrinterAdapter struct {
	OldPrinter *LegacyPrinter
}

func (a *LegacyPrinterAdapter) Print(msg string) string {
	return a.OldPrinter.PrintLegacy(msg)
}
