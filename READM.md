**ymd** is a Go library for handling YYYYMMDD format easily.


# News
* v0.1.0 released and tagged on Nov 23, 2022.

# Features
### YYYYMMDD features
* Covert "YYYYMMDD" string to "YYYY-MM-DD" string
* Add and Sub day with "YYYYMMDD" string
* Add and Sub month with "YYYYMMDD" string
* Add and Sub year with "YYYYMMDD" string
* Calculate the duration of two "YYYYMMDD" string
* Get first "YYYYMMDD" date with the "YYYYMMDD"
* Get end "YYYYMMDD" date with the "YYYYMMDD"
* Get today's "YYYYMMDD" string
* Get "YYYYMM" string from "YYYYMMDD" string
* Explode "YYYYMMDD" string to "YYYY","MM","DD"

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