package words

import (
	"context"

	"vocabulary/api/words/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	v, err := c.words.Detail(ctx, req.Id, uid)
	if err != nil {
		return nil, err
	}
	return &v1.DetailRes{
		Id:                 v.Id,
		Word:               v.Word,
		Definition:         v.Definition,
		ExampleSentence:    v.ExampleSentence,
		ChineseTranslation: v.ChineseTranslation,
		Pronunciation:      v.Pronunciation,
		ProficiencyLevel:   v1.ProficiencyLevel(v.ProficiencyLevel),
		CreatedAt:          v.CreatedAt,
		UpdatedAt:          v.UpdatedAt,
	}, nil
}
