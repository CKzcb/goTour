/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: word
 * @Version: 1.0.0
 * @Date: 2023/3/10 15:17
 */

package word

import (
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"strings"
	"unicode"
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
			content = ToUpper(str)
		case ModeLower:
			content = ToLower(str)
		case ModeCamelCaseToUnderscore:
			content = CamelCaseToUnderscore(str)
		case ModeUnderscoreToUpperCamelCase:
			content = UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = UnderscoreToLowerCamelCase(str)
		default:
			log.Fatalf("system not support this mode %v", mode)

		}
		log.Println("result is: ", content)
	},
}

//
// ToUpper
//  @Description: make s to upper
//  @param s
//  @return string
//
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//
// ToLower
//  @Description: make s to lower
//  @param s
//  @return string
//
func ToLower(s string) string {
	return strings.ToLower(s)
}

//
// UnderscoreToUpperCamelCase
//  @Description:
//  @param s
//  @return string
//
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = cases.Title(language.Und, cases.NoLower).String(s)
	return strings.Replace(s, " ", "", -1)
}

//
// UnderscoreToLowerCamelCase
//  @Description:
//  @param s
//  @return string
//
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//
// CamelCaseToUnderscore
//  @Description:
//  @param s
//  @return string
//
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
