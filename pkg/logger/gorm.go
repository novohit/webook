package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

// GormLoggerV1 方法一
type GormLoggerV1 struct {
	logger ZapLogger
}

func NewGormLoggerV1(logger ZapLogger) *GormLoggerV1 {
	return &GormLoggerV1{logger: logger}
}

func (l *GormLoggerV1) Printf(msg string, args ...interface{}) {
	//l.logger.Info(msg, zap.Any("args", args))
	//l.logger.Info(msg)
	fmt.Printf(msg, args...)
	fmt.Println()
}

// GormLoggerV2 方法二
type GormLoggerV2 struct {
	logger ZapLogger
}

func NewGormLoggerV2(logger ZapLogger) *GormLoggerV2 {
	return &GormLoggerV2{logger: logger}
}

func (l *GormLoggerV2) LogMode(level logger.LogLevel) logger.Interface {
	//TODO implement me
	return &GormLoggerV2{}
}

func (l *GormLoggerV2) Info(ctx context.Context, msg string, data ...interface{}) {
	//TODO implement me
	//panic("implement me")
	l.logger.Info(msg, zap.Any("data", data))
}

func (l *GormLoggerV2) Warn(ctx context.Context, msg string, data ...interface{}) {
	//TODO implement me
	l.logger.Info(msg, zap.Any("data", data))
}

func (l *GormLoggerV2) Error(ctx context.Context, msg string, data ...interface{}) {
	//TODO implement me
	l.logger.Info(msg, zap.Any("data", data))
}

func (l *GormLoggerV2) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//TODO implement me
	// 获取运行时间
	duration := time.Since(begin).Milliseconds()
	// 获取 SQL 语句和返回条数
	sql, rows := fc()
	// Gorm 错误时记录错误日志
	if err != nil {
		l.logger.Error("SQL ERROR", zap.String("sql", sql), zap.Int64("rows", rows), zap.Int64("dur(ms)", duration))
	}
	l.logger.Info("SQL DEBUG", zap.String("sql", sql), zap.Int64("rows", rows), zap.Int64("dur(ms)", duration))
	// 慢查询日志
	//if duration > l.SlowThreshold.Milliseconds() {
	//	logger.New(ctx).Warn("SQL SLOW", "sql", sql, "rows", rows, "dur(ms)", duration)
	//} else {
	//	logger.New(ctx).Debug("SQL DEBUG", "sql", sql, "rows", rows, "dur(ms)", duration)
	//}
}
