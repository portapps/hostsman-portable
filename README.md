<p align="center"><a href="https://github.com/crazy-max/HostsManPortable" target="_blank"><img width="100" src="https://github.com/crazy-max/HostsManPortable/blob/master/res/logo.png"></a></p>

<p align="center">
  <a href="https://github.com/crazy-max/HostsManPortable/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/HostsManPortable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/HostsManPortable/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/HostsManPortable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/crazy-max/HostsManPortable"><img src="https://img.shields.io/appveyor/ci/crazy-max/HostsManPortable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://goreportcard.com/report/github.com/crazy-max/HostsManPortable"><img src="https://goreportcard.com/badge/github.com/crazy-max/HostsManPortable?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/HostsManPortable"><img src="https://img.shields.io/codacy/grade/f8c77f7d45d34409b16d3b957ef80cf7.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=EZWH7N7DYDECQ"><img src="https://img.shields.io/badge/donate-paypal-blue.svg?style=flat-square" alt="Donate Paypal"></a>
  <a href="https://flattr.com/submit/auto?user_id=crazymax&url=https://github.com/crazy-max/HostsManPortable"><img src="https://img.shields.io/badge/flattr-this-green.svg?style=flat-square" alt="Flattr this!"></a>
</p>

## About

A single EXE written in [Go](https://golang.org/) to make [HostsMan](http://www.abelhadigital.com/hostsman) portable on Windows systems.<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

## Installation

There are four different kinds of artifacts :

* `HostsManPortable-x.x.x-x-setup.exe` : Full portable release of HostsMan as a setup. **Recommended way**!
* `HostsManPortable-x.x.x-x.{7z,zip}` : Full portable release of HostsMan as an archive.
* `HostsManPortable.exe` : Only the portable binary (must be dropped in the root folder near `hm.exe`)
* `HostsMan_x.x.x.exe` : The original release from the [official website](http://www.abelhadigital.com/hostsman).

For a **fresh installation**, install `HostsManPortable-x.x.x-x-setup.exe` where you want then run `HostsManPortable.exe`.

If **you have already installed HostsMan from the original setup or zip file**, do the same thing as a fresh installation and run `HostsManPortable.exe` a first time.<br />
The data located in `%APPDATA%\abelhadigital.com\HostsMan` and `%PROGRAMDATA%\abelhadigital.com\HostsMan` will be moved to their respectives folders.<br />
Then you can remove HostsMan from your computer.

**For an upgrade**, simply download and install the [latest setup](https://github.com/crazy-max/HostsManPortable/releases/latest).

## How can i help ?

We welcome all kinds of contributions :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
Any funds donated will be used to help further development on this project! :gift_heart:

<p>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=EZWH7N7DYDECQ">
    <img src="https://github.com/crazy-max/HostsManPortable/blob/master/res/paypal.png" alt="Donate Paypal">
  </a>
  <a href="https://flattr.com/submit/auto?user_id=crazymax&url=https://github.com/crazy-max/HostsManPortable">
    <img src="https://github.com/crazy-max/HostsManPortable/blob/master/res/flattr.png" alt="Flattr this!">
  </a>
</p>

## License

MIT. See `LICENSE` for more details.
