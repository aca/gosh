package gostrings

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/aca/gosh/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:          "gostrings",
	SilenceUsage: true,
}

func init() {
	Cmd.AddCommand(utils.NewCompletionCommand("gostrings"))

	// func Compare(a, b string) int
	Cmd.AddCommand(cmdCompare)

	// func Contains(s, substr string) bool
	Cmd.AddCommand(cmdContains)

	// func ContainsAny(s, chars string) bool
	Cmd.AddCommand(cmdContainsAny)

	// func ContainsRune(s string, r rune) bool

	// func Count(s, substr string) int
	Cmd.AddCommand(cmdCount)

	// func EqualFold(s, t string) bool

	// func Fields(s string) []string
	Cmd.AddCommand(cmdFields)

	// func HasPrefix(s, prefix string) bool
	Cmd.AddCommand(cmdHasPrefix)

	// func HasSuffix(s, suffix string) bool
	Cmd.AddCommand(cmdHasSuffix)

	// func Index(s, substr string) int
	Cmd.AddCommand(cmdIndex)

	// func IndexAny(s, chars string) int
	Cmd.AddCommand(cmdIndexAny)

	// func IndexRune(s string, r rune) int
	Cmd.AddCommand(cmdIndexRune)

	// func Join(elems []string, sep string) string

	// func LastIndex(s, substr string) int
	Cmd.AddCommand(cmdLastIndex)
	// func LastIndexAny(s, chars string) int
	Cmd.AddCommand(cmdLastIndexAny)

	// func Map(mapping func(rune) rune, s string) string
	// func Repeat(s string, count int) string
	Cmd.AddCommand(cmdRepeat)

	// func Replace(s, old, new string, n int) string
	Cmd.AddCommand(cmdReplace)
	// func ReplaceAll(s, old, new string) string
	Cmd.AddCommand(cmdReplaceAll)

	// func Split(s, sep string) []string
	Cmd.AddCommand(cmdSplit)

	// func SplitAfter(s, sep string) []string
	Cmd.AddCommand(cmdSplitAfter)

	// func SplitAfterN(s, sep string, n int) []string
	Cmd.AddCommand(cmdSplitAfterN)

	// func SplitN(s, sep string, n int) []string
	Cmd.AddCommand(cmdSplitN)

	// func Title(s string) string
	Cmd.AddCommand(cmdTitle)
	// func ToLower(s string) string
	Cmd.AddCommand(cmdToLower)
	// func ToTitle(s string) string
	Cmd.AddCommand(cmdToTitle)
	// func ToUpper(s string) string
	Cmd.AddCommand(cmdToUpper)

	// func ToValidUTF8(s, replacement string) string

	// func Trim(s, cutset string) string
	Cmd.AddCommand(cmdTrim)

	// func TrimLeft(s, cutset string) string
	Cmd.AddCommand(cmdTrimLeft)

	// func TrimPrefix(s, prefix string) string
	Cmd.AddCommand(cmdTrimPrefix)

	// func TrimRight(s, cutset string) string
	Cmd.AddCommand(cmdTrimRight)

	// func TrimSpace(s string) string
	Cmd.AddCommand(cmdTrimSpace)

	// func TrimSuffix(s, suffix string) string
	Cmd.AddCommand(cmdTrimSuffix)

	cmdSplitAfterN.Flags().StringP("output", "o", "", "output format")
	cmdSplitN.Flags().StringP("output", "o", "", "output format")
	cmdSplit.Flags().StringP("output", "o", "", "output format")
}

var cmdRepeat = &cobra.Command{
	Use:   "repeat",
	Short: "func Repeat(s string, count int) string",
	Long: `func Repeat(s string, count int) string

Repeat returns a new string consisting of count copies of the string s.

It panics if count is negative or if
the result of (len(s) * count) overflows.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var counts string

		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			counts = args[0]
		} else {
			s = args[0]
			counts = args[1]
		}

		count, err := strconv.Atoi(counts)
		if err != nil {
			return err
		}

		fmt.Print(strings.Repeat(s, count))
		return nil
	},
}

var cmdReplaceAll = &cobra.Command{
	Use:   "replaceall",
	Short: "func ReplaceAll(s, old, new string) string",
	Long: `func ReplaceAll(s, old, new string) string

ReplaceAll returns a copy of the string s with all
non-overlapping instances of old replaced by new.
If old is empty, it matches at the beginning of the string
and after each UTF-8 sequence, yielding up to k+1 replacements
for a k-rune string.`,
	Args:                  cobra.RangeArgs(2, 3),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var old string
		var news string

		if len(args) == 2 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			old = args[0]
			news = args[1]
		} else {
			s = args[0]
			old = args[1]
			news = args[2]
		}

		fmt.Print(strings.ReplaceAll(s, old, news))
		return nil
	},
}

var cmdReplace = &cobra.Command{
	Use:   "replace",
	Short: "func Replace(s, old, new string, n int) string",
	Long: `func Replace(s, old, new string, n int) string

Replace returns a copy of the string s with the first n
non-overlapping instances of old replaced by new.
If old is empty, it matches at the beginning of the string
and after each UTF-8 sequence, yielding up to k+1 replacements
for a k-rune string.
If n < 0, there is no limit on the number of replacements.`,
	Args:                  cobra.RangeArgs(3, 4),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var old string
		var news string
		var sn string

		if len(args) == 3 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			old = args[0]
			news = args[1]
			sn = args[2]
		} else {
			s = args[0]
			old = args[1]
			news = args[2]
			sn = args[3]
		}

		n, err := strconv.Atoi(sn)
		if err != nil {
			return err
		}

		fmt.Print(strings.Replace(s, old, news, n))
		return nil
	},
}

var cmdCompare = &cobra.Command{
	Use:   "compare",
	Short: "func Compare(a, b string) int",
	Long: `func Compare(a, b string) int

Compare returns an integer comparing two strings lexicographically.
The result will be 0 if a==b, -1 if a < b, and +1 if a > b.

Compare is included only for symmetry with package bytes.
It is usually clearer and always faster to use the built-in
string comparison operators ==, <, >, and so on.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var a string
		var b string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			a = string(stdin)
			b = args[0]
		} else {
			a = args[0]
			b = args[1]
		}
		if strings.Compare(a, b) == 0 {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
		return nil
	},
}

var cmdContains = &cobra.Command{
	Use:   "contains",
	Short: "func Contains(s, substr string) bool",
	Long: `func Contains(s, substr string) bool

Contains reports whether substr is within s.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var substr string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			substr = args[0]
		} else {
			s = args[0]
			substr = args[1]
		}
		if strings.Contains(s, substr) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
		return nil
	},
}

var cmdContainsAny = &cobra.Command{
	Use:   "containsany",
	Short: "func ContainsAny(s, chars string) bool",
	Long: `func ContainsAny(s, chars string) bool

ContainsAny reports whether any Unicode code points in chars are within s.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var chars string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			chars = args[0]
		} else {
			s = args[0]
			chars = args[1]
		}
		if strings.ContainsAny(s, chars) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
		return nil
	},
}

var cmdTitle = &cobra.Command{
	Use:                   "title",
	Short:                 "func Title(s string) string",
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
		} else {
			s = args[0]
		}
		fmt.Print(strings.Title(s))
		return nil
	},
}

var cmdToUpper = &cobra.Command{
	Use:   "toupper",
	Short: "func ToUpper(s string) string",
	Long: `func ToUpper(s string) string

ToUpper returns s with all Unicode letters mapped to their upper case.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
		} else {
			s = args[0]
		}
		fmt.Print(strings.ToUpper(s))
		return nil
	},
}

var cmdToLower = &cobra.Command{
	Use:   "tolower",
	Short: "func ToLower(s string) string",
	Long: `func ToLower(s string) string

 ToLower returns s with all Unicode letters mapped to their lower case.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
		} else {
			s = args[0]
		}
		fmt.Print(strings.ToLower(s))
		return nil
	},
}

var cmdToTitle = &cobra.Command{
	Use:   "totitle",
	Short: "func ToTitle(s string) string",
	Long: `func ToTitle(s string) string

ToTitle returns a copy of the string s with all Unicode letters mapped to their
Unicode title case.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
		} else {
			s = args[0]
		}
		fmt.Print(strings.ToTitle(s))
		return nil
	},
}

var cmdFields = &cobra.Command{
	Use:   "fields",
	Short: "func Fields(s string) []string",
	Long: `func Fields(s string) []string

Fields splits the string s around each instance of one or more consecutive
white space characters, as defined by unicode.IsSpace, returning a slice of
substrings of s or an empty slice if s contains only white space.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
		} else {
			s = args[0]
		}
		for _, v := range strings.Fields(s) {
			fmt.Println(v)
		}
		return nil
	},
}

var cmdHasPrefix = &cobra.Command{
	Use:   "hasprefix",
	Short: "func HasPrefix(s, prefix string) bool",
	Long: `func HasPrefix(s, prefix string) bool

HasPrefix tests whether the string s begins with prefix.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var prefix string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			prefix = args[0]
		} else {
			s = args[0]
			prefix = args[1]
		}
		if strings.HasPrefix(s, prefix) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
		return nil
	},
}

var cmdHasSuffix = &cobra.Command{
	Use:   "hassuffix",
	Short: "func HasSuffix(s, suffix string) bool",
	Long: `func HasSuffix(s, suffix string) bool

HasSuffix tests whether the string s ends with suffix.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var suffix string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			suffix = args[0]
		} else {
			s = args[0]
			suffix = args[1]
		}
		if strings.HasSuffix(s, suffix) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
		return nil
	},
}

var cmdIndex = &cobra.Command{
	Use:   "index",
	Short: "func Index(s, substr string) int",
	Long: `func Index(s, substr string) int

Index returns the index of the first instance of substr in s, or -1 if substr
is not present in s.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var substr string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			substr = args[0]
		} else {
			s = args[0]
			substr = args[1]
		}
		fmt.Print(strings.Index(s, substr))
		return nil
	},
}

var cmdLastIndex = &cobra.Command{
	Use:   "index",
	Short: "func LastIndex(s, substr string) int",
	Long: `func LastIndex(s, substr string) int

LastIndex returns the index of the last instance of substr in s, or -1 if
substr is not present in s.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var substr string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			substr = args[0]
		} else {
			s = args[0]
			substr = args[1]
		}
		fmt.Print(strings.LastIndex(s, substr))
		return nil
	},
}

var cmdIndexAny = &cobra.Command{
	Use:   "indexany",
	Short: "func IndexAny(s, chars string) int",
	Long: `func IndexAny(s, chars string) int

IndexAny returns the index of the first instance of any Unicode code point from
chars in s, or -1 if no Unicode code point from chars is present in s.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var chars string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			chars = args[0]
		} else {
			s = args[0]
			chars = args[1]
		}
		fmt.Print(strings.IndexAny(s, chars))
		return nil
	},
}

var cmdLastIndexAny = &cobra.Command{
	Use:   "indexany",
	Short: "func LastIndexAny(s, chars string) int",
	Long: `func LastIndexAny(s, chars string) int

LastIndexAny returns the index of the last instance of any Unicode code point
from chars in s, or -1 if no Unicode code point from chars is present in s.
  `,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var chars string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			chars = args[0]
		} else {
			s = args[0]
			chars = args[1]
		}
		fmt.Print(strings.LastIndexAny(s, chars))
		return nil
	},
}

var cmdIndexRune = &cobra.Command{
	Use:   "indexrune",
	Short: "func IndexRune(s string, r rune) int",
	Long: `func IndexRune(s string, r rune) int

IndexRune returns the index of the first instance of the Unicode code point r,
or -1 if rune is not present in s. If r is utf8.RuneError, it returns the first
instance of any invalid UTF-8 byte sequence.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var rs string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			rs = args[0]
		} else {
			s = args[0]
			rs = args[1]
		}

		r := []rune(rs)
		if len(r) != 1 {
			return errors.New("invalid rune arguments")
		}
		fmt.Print(strings.IndexRune(s, r[0]))
		return nil
	},
}

var cmdSplitN = &cobra.Command{
	Use:   "splitn",
	Short: "func SplitN(s, sep string, n int) []string",
	Long: `func SplitN(s, sep string, n int) []string

SplitN slices s into substrings separated by sep and returns a slice of
the substrings between those separators.

The count determines the number of substrings to return:
  n > 0: at most n substrings; the last substring will be the unsplit remainder.
  n == 0: the result is nil (zero substrings)
  n < 0: all substrings

Edge cases for s and sep (for example, empty strings) are handled
as described in the documentation for Split.`,
	Args:                  cobra.RangeArgs(2, 3),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var sep string
		var n int

		var sn string
		if len(args) == 2 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			sep = args[0]
			sn = args[1]
		} else {
			s = args[0]
			sep = args[1]
			sn = args[2]
		}

		n, err := strconv.Atoi(sn)
		if err != nil {
			return err
		}

		for _, v := range strings.SplitN(s, sep, n) {
			fmt.Println(v)
		}
		return nil
	},
}

var cmdSplitAfterN = &cobra.Command{
	Use:   "splitaftern",
	Short: "func SplitAfterN(s, sep string, n int) []string",
	Long: `func SplitAfterN(s, sep string, n int) []string

SplitAfterN slices s into substrings after each instance of sep and
returns a slice of those substrings.

The count determines the number of substrings to return:
  n > 0: at most n substrings; the last substring will be the unsplit remainder.
  n == 0: the result is nil (zero substrings)
  n < 0: all substrings

Edge cases for s and sep (for example, empty strings) are handled
as described in the documentation for SplitAfter.`,
	Args:                  cobra.RangeArgs(2, 3),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var sep string
		var n int

		var sn string
		if len(args) == 2 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			sep = args[0]
			sn = args[1]
		} else {
			s = args[0]
			sep = args[1]
			sn = args[2]
		}

		n, err := strconv.Atoi(sn)
		if err != nil {
			return err
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		splitted := strings.SplitAfterN(s, sep, n)

		switch output {
		case "":
			for _, v := range splitted {
				fmt.Println(v)
			}
		case "json":
			enc := json.NewEncoder(os.Stdout)
			enc.Encode(splitted)
		default:
			return utils.ErrInvalidOutputFormat
		}
		return nil
	},
}

var cmdSplit = &cobra.Command{
	Use:   "split",
	Short: "func Split(s, sep string) []string",
	Long: `func Split(s, sep string) []string

Split slices s into all substrings separated by sep and returns a slice of the
substrings between those separators.

If s does not contain sep and sep is not empty, Split returns a slice of length
1 whose only element is s.

If sep is empty, Split splits after each UTF-8 sequence. If both s and sep are
empty, Split returns an empty slice.

It is equivalent to SplitN with a count of -1.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var sep string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			sep = args[0]
		} else {
			s = args[0]
			sep = args[1]
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		splitted := strings.Split(s, sep)

		switch output {
		case "":
			for _, v := range splitted {
				fmt.Println(v)
			}
		case "json":
			enc := json.NewEncoder(os.Stdout)
			enc.Encode(splitted)
		}
		return nil
	},
}

var cmdSplitAfter = &cobra.Command{
	Use:   "splitafter",
	Short: "func SplitAfter(s, sep string) []string",
	Long: `func SplitAfter(s, sep string) []string

SplitAfter slices s into all substrings after each instance of sep and returns
a slice of those substrings.

If s does not contain sep and sep is not empty, SplitAfter returns a slice of
length 1 whose only element is s.

If sep is empty, SplitAfter splits after each UTF-8 sequence. If both s and sep
are empty, SplitAfter returns an empty slice.

It is equivalent to SplitAfterN with a count of -1.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var sep string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			sep = args[0]
		} else {
			s = args[0]
			sep = args[1]
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		splitted := strings.SplitAfter(s, sep)

		switch output {
		case "":
			for _, v := range splitted {
				fmt.Println(v)
			}
		case "json":
			enc := json.NewEncoder(os.Stdout)
			enc.Encode(splitted)
		}
		return nil
	},
}

var cmdCount = &cobra.Command{
	Use:   "count",
	Short: "func Count(s, substr string) int",
	Long: `func Count(s, substr string) int

Count counts the number of non-overlapping instances of substr in s. If substr
is an empty string, Count returns 1 + the number of Unicode code points in s.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var substr string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			substr = args[0]
		} else {
			s = args[0]
			substr = args[1]
		}
		fmt.Print(strings.Count(s, substr))
		return nil
	},
}

var cmdTrim = &cobra.Command{
	Use:   "trim",
	Short: "func Trim(s, cutset string) string",
	Long: `func Trim(s, cutset string) string

Trim returns a slice of the string s with all leading and trailing Unicode code
points contained in cutset removed.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var cutset string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			cutset = args[0]
		} else {
			s = args[0]
			cutset = args[1]
		}
		fmt.Print(strings.Trim(s, cutset))
		return nil
	},
}

var cmdTrimLeft = &cobra.Command{
	Use:   "trimleft",
	Short: "func TrimLeft(s, cutset string) string",
	Long: `func TrimLeft(s, cutset string) string

TrimLeft returns a slice of the string s with all leading Unicode code points
contained in cutset removed.

To remove a prefix, use TrimPrefix instead.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var cutset string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			cutset = args[0]
		} else {
			s = args[0]
			cutset = args[1]
		}
		fmt.Print(strings.TrimLeft(s, cutset))
		return nil
	},
}

var cmdTrimPrefix = &cobra.Command{
	Use:   "trimprefix",
	Short: "func TrimPrefix(s, prefix string) string",
	Long: `func TrimPrefix(s, prefix string) string

TrimPrefix returns s without the provided leading prefix string. If s doesn't
start with prefix, s is returned unchanged.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var prefix string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			prefix = args[0]
		} else {
			s = args[0]
			prefix = args[1]
		}
		fmt.Print(strings.TrimPrefix(s, prefix))
		return nil
	},
}

var cmdTrimRight = &cobra.Command{
	Use:   "trimright",
	Short: "func TrimRight(s, cutset string) string",
	Long: `func TrimRight(s, cutset string) string

TrimRight returns a slice of the string s, with all trailing Unicode code
points contained in cutset removed.

To remove a suffix, use TrimSuffix instead.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var cutset string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			cutset = args[0]
		} else {
			s = args[0]
			cutset = args[1]
		}
		fmt.Print(strings.TrimRight(s, cutset))
		return nil
	},
}

var cmdTrimSpace = &cobra.Command{
	Use:   "trimspace",
	Short: "func TrimSpace(s string) string",
	Long: `func TrimSpace(s string) string

TrimSpace returns a slice of the string s, with all leading and trailing white
space removed, as defined by Unicode.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
		} else {
			s = args[0]
		}
		fmt.Print(strings.TrimSpace(s))
		return nil
	},
}

var cmdTrimSuffix = &cobra.Command{
	Use:   "trimsuffix",
	Short: "func TrimSuffix(s, suffix string) string",
	Long: `func TrimSuffix(s, suffix string) string

TrimSuffix returns s without the provided trailing suffix string.
If s doesn't end with suffix, s is returned unchanged.`,
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var suffix string
		if len(args) == 1 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			s = string(stdin)
			suffix = args[0]
		} else {
			s = args[0]
			suffix = args[1]
		}
		fmt.Print(strings.TrimSuffix(s, suffix))
		return nil
	},
}
