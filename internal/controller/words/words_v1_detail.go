package words

import (
	"context"
	"vocabulary/api/words/v1"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	word, err := c.words.Detail(ctx, uid, req.Id)
	if err != nil {
		return nil, err
	}
	
	// 检查单词是否存在
	if word == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "单词不存在")
	}
	
	return &v1.DetailRes{
		Id:                 word.Id,
		Word:               word.Word,
		Definition:         word.Definition,
		ExampleSentence:    word.ExampleSentence,
		ChineseTranslation: word.ChineseTranslation,
		Pronunciation:      word.Pronunciation,
		ProficiencyLevel:   v1.ProficiencyLevel(word.ProficiencyLevel),
		CreatedAt:          word.CreatedAt,
		UpdatedAt:          word.UpdatedAt,
	}, nil
}
