package config

import (
	"github.com/bijialin/nb-boot/common/base"
	"github.com/bijialin/nb-boot/common/global"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func init() {
	initConfig()
	initLog()
	initDb()
	initRedis()

}

func initConfig() {
	configPath := "./config"
	configType := "yml"
	configPrefix := "app"
	v := viper.New()
	v.SetConfigType(configType)
	// 从default中读取默认的配置
	v.SetConfigName(configPrefix)
	v.AddConfigPath(configPath)
	err := v.ReadInConfig()
	if err != nil {
		return
	}

	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}
	env := os.Getenv("run-mode")
	if env != "" {
		viper.SetDefault("run-mode", env)
	}
	vEnv := v.GetString("run-mode")

	if vEnv == "" {
		viper.SetDefault("run-mode", "dev")
	}

	// 根据配置的env读取相应的配置信息
	if env != "" {
		viper.SetConfigName(configPrefix + "-" + env)
	} else {
		viper.SetConfigName(configPrefix + "-" + vEnv)
	}
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	err = viper.ReadInConfig()
	if err != nil {
		panic("init project config error， please check your config file")
	}
	bindErr := viper.Unmarshal(&global.AppConfig)
	if bindErr != nil {
		panic("init project config error， please check your config file")
	}
	//autoLoad := AppConfig.AutoLoadConfig
	//if autoLoad {
	//	viper
	//}
}

func initDb() {
	global.Log.Info("start init nb-boot db ")

	//dbUrl := viper.GetString("database.mysql.url")
	dbUrl := global.AppConfig.Database.Mysql.Url
	if dbUrl != "" {
		// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
		crtDb, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
		if err != nil {
			panic("init db error")
		}
		global.Db = crtDb
	}

	global.Log.Info("finished init nb-boot db ")

}

func initRedis() {

	//redisUrl := viper.GetString("redis.url")
	redisUrl := global.AppConfig.Redis.Host
	if redisUrl != "" {
		global.Log.Info("start init nb-boot db ")
	}
}

func initLog() {
	env := viper.Get("run-mode")
	if env == "prod" {
		writeSyncer := getLogWriter(global.AppConfig.Log)
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

		logger := zap.New(core, zap.AddCaller())
		global.Log = logger.Sugar()
	} else {
		l, _ := zap.NewDevelopment()
		global.Log = l.Sugar()
	}

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(config base.Log) zapcore.WriteSyncer {
	path := config.Path
	if path == "" {
		path = "./"
	}

	fileName := config.FileName
	if fileName == "" {
		fileName = "app.log"
	}

	if config.PackLog {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   path + fileName,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		}
		return zapcore.AddSync(lumberJackLogger)
	}
	file, _ := os.Create(path + fileName)
	return zapcore.AddSync(file)
}
