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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one entry from database",
	Long: `Delete one entry from database. Entry to delete should be
specified by title.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("delete called")
		title, err := cmd.Flags().GetString("title")
		if err != nil || len(title) == 0 {
			fmt.Println("Error: don't found entry's title")
			return
		}

		d := db.NewDB()
		defer d.Close()
		err = d.Delete(title)
		if err != nil {
			fmt.Println("Error: delete entry failed - ", err.Error())
			return
		}
		fmt.Println("Delete entry success")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().StringP("title", "t", "", "password entry's title")
}
