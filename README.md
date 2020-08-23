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

### Interactive Mode

```shell script
$ gool i
Gool Shell
\h: show help
=# \c jwt-decode
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

| Script     | Description                   |
|------------|-------------------------------|
| jwt-decode | Decode [JWT](https://jwt.io/) |



