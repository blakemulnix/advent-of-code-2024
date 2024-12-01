package format

import (
	"fmt"
)

const (
	colorWhite  = "\033[1;37m"
	colorGreen  = "\033[1;32m"
	colorYellow = "\033[1;33m"
	colorReset  = "\033[0m"
	decorBold      = "\033[1m"
	decorUnderline = "\033[4m"
	decorItalic    = "\033[3m"
) 

func PrintHeader() {
	fmt.Printf("🎄 %sAdvent of Code 2024%s 🎄\n\n",
		colorGreen, colorReset)
}

func PrintFooter() {
	fmt.Printf("%s❄️  %sHappy Coding!%s %s❄️%s\n\n",
		colorWhite, colorGreen, colorWhite, colorWhite, colorReset)
}

func PrintAnswer(day int, part int, answer any, duration string) {
	fmt.Printf("   %s%s%sDay %d → Part %d%s\n", 
		colorGreen, decorBold, decorUnderline, day, part, colorReset)
	fmt.Printf("⭐️ %sAnswer: %s%v%s\n",
		colorGreen, colorYellow, answer, colorReset)
	fmt.Printf("⏱️  %sTime: %s%s%s\n\n",
		colorGreen, colorYellow, duration, colorReset)
}
