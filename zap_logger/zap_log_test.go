package zap_logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestZapLogger(t *testing.T) {
	// 测试Info级别日志
	logger.Info("用户value1执行了操作，处理次数为100")

	// 测试Error级别日志
	logger.Error("系统发生错误：error_value，错误代码：500")

	// 测试带有结构化字段的日志
	logger.Info("用户张三，年龄25岁，兴趣爱好包括：读书、运动")

	// 测试Debug级别日志
	logger.Debug("调试信息：debug_value")

	// 测试Warn级别日志
	logger.Warn("警告信息：warn_value")
}
func TestZapLoggerJSON(t *testing.T) {
	level.SetLevel(zap.DebugLevel)
	// 测试Info级别日志,带有额外字段
	logger.Info("这是一条测试Info日志",
		zap.String("key1", "value1"),
		zap.Int("key2", 100))

	// 测试Error级别日志,带有错误信息
	logger.Error("这是一条测试Error日志",
		zap.String("error_key", "error_value"),
		zap.Int("error_code", 500))

	// 测试带有结构化字段的日志
	logger.Info("带结构化字段的日志",
		zap.String("user", "张三"),
		zap.Int("age", 25),
		zap.Strings("hobbies", []string{"读书", "运动"}))

	// 测试Warn级别日志
	logger.Warn("这是一条测试Warn日志",
		zap.String("warn_key", "warn_value"))

	// 测试Debug级别日志
	logger.Debug("调试信息：debug_value")
}
