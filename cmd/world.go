/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: word
 * @Version: 1.0.0
 * @Date: 2023/3/10 15:17
 */

package cmd

import (
	"github.com/CKzcb/goTour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper                      = iota + 1 // words translate to upper of all
	ModeLower                                 // words translate to lower of all
	ModeUnderscoreToUpperCamelCase            // upper words translate to upper camel case
	ModeUnderscoreToLowerCamelCase            // lower words translate to lower camel case
	ModeCamelCaseToUnderscore                 // camel case translate to underscore
)

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "please input words.")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "please input select mode")
}

var desc = strings.Join([]string{
	"this sub cmd support the following cmd: ",
	"1. words translate to upper of all",
	"2. words translate to lower of all",
	"3. upper words translate to upper camel case",
	"4. lower words translate to lower camel case",
	"5. camel case translate to underscore",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		//mode:= args[0]
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		default:
			log.Fatalf("system not support this mode %v", mode)

		}
		log.Println("result is: ", content)
	},
}
