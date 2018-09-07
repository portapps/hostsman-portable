//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "hostsman-portable"
	Papp.Name = "HostsMan"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")
	Papp.Process = PathJoin(Papp.AppPath, "hm.exe")
	Papp.Args = nil
	Papp.WorkingDir = Papp.AppPath

	// Create data folders
	progDataPath := CreateFolder(PathJoin(Papp.DataPath, "ProgramData", "abelhadigital.com", "HostsMan"))
	userProfilePath := CreateFolder(PathJoin(Papp.DataPath, "UserProfile", "AppData", "Roaming", "abelhadigital.com", "HostsMan"))
	Log.Infof("Prog data path: %s", progDataPath)
	Log.Infof("User profile path: %s", userProfilePath)

	OverrideEnv("SystemDrive", Papp.DataPath)
	OverrideEnv("USERPROFILE", PathJoin(Papp.DataPath, "UserProfile"))
	Launch(os.Args[1:])
}
