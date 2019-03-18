//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("hostsman-portable", "HostsMan"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "hm.exe")

	// Create data folders
	progDataPath := utl.CreateFolder(app.DataPath, "ProgramData", "abelhadigital.com", "HostsMan")
	userProfilePath := utl.CreateFolder(app.DataPath, "UserProfile", "AppData", "Roaming", "abelhadigital.com", "HostsMan")
	Log.Info().Msgf("Prog data path: %s", progDataPath)
	Log.Info().Msgf("User profile path: %s", userProfilePath)

	utl.OverrideEnv("SystemDrive", app.DataPath)
	utl.OverrideEnv("USERPROFILE", utl.PathJoin(app.DataPath, "UserProfile"))

	app.Launch(os.Args[1:])
}
