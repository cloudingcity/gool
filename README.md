# Gool
[![Test](https://github.com/cloudingcity/gool/workflows/Test/badge.svg)](https://github.com/cloudingcity/gool/actions?query=workflow%3ATest)
[![Lint](https://github.com/cloudingcity/gool/workflows/Lint/badge.svg)](https://github.com/cloudingcity/gool/actions?query=workflow%3ALint)
[![Deploy](https://github.com/cloudingcity/gool/workflows/Deploy/badge.svg)](https://github.com/cloudingcity/gool/actions?query=workflow%3ADeploy)
[![codecov](https://codecov.io/gh/cloudingcity/gool/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudingcity/gool)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/gool)](https://goreportcard.com/report/github.com/cloudingcity/gool)

📦 A toolkit make your programmer life easier.

> Inspired by [Boop](https://github.com/IvanMathy/Boop)

![demo](https://user-images.githubusercontent.com/11569651/92252386-bd1c2c00-ef00-11ea-927b-8f6d752ba733.gif)

## Installation

### Homebrew

```shell script
brew install cloudingcity/tap/gool
```

### Docker

```shell script
docker run --rm -it ghost0436/gool
```

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
| md5                    | Computes the checksum                                                                                    |
| url-encode             | Encode url                                                                                               |
| url-decode             | Decode url                                                                                               |
| url-to-json            | Convert url to JSON                                                                                      |
| base64-encode          | Base64 encode                                                                                            |
| base64-decode          | Base64 decode                                                                                            |
| rand-string            | Generate random string of given length (characters: a-z, A-Z, 0-9)                                       |
| count                  | Get the characters length                                                                                |
| camel-case             | Coverts string to [camel case](https://en.wikipedia.org/wiki/Camel_case)                                 |
| kebab-case             | Coverts string to [kebab case](https://en.wikipedia.org/wiki/Letter_case#Special_case_styles)            |
| lower-case             | Coverts string to lower case                                                                             |
| snake-case             | Coverts string to [snake case](https://en.wikipedia.org/wiki/Snake_case)                                 |
| start-case             | Coverts string to [start case](https://en.wikipedia.org/wiki/Letter_case#Stylistic_or_specialised_usage) |
| upper-case             | Coverts string to upper case                                                                             |
| format-json            | Cleans and format JSON                                                                                   |

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

## Contributors 


<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://clouding.city"><img src="https://avatars3.githubusercontent.com/u/11569651?v=4" width="100px;" alt=""/><br /><sub><b>Clouding</b></sub></a><br /><a href="https://github.com/cloudingcity/gool/commits?author=cloudingcity" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/leozhantw"><img src="https://avatars2.githubusercontent.com/u/14068118?v=4" width="100px;" alt=""/><br /><sub><b>LeoZhan</b></sub></a><br /><a href="https://github.com/cloudingcity/gool/commits?author=leozhantw" title="Code">💻</a></td>
    <td align="center"><a href="https://letme.rocks"><img src="https://avatars1.githubusercontent.com/u/164212?v=4" width="100px;" alt=""/><br /><sub><b>Toby Lam</b></sub></a><br /><a href="https://github.com/cloudingcity/gool/commits?author=livekn" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->
