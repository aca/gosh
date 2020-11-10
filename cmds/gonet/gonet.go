package gonet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/aca/gosh/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:          "gonet",
	SilenceUsage: true,
}

func init() {
	Cmd.AddCommand(utils.NewCompletionCommand("gonet"))

	// func JoinHostPort(host, port string) string
	Cmd.AddCommand(cmdJoinHostPort)

	// func LookupAddr(addr string) (names []string, err error)
	Cmd.AddCommand(cmdLookupAddr)

	// func LookupCNAME(host string) (cname string, err error)
	Cmd.AddCommand(cmdLookupCNAME)

	// func LookupHost(host string) (addrs []string, err error)
	Cmd.AddCommand(cmdLookupHost)

	// func LookupTXT(name string) ([]string, error)
	Cmd.AddCommand(cmdLookupTXT)

	// func ParseCIDR(s string) (IP, *IPNet, error)
	Cmd.AddCommand(cmdParseCIDR)

	cmdLookupAddr.Flags().StringP("output", "o", "", "output format")
	cmdLookupHost.Flags().StringP("output", "o", "", "output format")
	cmdLookupTXT.Flags().StringP("output", "o", "", "output format")
}

var cmdParseCIDR = &cobra.Command{
	Use:   "parsecidr",
	Short: "func ParseCIDR(s string) (IP, *IPNet, error)",
	Long: `func ParseCIDR(s string) (IP, *IPNet, error)

ParseCIDR parses s as a CIDR notation IP address and prefix length,
like "192.0.2.0/24" or "2001:db8::/32", as defined in
RFC 4632 and RFC 4291.

It returns the IP address and the network implied by the IP and
prefix length.
For example, ParseCIDR("192.0.2.1/24") returns the IP address
192.0.2.1 and the network 192.0.2.0/24.`,
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

		ip, ipnet, err := net.ParseCIDR(s)
		if err != nil {
			return err
		}

		fmt.Println(ip)
		fmt.Println(ipnet)
		return nil
	},
}

var cmdJoinHostPort = &cobra.Command{
	Use:   "joinhostport",
	Short: "func JoinHostPort(host, port string) string",
	Long: `func JoinHostPort(host, port string) string

JoinHostPort combines host and port into a network address of the
form "host:port". If host contains a colon, as found in literal
IPv6 addresses, then JoinHostPort returns "[host]:port".

See func Dial for a description of the host and port parameters.`,
	Args:                  cobra.ExactArgs(2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		host := args[0]
		port := args[1]

		fmt.Print(net.JoinHostPort(host, port))
		return nil
	},
}

var cmdLookupHost = &cobra.Command{
	Use:   "lookuphost",
	Short: "func LookupHost(host string) (addrs []string, err error)",
	Long: `func LookupHost(host string) (addrs []string, err error)

LookupHost looks up the given host using the local resolver.
It returns a slice of that host's addresses.`,
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

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		names, err := net.LookupHost(s)
		if err != nil {
			return err
		}

		switch output {
		case "":
			for _, v := range names {
				fmt.Println(v)
			}
		case "json":
			enc := json.NewEncoder(os.Stdout)
			enc.Encode(names)
		default:
			return utils.ErrInvalidOutputFormat
		}
		return nil
	},
}

var cmdLookupAddr = &cobra.Command{
	Use:   "lookupaddr",
	Short: "func LookupAddr(addr string) (names []string, err error)",
	Long: `func LookupAddr(addr string) (names []string, err error)

LookupAddr performs a reverse lookup for the given address, returning a list
of names mapping to that address.

When using the host C library resolver, at most one result will be
returned. To bypass the host resolver, use a custom Resolver.`,
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

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		names, err := net.LookupAddr(s)
		if err != nil {
			return err
		}

		switch output {
		case "":
			for _, v := range names {
				fmt.Println(v)
			}
		case "json":
			enc := json.NewEncoder(os.Stdout)
			enc.Encode(names)
		default:
			return utils.ErrInvalidOutputFormat
		}
		return nil
	},
}

var cmdLookupCNAME = &cobra.Command{
	Use:   "lookupcname",
	Short: "func LookupCNAME(host string) (cname string, err error)",
	Long: `func LookupCNAME(host string) (cname string, err error)

LookupCNAME returns the canonical name for the given host.
Callers that do not care about the canonical name can call
LookupHost or LookupIP directly; both take care of resolving
the canonical name as part of the lookup.

A canonical name is the final name after following zero
or more CNAME records.
LookupCNAME does not return an error if host does not
contain DNS "CNAME" records, as long as host resolves to
address records.`,
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var host string
		if len(args) == 0 {
			stdin, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			host = string(stdin)
		} else {
			host = args[0]
		}

		cname, err := net.LookupCNAME(host)
		if err != nil {
			return err
		}

		fmt.Print(cname)
		return nil
	},
}

var cmdLookupTXT = &cobra.Command{
	Use:   "lookuptxt",
	Short: "func LookupTXT(name string) ([]string, error)",
	Long: `func LookupTXT(name string) ([]string, error)

LookupTXT returns the DNS TXT records for the given domain name.`,
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

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		names, err := net.LookupTXT(s)
		if err != nil {
			return err
		}

		switch output {
		case "":
			for _, v := range names {
				fmt.Println(v)
			}
		case "json":
			enc := json.NewEncoder(os.Stdout)
			enc.Encode(names)
		default:
			return utils.ErrInvalidOutputFormat
		}
		return nil
	},
}
