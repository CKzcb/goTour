/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: root
 * @Version: 1.0.0
 * @Date: 2023/3/10 15:19
 */

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
}
