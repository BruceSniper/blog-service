/**
 * @Author:      zhangyuyao
 * @Description: TODO
 * @File:        setting.go
 * @Version:     1.0.0
 * @Date:        2021/8/22 3:36 下午
 */

package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"

	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
