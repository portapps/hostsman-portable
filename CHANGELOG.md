# Changelog

## 4.8.106-9 (2019/02/17)

* Add checksums (portapps/portapps#11)
* Switch to TravisCI

## 4.8.106-8 (2018/02/11)

* Move ia32/x64 to win32/win64 for arch def
* Add portapp.json file
* Uncheck run app in setup

## 4.8.106-7 (2018/02/09)

* Ability to pass custom args to the portable process

## 4.8.106-6 (2017/11/21)

* Reduce dependencies to avoid heuristic detection
* Add UPX compression

## 4.8.106-5 (2017/11/20)

* Move app to a subfolder
* Remove unnecessary override of some environment variables

> :warning: **UPGRADE NOTES**
> * Delete all files and folders in root folder except `data` folder.

## 4.8.106-4 (2017/11/19)

* Upgrade to HostsMan : 4.8.106
* Move repository to [Portapps](https://github.com/portapps) org
* Switch to [Golang Dep](https://github.com/golang/dep) as dependency manager
* New logger
* Override APPDATA and PROGRAMDATA instead of using symlinks to store data
* Do not migrate old data folder
* Reduce dependencies and system calls to avoid heuristic detection
* Upgrade to Go 1.9.1

## 4.7.105-3 (2017/08/26)

* Upgrade to Go 1.9
* Add support guidelines
* Small refactoring

## 4.7.105-1 (2017/05/26)

* Upgrade to HostsMan : 4.7.105
* New LICENSE
* Upgrade project using Go

## 1.2 (2016/01/30)

* Add hm.ini and update.cfg (convert to update2.cfg) symlinks from AppData folder.

## 1.1 (2016/01/28)

* Forgot hm.cfg symlink.

## 1.0 (2016/01/28)

* Initial version based on HostsMan 4.6.103.
