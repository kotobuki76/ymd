<p align="center">
<h1 align="center">ymd</h1>
<p align="center"> Go library for handling YYYYMMDD format easily.</p>

</p>
<p align="center">
<a href="https://goreportcard.com/report/kotobuki76/ymd"><img src="https://goreportcard.com/badge/kotobuki76/ymd" alt="Go Report Card"></a> <a href="https://github.com/kotobuki76/ymd/releases/latest"><img src="https://img.shields.io/badge/version-0.1.0-blue.svg" alt="Release Version"></a> <a href="LICENSE"><img src="https://img.shields.io/github/license/kotobuki76/ymd.svg" alt="License"></a> 
</p>
 

# News
- v0.2.1 released and tagged on Dec 6, 2022.
- v0.2.0 released and tagged on Dec 6, 2022.
  - add validation
  - add Value Object Model
- v0.1.2 released and tagged on Nov 23, 2022.


# Features
### YYYYMMDD features
- Covert "YYYYMMDD" string to "YYYY-MM-DD" string
- Add and Sub day with "YYYYMMDD" string
- Add and Sub month with "YYYYMMDD" string
- Add and Sub year with "YYYYMMDD" string
- Calculate the duration of two "YYYYMMDD" string
- Get first "YYYYMMDD" date with the "YYYYMMDD"
- Get end "YYYYMMDD" date with the "YYYYMMDD"
- Get today's "YYYYMMDD" string
- Get "YYYYMM" string from "YYYYMMDD" string
- Explode "YYYYMMDD" string to "YYYY","MM","DD"
- Value Object Model for YYMMDD with Scan(), UnmarshalJSON() MarshalJSON()

### YYYY-MM-DD features
* Covert "YYYY-MM-DD" string to "YYYYMMDD" string

# Installation

```sh
# Go Modules
require github.com/kotobuki76/ymd v0.1.0
```


# Usage
### YYYYMMDD features

import
```go
import "github.com/kotobuki76/ymd"
```

#### Covert "YYYYMMDD" string to "YYYY-MM-DD" string
```go
got := Ymd("20221123").ToHyphen()

[output]
"2022-11-23"

```

#### Add and Sub day with "YYYYMMDD" string
```go
got := Ymd("20221123").AddDay(2)

[output]
"20221125"

```

```go
got := Ymd("20221123").AddDay(-2)

[output]
"20221121"

```

#### Add and Sub day with "YYYYMMDD" string
```go
got := Ymd("20221123").AddMonth(2)

[output]
"20231123"

```

```go
got := Ymd("20221123").AddMonth(-2)

[output]
"20220923"

```

#### Add and Sub yesr with "YYYYMMDD" string
```go
got := Ymd("20221123").AddYear(2)

[output]
"20241123"

```

```go
got := Ymd("20221123").AddYear(-2)

[output]
"20221123"

```

#### Calculate the duration of two "YYYYMMDD" string

```go
got := Ymd("20221123").Between(Ymd("20221125"))

[output]
2
```

#### Get first "YYYYMMDD" date with the "YYYYMMDD"
```go
got := Ymd("20221123").FirstDateOfMonth()

[output]
"20221101"
```

#### Get end "YYYYMMDD" date with the "YYYYMMDD"
```go
got := Ymd("20221123").EndDateOfMonth()

[output]
"20221130"
```

#### Get today's "YYYYMMDD" string
```go
got := TodayYmd()

[output]
"20221123"
```

#### Get "YYYYMM" string from "YYYYMMDD" string
```go
got := Ymd("20221123").ToYm()

[output]
"202211"
```

#### Explode "YYYYMMDD" string to "YYYY","MM","DD"
```go
got := Ymd("20221123").EXplode()

[output]
"20221","11","23"
```

### YYYY-MM-DD features
#### Covert "YYYY-MM-DD" string to "YYYYMMDD" string
```go
got := HyphenYmd("2022-11-23").ToYmd()

[output]
"20221123"
```


# Creator
[Toshinari Yamashita](https://github.com/kotobuki76) (kotobuki76@gmail.com)

# License
**ymd** released under MIT license, refer [LICENSE](LICENSE) file.