// Copyright © 2018 zergwangj <zergwangj@163.com>
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
	"github.com/zergwangj/zergpass/db"
	"fmt"
	"github.com/crackcell/gotabulate"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List database entries",
	Long: `List database entries.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		d := db.NewDB()
		defer d.Close()
		entries, err := d.List()
		if err != nil {
			fmt.Println("Error: list entries failed - ", err.Error())
			return
		}
		tabulator := gotabulate.NewTabulator()
		tabulator.SetFirstRowHeader(true)
		tabulator.SetFormat("psql")
		tables := make([][]string, 0)
		tables = append(tables, db.FieldStrings())
		for _, entry := range entries {
			tables = append(tables, entry.ValueStrings())
		}
		fmt.Print(tabulator.Tabulate(tables))
		fmt.Println(len(entries), "row(s) in db")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
