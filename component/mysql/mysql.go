package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Instance struct {
	Settings map[string]setting `json:"settings"` // 组件可能有多个实例，要想把多个配置项解析到Settings，需要在配置文件增加一层，名为settings
}

// Setting 配置项
type setting struct {
	DSN                       string        `json:"dsn"`
	MaxIdleConns              int           // 最大空闲连接
	MaxOpenConns              int           // 最大连接数
	ConnMaxLifetime           time.Duration // 连接可复用时间
	OpenAdvancedConfig        bool          // 开启高级配置
	DefaultStringSize         uint          // string 类型字段的默认长度
	DisableDatetimePrecision  bool          // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	DontSupportRenameIndex    bool          // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	DontSupportRenameColumn   bool          // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	SkipInitializeWithVersion bool          // 根据当前 MySQL 版本自动配置
}

var (
	componentName = "mysql"
	connects      = make(map[string]*gorm.DB)
)

func (i *Instance) Run() error {
	// 根据配置项个数创建对应数量的组件实例
	for k, s := range i.Settings {
		var db *gorm.DB
		var err error
		gcf := &gorm.Config{}
		if s.OpenAdvancedConfig {
			cf := mysql.Config{
				DSN:                       s.DSN,
				DefaultStringSize:         s.DefaultStringSize,
				DisableDatetimePrecision:  s.DisableDatetimePrecision,
				DontSupportRenameIndex:    s.DontSupportRenameIndex,
				DontSupportRenameColumn:   s.DontSupportRenameColumn,
				SkipInitializeWithVersion: s.SkipInitializeWithVersion,
			}
			db, err = gorm.Open(mysql.New(cf), gcf)
		} else {
			db, err = gorm.Open(mysql.Open(s.DSN), gcf)
		}

		if err != nil {
			return err
		}

		sqlDB, err := db.DB()
		if err != nil {
			return err
		}

		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(s.ConnMaxLifetime)

		connects[k] = db
	}

	return nil
}

func (i *Instance) Name() string {
	return componentName
}

func Main() *gorm.DB {
	return connects["main"]
}

// Get 获取指定mysql实例
func Get(name string) *gorm.DB {
	if c, ok := connects[name]; ok {
		return c
	}
	return nil
}
