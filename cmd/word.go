package cmd

import (
	interval "helloCLI/interval/word"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

const (
	ModeUpper                      = iota + 1 // 全部大写
	ModeLower                                 // 全部小写
	ModeUnderscoreToUpperCamelCase            // 下划线转驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转小写驼峰
	UpperCamelCaseToUnderscore                // 驼峰转下划线
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = interval.ToUpper(str)
		case ModeLower:
			content = interval.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = interval.UnderscoreToUpperCamelClass(str)
		case ModeUnderscoreToLowerCamelCase:
			content = interval.UnderscoreToLowerCamelClass(str)
		case UpperCamelCaseToUnderscore:
			content = interval.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word 查看帮助文档")
		}
		log.Printf("输出结果: %s", content)
	},
}

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下:",
	"1: 全部大写",
	"2: 全部小写",
	"3: 下划线转驼峰",
	"4: 下划线转小写驼峰",
	"5: 驼峰转下划线",
}, "\n")

var str string
var mode uint8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输出字符串内容")
	wordCmd.Flags().Uint8VarP(&mode, "mode", "m", 0, "请输出转换模式")
}
