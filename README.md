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

| Script                 | Description                   |
|------------------------|-------------------------------|
| timestamp-to-date, t2d | Covert unix timestamp to date |
| date-to-timestamp, d2t | Covert date to unix timestamp |
| jwt-decode             | Decode [JWT](https://jwt.io/) |


