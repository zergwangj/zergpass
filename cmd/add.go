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
	"fmt"
	"github.com/zergwangj/zergpass/db"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry to password database",
	Long: `Add a new entry to password database. You will supply values for
the entry's fields.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("add called")
		title, err := cmd.Flags().GetString("title")
		if err != nil || len(title) == 0 {
			fmt.Println("Error: don't found entry's title")
			return
		}
		username, err := cmd.Flags().GetString("username")
		if err != nil || len(username) == 0 {
			fmt.Println("Error: don't found entry's username")
			return
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil || len(password) == 0 {
			fmt.Println("Error: don't found entry's password")
			return
		}
		url, _ := cmd.Flags().GetString("url")
		notes, _ := cmd.Flags().GetString("notes")

		entry := db.NewEntry()
		entry.Title = title
		entry.Username = username
		entry.Password = password
		entry.Url = url
		entry.Notes = notes

		d := db.NewDB()
		defer d.Close()
		err = d.Add(entry)
		if err != nil {
			fmt.Println("Error: add entry failed - ", err.Error())
			return
		}
		fmt.Println("Add entry success")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("title", "t", "", "password entry's title")
	addCmd.Flags().StringP("url", "l", "", "password entry's URL")
	addCmd.Flags().StringP("username", "u", "", "password entry's username")
	addCmd.Flags().StringP("password", "p", "", "password entry's password")
	addCmd.Flags().StringP("notes", "n", "", "password entry's notes")
}
