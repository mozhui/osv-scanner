package reporter

import (
	"fmt"
	"io"

	"github.com/google/osv-scanner/internal/output"
	"github.com/google/osv-scanner/pkg/models"
)

type TableReporter struct {
	hasPrintedError bool
	stdout          io.Writer
	stderr          io.Writer
	markdown        bool
}

func NewTableReporter(stdout io.Writer, stderr io.Writer, markdown bool) *TableReporter {
	return &TableReporter{
		stdout:          stdout,
		stderr:          stderr,
		hasPrintedError: false,
		markdown:        markdown,
	}
}

func (r *TableReporter) PrintError(msg string) {
	fmt.Fprint(r.stderr, msg)
	r.hasPrintedError = true
}

func (r *TableReporter) HasPrintedError() bool {
	return r.hasPrintedError
}

func (r *TableReporter) PrintText(msg string) {
	fmt.Fprint(r.stdout, msg)
}

func (r *TableReporter) PrintResult(vulnResult *models.VulnerabilityResults) error {
	if r.markdown {
		output.PrintMarkdownTableResults(vulnResult, r.stdout)
	} else {
		output.PrintTableResults(vulnResult, r.stdout)
	}

	return nil
}
