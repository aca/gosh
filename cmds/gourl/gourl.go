package gourl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/aca/gosh/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:          "gourl",
	SilenceUsage: true,
}

func init() {
	Cmd.AddCommand(utils.NewCompletionCommand("gourl"))

	// func PathEscape(s string) string
	Cmd.AddCommand(cmdPathEscape)

	// func PathUnescape(s string) (string, error)
	Cmd.AddCommand(cmdPathUnescape)

	// func QueryEscape(s string) string
	Cmd.AddCommand(cmdQueryEscape)

	// func QueryUnescape(s string) (string, error)
	Cmd.AddCommand(cmdQueryUnescape)

	// func Parse(rawurl string) (*URL, error)
	Cmd.AddCommand(cmdParse)
	// func ParseRequestURI(rawurl string) (*URL, error)
	// func ParseQuery(query string) (Values, error)
	cmdPathEscape.Flags().StringP("output", "o", "", "output format")
}

var cmdPathEscape = &cobra.Command{
	Use:   "pathescape",
	Short: "func PathEscape(s string) string",
	Long: `func PathEscape(s string) string

PathEscape escapes the string so it can be safely placed inside a URL path segment,
replacing special characters (including /) with %XX sequences as needed.`,
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
		fmt.Print(url.PathEscape(s))
		return nil
	},
}

var cmdPathUnescape = &cobra.Command{
	Use:   "pathunescape",
	Short: "func PathUnescape(s string) (string, error)",
	Long: `func PathUnescape(s string) (string, error)

PathUnescape does the inverse transformation of PathEscape,
converting each 3-byte encoded substring of the form "%AB" into the
hex-decoded byte 0xAB. It returns an error if any % is not followed
by two hexadecimal digits.

PathUnescape is identical to QueryUnescape except that it does not
unescape '+' to ' ' (space).`,
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
		p, err := url.PathUnescape(s)
		if err != nil {
			return err
		}
		fmt.Print(p)
		return nil
	},
}

var cmdQueryEscape = &cobra.Command{
	Use:   "pathunescape",
	Short: "func QueryEscape(s string) string",
	Long: `func QueryEscape(s string) string

QueryEscape escapes the string so it can be safely placed
inside a URL query.`,
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
		q := url.QueryEscape(s)
		fmt.Print(q)
		return nil
	},
}

var cmdQueryUnescape = &cobra.Command{
	Use:   "queryunescape",
	Short: "func QueryUnescape(s string) (string, error)",
	Long: `func QueryUnescape(s string) (string, error)

QueryUnescape does the inverse transformation of QueryEscape,
converting each 3-byte encoded substring of the form "%AB" into the
hex-decoded byte 0xAB.
It returns an error if any % is not followed by two hexadecimal
digits.`,
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
		q, err := url.QueryUnescape(s)
		if err != nil {
			return err
		}
		fmt.Print(q)
		return nil
	},
}

var cmdParse = &cobra.Command{
	Use:   "parse",
	Short: "func Parse(rawurl string) (*URL, error)",
	Long: `func Parse(rawurl string) (*URL, error)

Parse parses rawurl into a URL structure.

The rawurl may be relative (a path, without a host) or absolute
(starting with a scheme). Trying to parse a hostname and path
without a scheme is invalid but may not necessarily return an
error, due to parsing ambiguities.`,
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
		u, err := url.Parse(s)
		if err != nil {
			return err
		}
		json.NewEncoder(os.Stdout).Encode(u)
		return nil
	},
}

var cmdParseRequestURI = &cobra.Command{
	Use:   "parserequesturi",
	Short: "func ParseRequestURI(rawurl string) (*URL, error)",
	Long: `func ParseRequestURI(rawurl string) (*URL, error)

ParseRequestURI parses rawurl into a URL structure. It assumes that
rawurl was received in an HTTP request, so the rawurl is interpreted
only as an absolute URI or an absolute path.
The string rawurl is assumed not to have a #fragment suffix.
(Web browsers strip #fragment before sending the URL to a web server.)`,
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
		u, err := url.ParseRequestURI(s)
		if err != nil {
			return err
		}
		json.NewEncoder(os.Stdout).Encode(u)
		return nil
	},
}
