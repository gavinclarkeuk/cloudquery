package main

import (
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	outputPath := "./docs"
	if err := docs.GenerateDocs(provider.Provider(), outputPath, true); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	}
}
