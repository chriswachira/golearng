package main

import (
	"log"
	"os"

	"github.com/chriswachira/golearng/cmd"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

// Global list of apps provided by Mimir - the wise one.
var mimirApps = []string{
	"rocketlauncher",
}

func main() {

	app := &cli.App{
		Name:  "mimir",
		Usage: "Ask for an app and you shall get. He knows everything.",
		Action: func(cCtx *cli.Context) error {

			firstArg := cCtx.Args().Get(0)

			if firstArg == "" {

				log.Fatal("Mimir says: No app argument was specified, exiting...")
			}

			if !slices.Contains(mimirApps, firstArg) {

				log.Fatal("Mimir says: The app you have specified is unknown. Please run the app with --help.")

			} else {
				if firstArg == "rocketlauncher" {

					cmd.StartRocketLaunch()
				}
			}
			return nil

		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
