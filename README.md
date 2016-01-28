# HostsMan Portable

Portable version of [HostsMan](http://www.abelhadigital.com/hostsman) freeware application by abelhadigital.com.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [About](#about)
- [Download](#download)
- [Usage](#usage)
- [Changelog](#changelog)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## About

HostsMan originally writes configuration files in the public application data folder (C:\ProgramData).<br />
I have made a batch file (``hm.bat``) encapsulated in an exe file (``HostsMan.exe``) with [Bat To Exe Converter](http://www.f2ko.de/en/b2e.php) that creates symlinks from the application data folder to the root folder of HostsMan.<br />
The program checks that ``hm.exe`` is not already launched to prevent multiple instances (not managed by original HostsMan release at this time).

## Download

[![HostsManPortable 1.0](https://img.shields.io/badge/download-HostsManPortable%201.0%20-brightgreen.svg)](https://github.com/crazy-max/HostsManPortable/releases/download/v1.0/HostsManPortable-1.0.zip)

## Usage

Extract zip file where you want and launch ``HostsMan.exe``.

## Changelog

See ``CHANGELOG.md``.

## License

LGPL. See ``LICENSE`` for more details.
