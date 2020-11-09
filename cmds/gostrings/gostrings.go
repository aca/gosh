package gostrings

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var CmdGostrings = &cobra.Command{
	Use:          "strings",
	SilenceUsage: true,
}

func init() {
	// func Compare(a, b string) int
	CmdGostrings.AddCommand(cmdCompare)

	// func Contains(s, substr string) bool
	CmdGostrings.AddCommand(cmdContains)

	// func ContainsAny(s, chars string) bool
	CmdGostrings.AddCommand(cmdContainsAny)

	// func ContainsRune(s string, r rune) bool

	// func Count(s, substr string) int
	CmdGostrings.AddCommand(cmdCount)

	// func EqualFold(s, t string) bool

	// func Fields(s string) []string
	CmdGostrings.AddCommand(cmdFields)

	// func HasPrefix(s, prefix string) bool
	CmdGostrings.AddCommand(cmdHasPrefix)

	// func HasSuffix(s, suffix string) bool
	CmdGostrings.AddCommand(cmdHasSuffix)

	// func Index(s, substr string) int
	CmdGostrings.AddCommand(cmdIndex)

	// func IndexAny(s, chars string) int
	CmdGostrings.AddCommand(cmdIndexAny)

	// func IndexRune(s string, r rune) int
	CmdGostrings.AddCommand(cmdIndexRune)

	// func Join(elems []string, sep string) string

	// func LastIndex(s, substr string) int
	CmdGostrings.AddCommand(cmdLastIndex)
	// func LastIndexAny(s, chars string) int
	CmdGostrings.AddCommand(cmdLastIndexAny)

	// func LastIndexByte(s string, c byte) int
	// func Map(mapping func(rune) rune, s string) string
	// func Repeat(s string, count int) string
	// func Replace(s, old, new string, n int) string
	// func ReplaceAll(s, old, new string) string

	// func Split(s, sep string) []string
	CmdGostrings.AddCommand(cmdSplit)

	// func SplitAfter(s, sep string) []string
	CmdGostrings.AddCommand(cmdSplitAfter)

	// func SplitAfterN(s, sep string, n int) []string
	CmdGostrings.AddCommand(cmdSplitAfterN)

	// func SplitN(s, sep string, n int) []string
	CmdGostrings.AddCommand(cmdSplitN)

	// func Title(s string) string
	CmdGostrings.AddCommand(cmdTitle)
	// func ToLower(s string) string
	CmdGostrings.AddCommand(cmdToLower)
	// func ToTitle(s string) string
	CmdGostrings.AddCommand(cmdToTitle)
	// func ToUpper(s string) string
	CmdGostrings.AddCommand(cmdToUpper)

	// func ToValidUTF8(s, replacement string) string

	// func Trim(s, cutset string) string
	CmdGostrings.AddCommand(cmdTrim)

	// func TrimLeft(s, cutset string) string
	CmdGostrings.AddCommand(cmdTrimLeft)

	// func TrimPrefix(s, prefix string) string
	CmdGostrings.AddCommand(cmdTrimPrefix)

	// func TrimRight(s, cutset string) string
	CmdGostrings.AddCommand(cmdTrimRight)

	// func TrimSpace(s string) string
	CmdGostrings.AddCommand(cmdTrimSpace)

	// func TrimSuffix(s, suffix string) string
	CmdGostrings.AddCommand(cmdTrimSuffix)
}

var cmdCompare = &cobra.Command{
	Use:                   "compare",
	Short:                 "func Compare(a, b string) int",
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
	Use:                   "contains",
	Short:                 "func Contains(s, substr string) bool",
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
	Use:                   "containsany",
	Short:                 "func ContainsAny(s, chars string) bool",
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
		fmt.Printf(strings.Title(s))
		return nil
	},
}

var cmdToUpper = &cobra.Command{
	Use:                   "toupper",
	Short:                 "func ToUpper(s string) string",
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
		fmt.Printf(strings.ToUpper(s))
		return nil
	},
}

var cmdToLower = &cobra.Command{
	Use:                   "tolower",
	Short:                 "func ToLower(s string) string",
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
		fmt.Printf(strings.ToLower(s))
		return nil
	},
}

var cmdToTitle = &cobra.Command{
	Use:                   "totitle",
	Short:                 "func ToTitle(s string) string",
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
		fmt.Printf(strings.ToTitle(s))
		return nil
	},
}

var cmdFields = &cobra.Command{
	Use:                   "fields",
	Short:                 "func Fields(s string) []string",
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
	Use:                   "hasprefix",
	Short:                 "func HasPrefix(s, prefix string) bool",
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
	Use:                   "hassuffix",
	Short:                 "func HasSuffix(s, suffix string) bool",
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
	Use:                   "index",
	Short:                 "func Index(s, substr string) int",
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
		fmt.Println(strings.Index(s, substr))
		return nil
	},
}

var cmdLastIndex = &cobra.Command{
	Use:                   "index",
	Short:                 "func LastIndex(s, substr string) int",
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
		fmt.Println(strings.LastIndex(s, substr))
		return nil
	},
}

var cmdIndexAny = &cobra.Command{
	Use:                   "indexany",
	Short:                 "func IndexAny(s, chars string) int",
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
		fmt.Println(strings.IndexAny(s, chars))
		return nil
	},
}

var cmdLastIndexAny = &cobra.Command{
	Use:                   "indexany",
	Short:                 "func LastIndexAny(s, chars string) int",
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
		fmt.Println(strings.LastIndexAny(s, chars))
		return nil
	},
}

var cmdIndexRune = &cobra.Command{
	Use:                   "indexrune",
	Short:                 "func IndexRune(s string, r rune) int",
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
		fmt.Println(strings.IndexRune(s, r[0]))
		return nil
	},
}

var cmdSplitN = &cobra.Command{
	Use:                   "splitn",
	Short:                 "func SplitN(s, sep string, n int) []string",
	Args:                  cobra.RangeArgs(2, 3),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var sep string
		var n int

		var sn string
		if len(args) == 0 {
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
	Use:                   "splitaftern",
	Short:                 "func SplitAfterN(s, sep string, n int) []string",
	Args:                  cobra.RangeArgs(2, 3),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var sep string
		var n int

		var sn string
		if len(args) == 0 {
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

		for _, v := range strings.SplitAfterN(s, sep, n) {
			fmt.Println(v)
		}
		return nil
	},
}

var cmdSplit = &cobra.Command{
	Use:                   "split",
	Short:                 "func Split(s, sep string) []string",
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var sep string
		if len(args) == 0 {
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
		for _, v := range strings.Split(s, sep) {
			fmt.Println(v)
		}
		return nil
	},
}

var cmdSplitAfter = &cobra.Command{
	Use:                   "splitafter",
	Short:                 "func SplitAfter(s, sep string) []string",
	Args:                  cobra.RangeArgs(1, 2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		var sep string
		if len(args) == 0 {
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
		for _, v := range strings.SplitAfter(s, sep) {
			fmt.Println(v)
		}
		return nil
	},
}

var cmdCount = &cobra.Command{
	Use:                   "count",
	Short:                 "func Count(s, substr string) int",
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
	Use:                   "trim",
	Short:                 "func Trim(s, cutset string) string",
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
	Use:                   "trimleft",
	Short:                 "func TrimLeft(s, cutset string) string",
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
	Use:                   "trimprefix",
	Short:                 "func TrimPrefix(s, prefix string) string",
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
	Use:                   "trimright",
	Short:                 "func TrimRight(s, cutset string) string",
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
	Use:                   "trimspace",
	Short:                 "func TrimSpace(s string) string",
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
	Use:                   "trimsuffix",
	Short:                 "func TrimSuffix(s, suffix string) string",
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
