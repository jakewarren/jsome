# jsome
[![Build Status](https://github.com/jakewarren/jsome/workflows/lint/badge.svg)](https://github.com/jakewarren/jsome/actions)
[![GitHub release](http://img.shields.io/github/release/jakewarren/jsome.svg?style=flat-square)](https://github.com/jakewarren/jsome/releases])
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/jsome/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/jsome)](https://goreportcard.com/report/github.com/jakewarren/jsome)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)
> cli utility for colorizing your JSON


## Install
### Option 1: Binary

Download the latest release from [https://github.com/jakewarren/jsome/releases/latest](https://github.com/jakewarren/jsome/releases/latest)

### Option 2: From source

```
go get github.com/jakewarren/jsome
```

## Screenshot

![](screenshot.png)

## Usage

```
Usage: jsome [<flags>] [FILE]

Example:
	jsome file.json
	cat file.json | jsome 

Optional flags:

  -h, --help   display help

```

## Acknowledgements

[Javascipt/Jsome](https://github.com/Javascipt/Jsome) - for inspiration of the tool, and for the name I shamelessly copied  
[nwidger/jsoncolor](https://github.com/nwidger/jsoncolor) - awesome library that powers this tool

## Changes

All notable changes to this project will be documented in the [changelog].

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this project adheres to [Semantic Versioning](http://semver.org/).

## License

MIT © 2018 Jake Warren

[changelog]: https://github.com/jakewarren/jsome/blob/master/CHANGELOG.md
