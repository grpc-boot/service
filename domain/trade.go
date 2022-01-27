package domain

import (
	"service/common/abstract"
)

// Trade 交易聚合抽象
type Trade interface {
	abstract.PoolObject
}
