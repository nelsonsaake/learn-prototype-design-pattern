package contracts

type Prototype interface {
	Clone() Prototype
}
