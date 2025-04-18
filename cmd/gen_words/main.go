package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	v1 "vocabulary/api/words/v1"
	"vocabulary/internal/dao"
	"vocabulary/internal/model/do"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

type Word struct {
	Word               string             `json:"word"`
	Definition         string             `json:"definition"`
	ExampleSentence    string             `json:"example_sentence"`
	ChineseTranslation string             `json:"chinese_translation"`
	Pronunciation      string             `json:"pronunciation"`
	ProficiencyLevel   v1.ProficiencyLevel `json:"proficiency_level"`
}

func main() {
	// 读取JSON文件
	data, err := os.ReadFile("cmd/gen_words/words.json")
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}

	var words []Word
	if err := json.Unmarshal(data, &words); err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return
	}

	ctx := context.Background()
	uid := uint(1)

	for _, w := range words {
		// 先检查是否存在
		count, err := dao.Words.Ctx(ctx).
			Where("uid", uid).
			Where("word", w.Word).
			Count()
		if err != nil {
			fmt.Printf("检查单词 %s 失败: %v\n", w.Word, err)
			continue
		}

		if count > 0 {
			// 如果存在则更新
			_, err = dao.Words.Ctx(ctx).Data(do.Words{
				Definition:         w.Definition,
				ExampleSentence:    w.ExampleSentence,
				ChineseTranslation: w.ChineseTranslation,
				Pronunciation:      w.Pronunciation,
				ProficiencyLevel:   w.ProficiencyLevel,
			}).Where("uid", uid).Where("word", w.Word).Update()
			
			if err != nil {
				fmt.Printf("更新单词 %s 失败: %v\n", w.Word, err)
			} else {
				fmt.Printf("更新单词: %s\n", w.Word)
			}
		} else {
			// 不存在则插入
			_, err = dao.Words.Ctx(ctx).Data(do.Words{
				Uid:                uid,
				Word:               w.Word,
				Definition:         w.Definition,
				ExampleSentence:    w.ExampleSentence,
				ChineseTranslation: w.ChineseTranslation,
				Pronunciation:      w.Pronunciation,
				ProficiencyLevel:   w.ProficiencyLevel,
			}).Insert()

			if err != nil {
				fmt.Printf("插入单词 %s 失败: %v\n", w.Word, err)
			} else {
				fmt.Printf("插入单词: %s\n", w.Word)
			}
		}
	}
}