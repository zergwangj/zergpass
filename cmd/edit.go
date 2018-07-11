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

const hiddenString = "!@#$%^&*"

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing database entry",
	Long: `Edit an existing database entry.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("edit called")
		title, err := cmd.Flags().GetString("title")
		if err != nil || len(title) == 0 {
			fmt.Println("Error: don't found entry's title")
			return
		}
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		url, _ := cmd.Flags().GetString("url")
		notes, _ := cmd.Flags().GetString("notes")

		d := db.NewDB()
		defer d.Close()
		entry, err := d.Get(title)
		if username != hiddenString && len(username) > 0 {
			entry.Username = username
		}
		if password != hiddenString && len(password) > 0 {
			entry.Password = password
		}
		if url != hiddenString && len(url) > 0 {
			entry.Url = url
		}
		if notes != hiddenString && len(notes) > 0 {
			entry.Notes = notes
		}
		err = d.Set(entry)
		if err != nil {
			fmt.Println("Error: edit entry failed - ", err.Error())
			return
		}
		fmt.Println("Edit entry success")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	editCmd.Flags().StringP("title", "t", "", "password entry's title")
	editCmd.Flags().StringP("url", "l", hiddenString, "password entry's URL")
	editCmd.Flags().StringP("username", "u", hiddenString, "password entry's username")
	editCmd.Flags().StringP("password", "p", hiddenString, "password entry's password")
	editCmd.Flags().StringP("notes", "n", hiddenString, "password entry's notes")
}
