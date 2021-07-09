/*
Copyright © 2021 Fregie <xiaohao950830@live.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/fregie/simple/proto/gen/go/api"
	inf "github.com/fregie/simple/proto/gen/go/simple-interface"
)

// getSessCmd represents the session command
var getSessCmd = &cobra.Command{
	Use:   "session",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showSession,
}

var createSessCmd = &cobra.Command{
	Use:   "session",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: createSession,
}

var delSessCmd = &cobra.Command{
	Use:   "session",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: delSession,
}

func showSession(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		pterm.Error.Printf("Need provide session ID")
		return
	}
	id := args[0]
	rsp, err := srv.GetSession(cmd.Context(), &api.GetSessionReq{ID: id})
	checkErr(err)
	checkRsp(rsp.Code, rsp.Msg)
	sess := rsp.Session
	pterm.FgLightCyan.Printf("ID:            %s\n", sess.ID)
	pterm.Printf("Proto:         %s\n", sess.Proto)
	pterm.Printf("Config type:   %s\n", sess.ConfigType)
	pterm.Print("Option:\n")
	pterm.Printf("    Upload rate limit:   %d mbps\n", sess.Opt.SendRateLimit)
	pterm.Printf("    Download rate limit: %d mbps\n", sess.Opt.RecvRateLimit)
}

func createSession(cmd *cobra.Command, args []string) {
	var proto, ctype *string
	var limit *uint64
	proto = cmd.Flags().String("proto", "trojan", "proto")
	ctype = cmd.Flags().String("type", "json", "config type")
	limit = cmd.Flags().Uint64("limit", 0, "speed limit")
	rsp, err := srv.CreateSession(cmd.Context(), &api.CreateSessionReq{
		Proto:      *proto,
		ConfigType: parseConfigType(*ctype),
		Opt: &inf.Option{
			SendRateLimit: *limit,
			RecvRateLimit: *limit,
		},
	})
	checkErr(err)
	checkRsp(rsp.Code, rsp.Msg)
	pterm.Success.Print("Create success!\n")
	pterm.FgLightCyan.Printf("ID:            %s\n", rsp.ID)
	pterm.Printf("Proto:         %s\n", rsp.Proto)
	pterm.Printf("Config type:   %s\n", rsp.ConfigType.String())
	pterm.Printf("Config:\n%s\n", rsp.Config)
}

func delSession(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		pterm.Error.Printf("Need provide session ID")
		return
	}
	id := args[0]
	rsp, err := srv.DeleteSession(cmd.Context(), &api.DeleteSessionReq{ID: id})
	checkErr(err)
	checkRsp(rsp.Code, rsp.Msg)
	pterm.Success.Printf("Delete %s\n", id)
}

func parseConfigType(t string) inf.ConfigType {
	switch t {
	case "json", "JSON":
		return inf.ConfigType_JSON
	case "url", "URL":
		return inf.ConfigType_URL
	case "yaml", "yml", "YAML", "YML":
		return inf.ConfigType_YAML
	default:
		return inf.ConfigType_TEXT
	}
}

func init() {
	getCmd.AddCommand(getSessCmd)
	createCmd.AddCommand(createSessCmd)
	deleteCmd.AddCommand(delSessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sessionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sessionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
