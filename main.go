//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/utl"
)

var (
	app *portapps.App
)

func init() {
	var err error

	// Init app
	if app, err = portapps.New("hostsman-portable", "HostsMan"); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "hm.exe")

	// Create data folders
	progDataPath := utl.CreateFolder(app.DataPath, "ProgramData", "abelhadigital.com", "HostsMan")
	userProfilePath := utl.CreateFolder(app.DataPath, "UserProfile", "AppData", "Roaming", "abelhadigital.com", "HostsMan")
	log.Info().Msgf("Prog data path: %s", progDataPath)
	log.Info().Msgf("User profile path: %s", userProfilePath)

	os.Setenv("SystemDrive", app.DataPath)
	os.Setenv("USERPROFILE", utl.PathJoin(app.DataPath, "UserProfile"))

	defer app.Close()
	app.Launch(os.Args[1:])
}
