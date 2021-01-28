package global

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	DB  *gorm.DB
	Log *zap.SugaredLogger
)