package main

import "github.com/SomniaStellarum/StellarUtilities/slog"

func main() {
	i := 1
	slog.DebugPrint("This is Debug information ", i)
	slog.SetDebug()
	slog.DebugPrint("This is more debug information ", i)
	slog.InfoPrint("This is informational logging ", i)
}
