package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func StartRocketLaunch() {

	const defaultCountdownSeconds int = 600

	countdownSecondsFromEnv, exists := os.LookupEnv("ROCKET_LAUNCH_COUNTDOWN_SECONDS")

	if exists {
		countdownSeconds, err := strconv.Atoi(countdownSecondsFromEnv)

		if err != nil {
			panic(err)
		}

		startCountdown(countdownSeconds)

	} else {

		startCountdown(defaultCountdownSeconds)

	}

	fmt.Println("\n\n🐧 Kowalski reports lift-off, we have lift-off! 🚀🚀")

}

func startCountdown(seconds int) {

	for i := seconds; i > 0; i-- {

		if i == 14 {
			fmt.Println("\n\nMain engines running... 🔥🔥🔥🔥\n\n")
		}

		fmt.Printf("🐧 Kowalski reporting countdown T-minus %s\n", strconv.Itoa(i))
		time.Sleep(time.Duration(time.Second * 1))
	}

}
