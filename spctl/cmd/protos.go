/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	pb "github.com/fregie/simple/proto/api"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// protosCmd represents the protos command
var protosCmd = &cobra.Command{
	Use:   "protos",
	Short: "Get available protos",
	Long:  `Get available protos`,
	Run:   getProtos,
}

func getProtos(cmd *cobra.Command, args []string) {
	rsp, err := srv.GetProtos(cmd.Context(), &pb.GetProtosReq{})
	checkErr(err)
	checkRsp(rsp.Code, rsp.Msg)
	list := []pterm.BulletListItem{}
	for _, proto := range rsp.Protos {
		list = append(list, pterm.BulletListItem{
			Level: 0,
			Text:  proto,
		})
	}
	pterm.FgLightCyan.Print("Support protos:\n")
	pterm.DefaultBulletList.WithItems(list).Render()
}

func init() {
	getCmd.AddCommand(protosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// protosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// protosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
