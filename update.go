package update

import (
	"github.com/rosstimothy/go-update/retriever"
	"fmt"
	"io"
	"log"
)


func Update(retriever retriever.Retriever, options Options) error {
	reader, err := retriever.Retrieve()
	if err != nil {
		return fmt.Errorf("unable to retrieve file: %v", err)
	}

	if reader == nil {
		log.Println("skipped update")
	}

	if closer, ok := reader.(io.Closer); ok {
		defer closer.Close()
	}

	err = apply(reader, options)
	if err != nil {
		return fmt.Errorf("failed to apply update: %v", err)
	}

	return nil
}