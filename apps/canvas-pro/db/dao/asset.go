package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type AssetDao struct {
	db.BaseDao[model.AssetModel]
}

func NewAssetDao(ctx context.Context, dao ...*db.BaseDao[model.AssetModel]) *AssetDao {
	var baseDao db.BaseDao[model.AssetModel]
	if len(dao) > 0 {
		baseDao = *dao[0]
	} else {
		baseDao = db.InitBaseDao[model.AssetModel](ctx)
	}
	return &AssetDao{BaseDao: baseDao}
}

var assetTableName = model.ProjectModel{}.TableName()

func (dao *AssetDao) DeleleGroupEffect(groupId, workspace uint64) error {
	return dao.GetOrm().Table(assetTableName).Where("group_id = ? AND workspace = ?", groupId, workspace).Updates(map[string]any{"group_id": 0}).Error

}
