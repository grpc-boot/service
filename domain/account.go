package domain

import "service/infrastructure/abstract"

// Account 账户聚合抽象
type Account interface {
	abstract.PoolObject
}
