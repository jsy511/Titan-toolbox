package ui

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func LiveDashboard() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Println("║                 TITAN DASHBOARD                  ║")
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Printf("║ OS:       %-39s ║\n", runtime.GOOS)
	fmt.Printf("║ Arch:     %-39s ║\n", runtime.GOARCH)
	fmt.Printf("║ CPU Cores: %-38d ║\n", runtime.NumCPU())
	fmt.Printf("║ Go:       %-39s ║\n", runtime.Version())
	fmt.Printf("║ PID:      %-39d ║\n", os.Getpid())
	fmt.Printf("║ Time:     %-39s ║\n", time.Now().Format("15:04:05"))
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Println("║                MODULE STATUS                     ║")
	fmt.Println("║                                                  ║")
	fmt.Println("║  ● Titan Core       ONLINE                       ║")
	fmt.Println("║  ● Titan Network    READY                        ║")
	fmt.Println("║  ● Titan DNS        READY                        ║")
	fmt.Println("║  ● Titan System     ONLINE                       ║")
	fmt.Println("║  ● Titan Tools      READY                        ║")
	fmt.Println("║  ● Titan Lab        READY                        ║")
	fmt.Println("║  ● Titan CTF        READY                        ║")
	fmt.Println("║  ● Titan Reports    READY                        ║")
	fmt.Println("╚══════════════════════════════════════════════════╝")
}