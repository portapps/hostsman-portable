//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	. "github.com/portapps/portapps"
)

func init() {
	App.ID = "hostsman-portable"
	App.Name = "HostsMan"
	Init()
}

func main() {
	//App.MainPath =
	App.DataPath = CreateFolder(App.DataPath)
	App.Process = RootPathJoin("hm.exe")
	App.Args = nil
	App.WorkingDir = App.MainPath

	// Create data folders
	progDataPath := CreateFolder(PathJoin(App.RootDataPath, "ProgramData", "abelhadigital.com", "HostsMan"))
	userProfilePath := CreateFolder(PathJoin(App.RootDataPath, "UserProfile", "AppData", "Roaming", "abelhadigital.com", "HostsMan"))
	Log.Infof("Prog data path: %s", progDataPath)
	Log.Infof("User profile path: %s", userProfilePath)

	OverrideEnv("SystemDrive", App.DataPath)
	OverrideEnv("ProgramData", PathJoin(App.DataPath, "ProgramData"))
	OverrideEnv("ALLUSERSPROFILE", PathJoin(App.DataPath, "ProgramData"))
	OverrideEnv("USERPROFILE", PathJoin(App.DataPath, "UserProfile"))
	OverrideEnv("APPDATA", PathJoin(App.DataPath, "UserProfile", "Roaming"))
	Launch()
}
