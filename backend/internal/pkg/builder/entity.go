package builder

type IRunner interface {
	Run(bindAddress string) error
}
