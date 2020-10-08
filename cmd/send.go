/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"errors"

	"github.com/TouchSafe3/ts2-api/accesscontrolservice"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send test board command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("need board command")
		}
		_, err := accesscontrolservice.BoardCommandFromString(args[0])
		return err

	},
	Run: func(cmd *cobra.Command, args []string) {
		bc, _ := accesscontrolservice.BoardCommandFromString(args[0])
		//error checking already done
		send(bc)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func send(bc accesscontrolservice.BoardCommand) {

	//TODO ADD ARG FOR ENDPOINT
}
