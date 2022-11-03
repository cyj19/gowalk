package mysql

import (
	"github.com/cyj19/gowalk/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Instance struct {
	Settings map[string]setting `json:"settings"` // 组件可能有多个实例，要想把多个配置项解析到Settings，需要在配置文件增加一层，名为settings
}

// setting 配置项
type setting struct {
	DSN                       string        `json:"dsn" mapstructure:"dsn"`
	MaxIdleConns              int           `json:"max_idle_conns" mapstructure:"max_idle_conns"`                             // 最大空闲连接
	MaxOpenConns              int           `json:"max_open_conns" mapstructure:"max_open_conns"`                             // 最大连接数
	ConnMaxLifetime           time.Duration `json:"conn_max_lifetime" mapstructure:"onn_max_lifetime"`                        // 连接可复用时间
	OpenAdvancedConfig        bool          `json:"open_advanced_config" mapstructure:"open_advanced_config"`                 // 开启高级配置
	DefaultStringSize         uint          `json:"default_string_size" mapstructure:"default_string_size"`                   // string 类型字段的默认长度
	DisableDatetimePrecision  bool          `json:"disable_datetime_precision" mapstructure:"disable_datetime_precision"`     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	DontSupportRenameIndex    bool          `json:"dont_support_rename_index" mapstructure:"dont_support_rename_index"`       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	DontSupportRenameColumn   bool          `json:"dont_support_rename_column" mapstructure:"dont_support_rename_column"`     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	SkipInitializeWithVersion bool          `json:"skip_initialize_with_version" mapstructure:"skip_initialize_with_version"` // 根据当前 MySQL 版本自动配置
	SlowThreshold             time.Duration `json:"slow_threshold" mapstructure:"slow_threshold"`
	Colorful                  bool          `json:"colorful" mapstructure:"colorful"`
	IgnoreRecordNotFoundError bool          `json:"ignore_record_not_found_error" mapstructure:"ignore_record_not_found_error"`
	LogLevel                  int           `json:"log_level" mapstructure:"log_level"`
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
		l := New(logx.Log(), logger.Config{
			SlowThreshold:             s.SlowThreshold,
			Colorful:                  s.Colorful,
			IgnoreRecordNotFoundError: s.IgnoreRecordNotFoundError,
			LogLevel:                  logger.LogLevel(s.LogLevel),
		})
		gcf := &gorm.Config{Logger: l}
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
