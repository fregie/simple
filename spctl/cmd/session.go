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

var (
	isShowConfig *bool
	proto, ctype *string
	limit        *uint64
	createName   *string
)

// getSessCmd represents the session command
var getSessCmd = &cobra.Command{
	Use:   "session",
	Short: "Get specified session",
	Long:  `Get specified session`,
	Args:  cobra.MinimumNArgs(1),
	Run:   showSession,
}

var createSessCmd = &cobra.Command{
	Use:   "session",
	Short: "Create a new sessoin",
	Long:  `Create a new sessoin`,
	Run:   createSession,
}

var delSessCmd = &cobra.Command{
	Use:   "session",
	Short: "Delete specified session",
	Long:  `Delete specified session`,
	Args:  cobra.MinimumNArgs(1),
	Run:   delSession,
}

func showSession(cmd *cobra.Command, args []string) {
	id := args[0]

	rsp, err := srv.GetSession(cmd.Context(), &api.GetSessionReq{IDorName: id})
	checkErr(err)
	checkRsp(rsp.Code, rsp.Msg)
	sess := rsp.Session
	pterm.FgLightCyan.Printf("ID:            %s\n", sess.ID)
	pterm.Print("Name:          ")
	pterm.FgGreen.Printf("%s\n", sess.Name)
	pterm.Printf("Proto:         %s\n", sess.Proto)
	pterm.Printf("Config type:   %s\n", sess.ConfigType)
	pterm.Print("Option:\n")
	pterm.Printf("    Upload rate limit:   %d mbps\n", sess.Opt.SendRateLimit)
	pterm.Printf("    Download rate limit: %d mbps\n", sess.Opt.RecvRateLimit)
	if *isShowConfig {
		pterm.Printf("Config:\n%s\n", sess.Config)
	}
}

func createSession(cmd *cobra.Command, args []string) {
	rsp, err := srv.CreateSession(cmd.Context(), &api.CreateSessionReq{
		Name:       *createName,
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
	pterm.Printf("Name:         %s\n", *createName)
	pterm.Printf("Proto:         %s\n", rsp.Proto)
	pterm.Printf("Config type:   %s\n", rsp.ConfigType.String())
	pterm.Printf("Config:\n%s\n", rsp.Config)
}

func delSession(cmd *cobra.Command, args []string) {
	id := args[0]
	rsp, err := srv.DeleteSession(cmd.Context(), &api.DeleteSessionReq{IDorName: id})
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
	isShowConfig = getSessCmd.Flags().Bool("conf", false, "show the detail of config")
	getCmd.AddCommand(getSessCmd)

	createName = createSessCmd.Flags().String("name", "", "session name")
	proto = createSessCmd.Flags().String("proto", "trojan", "proto")
	ctype = createSessCmd.Flags().String("type", "json", "config type")
	limit = createSessCmd.Flags().Uint64("limit", 0, "speed limit")
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
