package util

import (
	"regexp"
	"strings"
)

// 注：每一个hashMap的值为：
// {
// 	“isEnd”：bool
// 	”key“：hashMap
// }
// 保存创建测敏感词hashMap
var sensitiveWordHashMap map[string]interface{} = map[string]interface{}{}

// 根据敏感词列表初始化敏感词hashMap
func init() {
	// 遍历所有敏感词
	for _, word := range sensitiveWordList {

		nowHashMap := sensitiveWordHashMap
		// 先敏感词中的字母、数字与字拆分开
		pretreatmentWord := Pretreatment(word)
		// 遍历每个敏感词的每个字
		for i, value := range pretreatmentWord {
			// 将rune转换为string
			keyWord := string(value)
			// 查询hashMap中是否存在keyWord，如果不存在，则证明已keyWord开头的敏感词还不存在，则我们直接构建这样的一棵树。否则进入下一次读取
			if nowHashMap[keyWord] != nil {
				nowHashMap = nowHashMap[keyWord].(map[string]interface{})
			} else {
				newHashMap := make(map[string]interface{})
				newHashMap["isEnd"] = false

				nowHashMap[keyWord] = newHashMap
				nowHashMap = newHashMap
			}
			// 判断该字是否为该词最后一字
			if i >= len(pretreatmentWord)-1 {
				nowHashMap["isEnd"] = true
			}
		} // for i,value : range pretreatmentWord {....
	} // for _, word : range sensitiveWordList {..
}

// 判断敏感词中是否包含字符串、数字
func isWordContainCharacter(word string) bool {
	reg, _ := regexp.Compile("[\x41-\x5A]|[\x61-\x7A]")
	return reg.Find([]byte(word)) != nil
}

// 拆分敏感词,使得相邻的英文字母及数字当成一个整体
func Pretreatment(word string) []string {
	// 返回结果
	splitResult := make([]string, 0)
	tempData := ""
	for _, value := range word {
		if isWordContainCharacter(string(value)) {
			tempData = tempData + string(value)
		} else {
			if tempData != "" {
				splitResult = append(splitResult, string(tempData))
				tempData = ""
			}
			splitResult = append(splitResult, string(value))
		}
	}
	if tempData != "" {
		splitResult = append(splitResult, string(tempData))
	}
	return splitResult
}

// 检测敏感词
func IsSensitiveWordByPlayerNick(nick string) bool {
	nowHashMap := sensitiveWordHashMap
	// 输入数据预处理
	pretreatmentWord := Pretreatment(nick)
	for _, word := range pretreatmentWord {

		keyWord := string(word)
		if nowHashMap[keyWord] != nil {

			nowHashMap = nowHashMap[keyWord].(map[string]interface{})

			if nowHashMap["isEnd"].(bool) {
				return true
			}
		}
	}
	return false
}

// 检测特殊字符
func IsSpecialCharactersByPlayerNick(playerNick string) bool {
	// 特殊字符
	// reg, _ := regexp.Compile("[\x09-\x0D]|[\x21-\x2F]|[\x3A-\x40]|[\x5B-\x60]|[\x7B-\x7E]")
	reg, _ := regexp.Compile("[\x09-\x0D]")
	if reg.Find([]byte(playerNick)) != nil {
		return true
	}

	// 屏蔽Emoji表情
	for _, v := range playerNick {
		test := string(v)
		if (test >= "\U0001f300" && test <= "\U0001f5ff") || (test >= "\U0001f910" && test <= "\U0001f918") || (test >= "\U0001f980" && test <= "\U0001f984") || test == "\U0001f9c0" || (test >= "\U0001f600" && test <= "\U0001f64f") || (test >= "\U0001f680" && test <= "\U0001f6d0") || (test >= "\U0001f6e0" && test <= "\U0001f6ec") || (test >= "\U0001f6f0" && test <= "\U0001f6f3") || (test >= "\U00012600" && test <= "\U000127bf") {
			return true
		}
	}
	for i := 0; i < len(flags); i++ {
		if strings.Contains(playerNick, flags[i]) {
			return true
		}
	}
	return false
}
