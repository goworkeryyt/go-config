/**
 * @Time: 2022/2/21 16:10
 * @Author: yt.yin
 */

package database

// SQLite 数据配置
type SQLite struct {

	/** sqlite 文件存放位置 */
	DbPath                string          `mapstructure:"db-path"                  json:"dbpath"                   yaml:"db-path"`

	/** 最大空闲连接数 */
	MaxIdleConns          int             `mapstructure:"max-idle-conns"           json:"maxIdleConns"             yaml:"max-idle-conns"`

	/** 最大连接数 */
	MaxOpenConns          int             `mapstructure:"max-open-conns"           json:"maxOpenConns"             yaml:"max-open-conns"`

	/** sql日志等级 */
	LogLevel              string          `mapstructure:"log-level"                json:"logLevel"                 yaml:"log-level"`

	/** 是否执行清除命令 */
	Vacuum                bool            `mapstructure:"vacuum"                   json:"vacuum"                   yaml:"vacuum"`

	/** 连接最大空闲时间  单位：秒 */
	ConnMaxIdleTime       int             `mapstructure:"conn-max-idle-time"       json:"connMaxIdleTime"          yaml:"conn-max-idle-time"`

	/** 连接最大生命周期  单位：秒 */
	ConnMaxLifetime       int             `mapstructure:"conn-max-life-time"       json:"connMaxLifetime"          yaml:"conn-max-life-time"`
}
