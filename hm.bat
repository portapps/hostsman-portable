@ECHO OFF
SETLOCAL EnableDelayedExpansion

TASKLIST /FI "IMAGENAME eq hm.exe" /FO LIST | find "hm.exe">nul
IF %ERRORLEVEL% == 0 GOTO EXIT

SET currentPath=%~dp0
SET hmData=%ProgramData%\abelhadigital.com\HostsMan

IF EXIST %hmData% RD /S /Q %hmData%
MD %hmData%

MKLINK %hmData%\hm.cfg %currentPath%\hm.cfg
MKLINK %hmData%\prefs.cfg %currentPath%\prefs.cfg
MKLINK %hmData%\update.cfg %currentPath%\update.cfg

START %currentPath%\hm.exe

:EXIT
ENDLOCAL
