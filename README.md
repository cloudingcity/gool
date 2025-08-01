# Gool
[![Test](https://github.com/cloudingcity/gool/workflows/Test/badge.svg)](https://github.com/cloudingcity/gool/actions?query=workflow%3ATest)
[![Lint](https://github.com/cloudingcity/gool/workflows/Lint/badge.svg)](https://github.com/cloudingcity/gool/actions?query=workflow%3ALint)
[![Deploy](https://github.com/cloudingcity/gool/workflows/Deploy/badge.svg)](https://github.com/cloudingcity/gool/actions?query=workflow%3ADeploy)
[![codecov](https://codecov.io/gh/cloudingcity/gool/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudingcity/gool)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/gool)](https://goreportcard.com/report/github.com/cloudingcity/gool)

📦 A toolkit make your programmer life easier.

> Inspired by [Boop](https://github.com/IvanMathy/Boop)

![Kapture 2025-07-29 at 00 32 36](https://github.com/user-attachments/assets/2d752693-cff9-4a25-a852-ad1b44e3c36f)

## Installation

### cURL

```bash
# macOS Intel
curl -L https://github.com/cloudingcity/gool/releases/latest/download/gool_Darwin_x86_64.tar.gz | \
  tar xz && sudo mv gool /usr/local/bin/

# macOS Apple Silicon  
curl -L https://github.com/cloudingcity/gool/releases/latest/download/gool_Darwin_arm64.tar.gz | \
  tar xz && sudo mv gool /usr/local/bin/

# Linux x86_64
curl -L https://github.com/cloudingcity/gool/releases/latest/download/gool_Linux_x86_64.tar.gz | \
  tar xz && sudo mv gool /usr/local/bin/

# Or for a specific version:
curl -L https://github.com/cloudingcity/gool/releases/download/v1.0.0/gool_Darwin_x86_64.tar.gz | \
  tar xz && sudo mv gool /usr/local/bin/
```

### Homebrew

```shell script
brew install --cask cloudingcity/tap/gool
```

> **Note**: On macOS, you may see a security warning since the binary is not signed/notarized (I'm too poor for Apple's $99/year developer fee 😭). To bypass this:
> 1. Go to System Preferences → Security & Privacy → General
> 2. Click "Allow Anyway" next to the blocked app warning
> 3. Or run: `sudo xattr -rd com.apple.quarantine /opt/homebrew/bin/gool`

## Usage

### Interactive Shell

```shell script
$ gool 

 ██████┐  ██████┐  ██████┐ ██┐
██┌────┘ ██┌───██┐██┌───██┐██│
██│  ███┐██│   ██│██│   ██│██│
██│   ██│██│   ██│██│   ██│██│
└██████┌┘└██████┌┘└██████┌┘███████┐
 └─────┘  └─────┘  └─────┘ └──────┘

 /help for more information

ʕ◔ϖ◔ʔ ❯ /jwt-decode
jwt-decode ❯ eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
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
jwt-decode ❯
```

### Direct CLI

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

## Commands

| Command                | Description                                                                                                        |
|------------------------|--------------------------------------------------------------------------------------------------------------------|
| tsconv                 | Smart timestamp/date converter - auto-detects input type (default: show current time as both)                     |
| timestamp-to-date, t2d | Covert unix timestamp to date (default: current time)                                                              |
| date-to-timestamp, d2t | Covert date to unix timestamp (default: current time)                                                              |
| jwt-decode             | Decode [JWT](https://jwt.io/)                                                                                      |
| md5                    | Computes the checksum                                                                                              |
| url-encode             | Encode url                                                                                                         |
| url-decode             | Decode url                                                                                                         |
| url-to-json            | Convert url to JSON                                                                                                |
| base64-encode          | Base64 encode                                                                                                      |
| base64-decode          | Base64 decode                                                                                                      |
| uuid                   | Generate UUID v4 (default: 1)                                                                                      |
| count                  | Get the characters length                                                                                          |
| camel-case             | Coverts string to [camel case](https://en.wikipedia.org/wiki/Camel_case) (fooBar)                                  |
| kebab-case             | Coverts string to [kebab case](https://en.wikipedia.org/wiki/Letter_case#Special_case_styles) (foo-bar)            |
| lower-case             | Coverts string to lower case (foo bar)                                                                             |
| snake-case             | Coverts string to [snake case](https://en.wikipedia.org/wiki/Snake_case) (foo_bar)                                 |
| start-case             | Coverts string to [start case](https://en.wikipedia.org/wiki/Letter_case#Stylistic_or_specialised_usage) (Foo Bar) |
| upper-case             | Coverts string to upper case (FOO BAR)                                                                             |
| text-escape            | Escape quotes and backslashes in text                                                                              |
| text-unescape          | Unescape quotes and backslashes in text (pretty prints JSON)                                                       |
| format-json            | Cleans and format JSON                                                                                             |

## Library

```go
package main

import (
	"fmt"

	"github.com/cloudingcity/gool/pkg/cases"
	"github.com/cloudingcity/gool/pkg/date"
	"github.com/cloudingcity/gool/pkg/text"
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

	fmt.Println(text.Escape(`{"key": "say "value""}`))
	// {\"key\": \"say \"value\"\"}

	fmt.Println(text.Unescape(`{\"key\": \"say \"value\"\"}`))
	// {"key": "say "value""}
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
