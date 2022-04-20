package log

import (
	"github.com/go-kratos/kratos/v2/log"
	"os"
	"testing"
)

func TestInfo(t *testing.T) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	helper := NewHelper(logger)
	helper.Info("demo", "name", "leslie", "age", "18")
}
