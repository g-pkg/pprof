[![Build Status](https://travis-ci.com/g-pkg/pprof.svg?branch=master)](https://travis-ci.com/g-pkg/pprof) [![codecov](https://codecov.io/gh/g-pkg/pprof/branch/master/graph/badge.svg)](https://codecov.io/gh/g-pkg/pprof)

# pprof

## Installation

`go get -u github.com/g-pkg/pprof`

## Quick Start

```go
go func(){
    log.Println(pprof.Serve(":15002"))
}()
```