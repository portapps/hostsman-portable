@ECHO OFF
SETLOCAL EnableDelayedExpansion

TASKLIST /FI "IMAGENAME eq hm.exe" /FO LIST | find "hm.exe">nul
IF %ERRORLEVEL% == 0 GOTO EXIT

SET currentPath=%~dp0

REM Public data
SET hmPublicData=%PROGRAMDATA%\abelhadigital.com\HostsMan

IF EXIST %hmPublicData% RD /S /Q %hmPublicData%
MD %hmPublicData%

MKLINK %hmPublicData%\hm.cfg %currentPath%\hm.cfg
MKLINK %hmPublicData%\prefs.cfg %currentPath%\prefs.cfg
MKLINK %hmPublicData%\update.cfg %currentPath%\update.cfg

REM User data
SET hmAppData=%APPDATA%\abelhadigital.com\HostsMan

IF EXIST %hmAppData% RD /S /Q %hmAppData%
MD %hmAppData%

MKLINK %hmAppData%\hm.ini %currentPath%\hm.ini
MKLINK %hmAppData%\update.cfg %currentPath%\update2.cfg

REM Start HostsMAn
START %currentPath%\hm.exe

:EXIT
ENDLOCAL
