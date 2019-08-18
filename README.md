# ascii2svg

:smile: ascii2svg creates .svg images from an ascii text

[![CircleCI](https://circleci.com/gh/moul/ascii2svg.svg?style=shield)](https://circleci.com/gh/moul/ascii2svg)
[![GoDoc](https://godoc.org/moul.io/ascii2svg?status.svg)](https://godoc.org/moul.io/ascii2svg)
[![License](https://img.shields.io/github/license/moul/ascii2svg.svg)](https://github.com/moul/ascii2svg/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/moul/ascii2svg.svg)](https://github.com/moul/ascii2svg/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/ascii2svg)](https://goreportcard.com/report/moul.io/ascii2svg)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/ascii2svg/badge)](https://www.codefactor.io/repository/github/moul/ascii2svg)
[![codecov](https://codecov.io/gh/moul/ascii2svg/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/ascii2svg)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/ascii2svg.svg)](https://microbadger.com/images/moul/ascii2svg)
[![Sourcegraph](https://sourcegraph.com/github.com/moul/ascii2svg/-/badge.svg)](https://sourcegraph.com/github.com/moul/ascii2svg?badge)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)


## Usage

```console
$ cat SECURITY.md | ascii2svg
<?xml version="1.0" encoding="UTF-8"?>
<svg width="800px" height="600px" viewBox="0 0 800 600" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
  <g>
    <rect x="0" y="0" width="800" height="600" fill="black" stroke="white"></rect>
    <foreignObject x="15" y="5" width="770" height="590">
      <div xmlns="http://www.w3.org/1999/xhtml" style="width:800px; height:600px; overflow-y:auto; color: white; font-size: 20px;">
        <pre># Reporting security issues

I take security seriously. If you discover a security issue, please bring it to my attention right away!

### Reporting a Vulnerability

Please **DO NOT** file a public issue, instead send your report privately to m+security-report@42.am.

Security reports are greatly appreciated and I will publicly thank you for it, although I keep your name confidential if you request it.
</pre>
      </div>
    </foreignObject>
  </g>
</svg>
```

## Install

### Using go

```console
$ go get -u moul.io/ascii2svg
```

### Using brew

```console
$ brew install moul/moul/ascii2svg
```

### Download releases

https://github.com/moul/ascii2svg/releases

## License

© 2019 [Manfred Touron](https://manfred.life) -
[Apache-2.0 License](https://github.com/moul/ascii2svg/blob/master/LICENSE)
