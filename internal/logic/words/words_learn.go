package words

import (
	"context"
	v1 "vocabulary/api/words/v1"
	"vocabulary/internal/dao"
	"vocabulary/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (w *Words) Rand(ctx context.Context, uid uint, limit uint) (words []entity.Words, err error) {
	if (limit <= 0) {
		limit = 10
	}

	var (
		cls = dao.Words.Columns()
		orm = dao.Words.Ctx(ctx)
		list = make([]entity.Words, limit)
	)

	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}

	err = orm.Limit(int(limit)).OrderRandom().Scan(&list)
	return list, err
}

func (w *Words) SetLevel(ctx context.Context, uid, id uint, level v1.ProficiencyLevel) (err error) {
	if level < 1 || level > 5 {
		return gerror.New("熟练度必须在1-5之间")
	}

	var (
		cls = dao.Words.Columns()
		orm = dao.Words.Ctx(ctx)
	)

	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}

	_, err = orm.Data(cls.ProficiencyLevel, level).Where(cls.Id, id).Update()
	return
}	

