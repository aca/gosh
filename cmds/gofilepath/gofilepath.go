package gofilepath

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aca/gosh/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:          "gofilepath",
	SilenceUsage: true,
}

func init() {
	Cmd.AddCommand(utils.NewCompletionCommand("gofilepath"))

	// func Compare(a, b string) int

	// func Abs(path string) (string, error)
	Cmd.AddCommand(cmdAbs)

	// func Base(path string) string
	Cmd.AddCommand(cmdBase)

	// func Clean(path string) string
	Cmd.AddCommand(cmdClean)

	// func Dir(path string) string
	Cmd.AddCommand(cmdDir)

	// func EvalSymlinks(path string) (string, error)
	Cmd.AddCommand(cmdEvalSymlinks)

	// func Ext(path string) string
	Cmd.AddCommand(cmdExt)

	// func FromSlash(path string) string
	// func Glob(pattern string) (matches []string, err error)
	Cmd.AddCommand(cmdGlob)

	// func HasPrefix(p, prefix string) bool

	// func IsAbs(path string) bool
	Cmd.AddCommand(cmdIsAbs)

	// func Join(elem ...string) string
	// func Match(pattern, name string) (matched bool, err error)
	// func Rel(basepath, targpath string) (string, error)
	Cmd.AddCommand(cmdRel)

	// func Split(path string) (dir, file string)
	Cmd.AddCommand(cmdSplit)

	// func SplitList(path string) []string
	Cmd.AddCommand(cmdSplitList)

	// func ToSlash(path string) string
	// func VolumeName(path string) string

	cmdSplit.Flags().StringP("output", "o", "", "output format")
	cmdGlob.Flags().StringP("output", "o", "", "output format")
	cmdSplitList.Flags().StringP("output", "o", "", "output format")
}

var cmdSplitList = &cobra.Command{
	Use:   "splitlist",
	Short: "func SplitList(path string) []string",
	Long: `func SplitList(path string) []string

Split splits path immediately following the final Separator, separating it into
a directory and file name component. If there is no Separator in path, Split
returns an empty dir and file set to path. The returned values have the
property that path = dir+file.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		l := filepath.SplitList(path)

		switch output {
		case "":
			for _, match := range l {
				fmt.Println(match)
			}
		case "json":
			enc := json.NewEncoder(os.Stdout)
			err := enc.Encode(&l)
			if err != nil {
				return err
			}
		default:
			return utils.ErrInvalidOutputFormat
		}

		return nil
	},
}

var cmdEvalSymlinks = &cobra.Command{
	Use:   "evalsymlinks",
	Short: "func EvalSymlinks(path string) (string, error)",
	Long: `func EvalSymlinks(path string) (string, error)

EvalSymlinks returns the path name after the evaluation of any symbolic links.
If path is relative the result will be relative to the current directory,
unless one of the components is an absolute symbolic link. EvalSymlinks calls
Clean on the result.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}

		s, err := filepath.EvalSymlinks(path)
		if err != nil {
			return err
		}

		fmt.Print(s)

		return nil
	},
}

var cmdGlob = &cobra.Command{
	Use:   "glob",
	Short: "func Glob(pattern string) (matches []string, err error)",
	Long: `func Glob(pattern string) (matches []string, err error)

Glob returns the names of all files matching pattern or nil
if there is no matching file. The syntax of patterns is the same
as in Match. The pattern may describe hierarchical names such as
/usr/*/bin/ed (assuming the Separator is '/').

Glob ignores file system errors such as I/O errors reading directories.
The only possible returned error is ErrBadPattern, when pattern
is malformed.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		matches, err := filepath.Glob(path)
		if err != nil {
			return err
		}

		switch output {
		case "":
			for _, match := range matches {
				fmt.Println(match)
			}
		case "json":
			enc := json.NewEncoder(os.Stdout)
			err := enc.Encode(&matches)
			if err != nil {
				return err
			}
		default:
			return utils.ErrInvalidOutputFormat
		}

		return nil
	},
}

var cmdRel = &cobra.Command{
	Use:   "rel",
	Short: "func Rel(basepath, targpath string) (string, error)",
	Long: `func Rel(basepath, targpath string) (string, error)

Rel returns a relative path that is lexically equivalent to targpath when
joined to basepath with an intervening separator. That is,
Join(basepath, Rel(basepath, targpath)) is equivalent to targpath itself.
On success, the returned path will always be relative to basepath,
even if basepath and targpath share no elements.
An error is returned if targpath can't be made relative to basepath or if
knowing the current working directory would be necessary to compute it.
Rel calls Clean on the result.`,

	Args:                  cobra.ExactArgs(2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		basepath := args[0]
		targpath := args[1]

		relpath, err := filepath.Rel(basepath, targpath)
		if err != nil {
			return err
		}

		fmt.Print(relpath)
		return nil
	},
}

var cmdAbs = &cobra.Command{
	Use:   "abs",
	Short: "func Abs(path string) (string, error)",
	Long: `func Abs(path string) (string, error)

Abs returns an absolute representation of path.
If the path is not absolute it will be joined with the current
working directory to turn it into an absolute path. The absolute
path name for a given file is not guaranteed to be unique.
Abs calls Clean on the result.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}
		abs, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		fmt.Print(abs)
		return nil
	},
}

var cmdSplit = &cobra.Command{
	Use:   "split",
	Short: "func Split(path string) (dir, file string)",
	Long: `func Split(path string) (dir, file string)

Split splits path immediately following the final Separator,
separating it into a directory and file name component.
If there is no Separator in path, Split returns an empty dir
and file set to path.
The returned values have the property that path = dir+file.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		dir, file := filepath.Split(path)

		switch output {
		case "":
			fmt.Println(dir)
			fmt.Println(file)

		case "json":
			st := &struct {
				Dir  string `json:"dir"`
				File string `json:"file"`
			}{
				Dir:  dir,
				File: file,
			}

			enc := json.NewEncoder(os.Stdout)
			err := enc.Encode(st)
			if err != nil {
				return err
			}
		default:
			return utils.ErrInvalidOutputFormat
		}

		return nil
	},
}

var cmdClean = &cobra.Command{
	Use:   "clean",
	Short: "func Clean(path string) string",
	Long: `func Clean(path string) string

Clean returns the shortest path name equivalent to path
by purely lexical processing. It applies the following rules
iteratively until no further processing can be done:

 1. Replace multiple Separator elements with a single one.
 2. Eliminate each . path name element (the current directory).
 3. Eliminate each inner .. path name element (the parent directory)
    along with the non-.. element that precedes it.
 4. Eliminate .. elements that begin a rooted path:
    that is, replace "/.." by "/" at the beginning of a path,
    assuming Separator is '/'.

The returned path ends in a slash only if it represents a root directory,
such as "/" on Unix or "C:\" on Windows.

Finally, any occurrences of slash are replaced by Separator.

If the result of this process is an empty string, Clean
returns the string ".".

See also Rob Pike, ""Lexical File Names in Plan 9 or
Getting Dot-Dot Right,""
https://9p.io/sys/doc/lexnames.html`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}
		fmt.Print(filepath.Clean(path))
		return nil
	},
}

var cmdDir = &cobra.Command{
	Use:   "dir",
	Short: "func Dir(path string) string",
	Long: `func Dir(path string) string

Dir returns all but the last element of path, typically the path's directory.
After dropping the final element, Dir calls Clean on the path and trailing
slashes are removed.
If the path is empty, Dir returns ".".
If the path consists entirely of separators, Dir returns a single separator.
The returned path does not end in a separator unless it is the root directory.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}
		fmt.Print(filepath.Dir(path))
		return nil
	},
}

var cmdBase = &cobra.Command{
	Use:   "base",
	Short: "func Base(path string) string",
	Long: `func Base(path string) string

Base returns the last element of path.
Trailing path separators are removed before extracting the last element.
If the path is empty, Base returns ".".
If the path consists entirely of separators, Base returns a single separator.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}
		fmt.Print(filepath.Base(path))
		return nil
	},
}

var cmdIsAbs = &cobra.Command{
	Use:   "isabs",
	Short: "func IsAbs(path string) bool",
	Long: `func IsAbs(path string) bool

Abs returns an absolute representation of path.
If the path is not absolute it will be joined with the current
working directory to turn it into an absolute path. The absolute
path name for a given file is not guaranteed to be unique.
Abs calls Clean on the result.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}
		if filepath.IsAbs(path) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
		return nil
	},
}

var cmdExt = &cobra.Command{
	Use:   "ext",
	Short: "func Ext(path string) string",
	Long: `func Ext(path string) string

Ext returns the file name extension used by path.
The extension is the suffix beginning at the final dot
in the final element of path; it is empty if there is
no dot.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			path = string(stdin)
		} else {
			path = args[0]
		}
		fmt.Print(filepath.Ext(path))
		return nil
	},
}
