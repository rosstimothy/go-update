package retriever

import "io"

type Retriever interface {
	Retrieve() (io.Reader, error)
}
