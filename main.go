package main

func main() {
	legacy := &LegacyPrinter{}
	adapter := &LegacyPrinterAdapter{OldPrinter: legacy}
	adapter.Print("Hello World")
}
