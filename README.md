# gosh
Shell utilies for gophers.  

### Why?

- [How to trim whitespace from a Bash variable?](https://stackoverflow.com/questions/369758/how-to-trim-whitespace-from-a-bash-variable)
- [Extract filename and extension in Bash](https://stackoverflow.com/questions/965053/extract-filename-and-extension-in-bash)


Cryptic syntax of bash, sed, awk and subtle differences between sed/gsed, awk/gawk... I always had a hard time memorising all these stuff. 
I wanted simple / universal way to manipulate strings in command line like [Fish's string builtin](http://fishshell.com/docs/current/cmds/string.html).
But then I thought it would be much better if we can have some kind of standard shell libraries not just for strings, but for others also, based on [Go](https://golang.org/).


### Examples
- Most commands also reads input either from stdin or as arguments. 

  [func Ext(path string) string](https://golang.org/pkg/path/filepath/#Ext)
  ```sh
  $ gofilepath ext "gonet/main.go"
  .go

  $ echo -n "gonet/main.go" |  gostrings ext
  .go
  ```

- Functions which returns boolean or error, exit with error code 1 if false or err is not neal.

  [func HasPrefix(s, prefix string) bool](https://golang.org/pkg/strings/#HasPrefix)
  ```sh
  $ gostrings hasprefix "foo" "so"; echo $?
  1
  ```

- For functions that returns slice or multiple value, result is printed in order seperated by line.

  [func Split(s, sep string) []string](https://golang.org/pkg/strings/#Split)
  ```sh
  $ echo -n 'a,b' | gostrings split ','
  a
  b
  ```

  But also you can specify output format in json.

  [func Split(path string) (dir, file string)](https://golang.org/pkg/filepath/#Split)
  ```sh
  $ gofilepath split -o json -- "a/b.go" | jq .
  {
    "dir": "a/",
    "file": "b.go"
  }
  ```

  [func LookupHost(host string) (addrs []string, err error)](https://golang.org/pkg/net/#LookupHost)
  ```
  $ gonet lookuphost "google.com" -o json
  ["172.217.175.238","2404:6800:4004:811::200e"]
  ```


### Install
- Seperate packages
  ```
  go get -u github.com/aca/gosh/gostrings
  go get -u github.com/aca/gosh/gonet
  go get -u github.com/aca/gosh/gofilepath
  go get -u github.com/aca/gosh/gourl
  ```

- Unified packages
  ```sh
  go get -u github.com/aca/gosh

  # Set alias
  # alias gostrings="gosh gostrings"
  # alias gonet="gosh gonet"
  # alias gofilepath="gosh gofilepath"
  # alias gourl="gosh gourl"
  ```

- Completion
  ```sh
  # check installation guide for each command
  gostrings completion --help
  ```

### Usage
**gostrings**
```
$ gostrings --help
Usage:
  gostrings [command]

Available Commands:
  compare     func Compare(a, b string) int
  completion  Generate completion script
  contains    func Contains(s, substr string) bool
  containsany func ContainsAny(s, chars string) bool
  count       func Count(s, substr string) int
  fields      func Fields(s string) []string
  hasprefix   func HasPrefix(s, prefix string) bool
  hassuffix   func HasSuffix(s, suffix string) bool
  help        Help about any command
  index       func Index(s, substr string) int
  index       func LastIndex(s, substr string) int
  indexany    func LastIndexAny(s, chars string) int
  indexany    func IndexAny(s, chars string) int
  indexrune   func IndexRune(s string, r rune) int
  repeat      func Repeat(s string, count int) string
  replace     func Replace(s, old, new string, n int) string
  replaceall  func ReplaceAll(s, old, new string) string
  split       func Split(s, sep string) []string
  splitafter  func SplitAfter(s, sep string) []string
  splitaftern func SplitAfterN(s, sep string, n int) []string
  splitn      func SplitN(s, sep string, n int) []string
  title       func Title(s string) string
  tolower     func ToLower(s string) string
  totitle     func ToTitle(s string) string
  toupper     func ToUpper(s string) string
  trim        func Trim(s, cutset string) string
  trimleft    func TrimLeft(s, cutset string) string
  trimprefix  func TrimPrefix(s, prefix string) string
  trimright   func TrimRight(s, cutset string) string
  trimspace   func TrimSpace(s string) string
  trimsuffix  func TrimSuffix(s, suffix string) string

Flags:
  -h, --help   help for gostrings

Use "gostrings [command] --help" for more information about a command.
```

**gofilepath**
```
$ gofilepath --help
Usage:
  gofilepath [command]

Available Commands:
  abs          func Abs(path string) (string, error)
  base         func Base(path string) string
  clean        func Clean(path string) string
  completion   Generate completion script
  dir          func Dir(path string) string
  evalsymlinks func EvalSymlinks(path string) (string, error)
  ext          func Ext(path string) string
  glob         func Glob(pattern string) (matches []string, err error)
  help         Help about any command
  isabs        func IsAbs(path string) bool
  rel          func Rel(basepath, targpath string) (string, error)
  split        func Split(path string) (dir, file string)
  splitlist    func SplitList(path string) []string

Flags:
  -h, --help   help for gofilepath

Use "gofilepath [command] --help" for more information about a command.
```

**gonet**
```
$ gonet --help
Usage:
  gonet [command]

Available Commands:
  completion   Generate completion script
  help         Help about any command
  joinhostport func JoinHostPort(host, port string) string
  lookupaddr   func LookupAddr(addr string) (names []string, err error)
  lookupcname  func LookupCNAME(host string) (cname string, err error)
  lookuphost   func LookupHost(host string) (addrs []string, err error)
  lookuptxt    func LookupTXT(name string) ([]string, error)
  parsecidr    func ParseCIDR(s string) (IP, *IPNet, error)

Flags:
  -h, --help   help for gonet

Use "gonet [command] --help" for more information about a command.
```

**gourl**
```
$ gourl --help
Usage:
  gourl [command]

Available Commands:
  completion    Generate completion script
  help          Help about any command
  parse         func Parse(rawurl string) (*URL, error)
  pathescape    func PathEscape(s string) string
  pathunescape  func PathUnescape(s string) (string, error)
  pathunescape  func QueryEscape(s string) string
  queryunescape func QueryUnescape(s string) (string, error)

Flags:
  -h, --help   help for gourl

Use "gourl [command] --help" for more information about a command.
```
