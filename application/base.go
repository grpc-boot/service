package application

type App interface {
	getDomain(ver string) (impl interface{})
}
