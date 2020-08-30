# Gool

![Test](https://github.com/cloudingcity/gool/workflows/Test/badge.svg)
![Lint](https://github.com/cloudingcity/gool/workflows/Lint/badge.svg)
[![codecov](https://codecov.io/gh/cloudingcity/gool/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudingcity/gool)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/gool)](https://goreportcard.com/report/github.com/cloudingcity/gool)

A toolkit make your programmer life easier.

> Inspired by [Boop](https://github.com/IvanMathy/Boop)

## Usage

```shell script
$ gool jwt-decode eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "sub": "1234567890",
    "name": "John Doe",
    "iat": 1516239022
  },
  "signature": "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

### Gool Shell

```shell script
$ gool shell 
Gool Shell
\h: show help
=# \s jwt-decode
jwt-decode=# eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "sub": "1234567890",
    "name": "John Doe",
    "iat": 1516239022
  },
  "signature": "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
jwt-decode=#
```


## Scripts

| Script                 | Description                                                                                              |
|------------------------|----------------------------------------------------------------------------------------------------------|
| timestamp-to-date, t2d | Covert unix timestamp to date                                                                            |
| date-to-timestamp, d2t | Covert date to unix timestamp                                                                            |
| jwt-decode             | Decode [JWT](https://jwt.io/)                                                                            |
| camel-case             | Coverts string to [camel case](https://en.wikipedia.org/wiki/Camel_case)                                 |
| kebab-case             | Coverts string to [kebab case](https://en.wikipedia.org/wiki/Letter_case#Special_case_styles)            |
| lower-case             | Coverts string to lower case                                                                             |
| snake-case             | Coverts string to [snake case](https://en.wikipedia.org/wiki/Snake_case)                                 |
| start-case             | Coverts string to [start case](https://en.wikipedia.org/wiki/Letter_case#Stylistic_or_specialised_usage) |
| upper-case             | Coverts string to upper case                                                                             |

## Library

```go
package main

import (
	"fmt"

	"github.com/cloudingcity/gool/pkg/cases"
	"github.com/cloudingcity/gool/pkg/date"
	"github.com/cloudingcity/gool/pkg/timestamp"
)

func main() {
	fmt.Println(cases.Snake("HelloWorld"))
	// hello_world

	fmt.Println(cases.Camel("hello world"))
	// helloWorld

	fmt.Println(date.ToTimestamp("2020-09-01"))
	// 1598918400 <nil>
	
	fmt.Println(timestamp.ToDate(1598918400))
	// 2020-09-01 08:00:00 +0800 CST
}
```
