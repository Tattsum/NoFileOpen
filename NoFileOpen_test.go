package NoFileOpen_test

import (
	"testing"

	"github.com/Tattsum/NoFileOpen"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, NoFileOpen.Analyzer, "a", "b")
}
