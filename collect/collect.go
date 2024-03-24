package collect

type Fetch interface {
	Get(url string) ([]byte, error)
}
