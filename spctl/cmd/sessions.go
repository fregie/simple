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
	pb "github.com/fregie/simple/proto/gen/go/api"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// sessionsCmd represents the sessions command
var sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "Get all sessions",
	Long:  `Get all sessions`,
	Run:   showSessions,
}

func showSessions(cmd *cobra.Command, args []string) {
	rsp, err := srv.GetAllSessions(cmd.Context(), &pb.GetAllSessionsReq{})
	checkErr(err)
	checkRsp(rsp.Code, rsp.Msg)
	data := make([][]string, 0)
	data = append(data, []string{"Name", "ID", "proto", "config type"})
	for _, sess := range rsp.Sessions {
		data = append(data, []string{sess.Name, sess.ID, sess.Proto, sess.ConfigType.String()})
	}
	pterm.DefaultTable.WithHasHeader().WithData(data).Render()
}

func init() {
	getCmd.AddCommand(sessionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sessionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sessionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
