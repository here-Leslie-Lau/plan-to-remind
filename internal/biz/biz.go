package biz

import (
	"github.com/google/wire"
	"plan-to-remind/internal/pkg/json"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCronSpecUsecase, NewPlanUsecase, NewTimerUsecase,
	json.NewParser,
)
