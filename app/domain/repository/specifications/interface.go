package specifications

import (
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// I defines persistence query
type I interface {
	GormQuery(db *gorm.DB) *gorm.DB
	CrucialQueryCondition() string
}

type specWrapper struct {
	i      I
	isRead bool
}

func (sw specWrapper) GormQuery(db *gorm.DB) *gorm.DB {
	tx := sw.i.GormQuery(db)

	if sw.isRead {
		tx = tx.Clauses(dbresolver.Read)
	}

	return tx
}

func (sw specWrapper) CrucialQueryCondition() string {
	return sw.i.CrucialQueryCondition()
}

func WithReplica(spec I) I {
	return specWrapper{
		i:      spec,
		isRead: true,
	}
}
