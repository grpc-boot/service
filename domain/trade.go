package domain

import "service/infrastructure/abstract"

// Trade 交易聚合抽象
type Trade interface {
	abstract.PoolObject
}
