# gosh
Shell utilies for gophers.  

Why?
----
Simple / universal way to manipulate string in command line like [Fish's string](http://fishshell.com/docs/current/cmds/string.html).  
Based on Go standard library.

Example
---
[func Fields(s string) []string](https://golang.org/pkg/strings/#Fields)
```sh
$ gostrings fields "foo bar"
foo
bar

$ echo -n "foo bar" |  gostrings fields
foo
bar
```
---
[func HasPrefix(s, prefix string) bool](https://golang.org/pkg/strings/#HasPrefix)
```sh
$ gostrings hasprefix "foo" "so"
$ echo $?
1

$ echo "foo" | gostrings hasprefix "fo"
$ echo $?
0
```

Install
-------

Each packages
```
go get -u github.com/aca/gosh/gostrings
```

Unified packages
```
go get -u github.com/aca/gosh

alias gostrings="gosh strings"
```

Completion
----------
```
TODO
```

Usage
-----
**gostrings**
```
» gostrings --help
Usage:
  strings [command]

Available Commands:
  compare     func Compare(a, b string) int
  contains    func Contains(s, substr string) bool
  containsany func ContainsAny(s, chars string) bool
  count       func Count(s, substr string) int
  fields      func Fields(s string) []string
  hasprefix   func HasPrefix(s, prefix string) bool
  hassuffix   func HasSuffix(s, suffix string) bool
  help        Help about any command
  index       func LastIndex(s, substr string) int
  index       func Index(s, substr string) int
  indexany    func IndexAny(s, chars string) int
  indexany    func LastIndexAny(s, chars string) int
  indexrune   func IndexRune(s string, r rune) int
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
  -h, --help   help for strings

Use "strings [command] --help" for more information about a command.
```
