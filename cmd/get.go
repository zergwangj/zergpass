// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/zergwangj/zergpass/db"
	"github.com/crackcell/gotabulate"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print a stored password or copy to the system clipboard",
	Long: `Print a stored password or copy to the system clipboard. Entry to print should be
specified by title.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("get called")
		title, err := cmd.Flags().GetString("title")
		if err != nil || len(title) == 0 {
			fmt.Println("Error: don't found entry's title")
			return
		}

		d := db.NewDB()
		defer d.Close()
		entry, err := d.Get(title)
		if err != nil {
			fmt.Println("Error: get entry failed - ", err.Error())
			return
		}
		tabulator := gotabulate.NewTabulator()
		tabulator.SetFirstRowHeader(true)
		tabulator.SetFormat("psql")
		tables := make([][]string, 0)
		tables = append(tables, db.FieldStrings())
		tables = append(tables, entry.ValueStrings())
		fmt.Print(tabulator.Tabulate(tables))
		fmt.Println("1 row(s) in db")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.Flags().StringP("title", "t", "", "password entry's title")
}
