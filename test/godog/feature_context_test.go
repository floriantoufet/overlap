package godog

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cucumber/godog"
)

var output string

// FeatureContext maps all steps regex with its functions
// Sets After and Before Scenario behaviours
func FeatureContext(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call the overlap CLI with ([^"]*) as first CIDR and ([^"]*) as second CIDR$`, iCallTheOverlapCLI)
	ctx.Step(`^I call the overlap CLI with ([^"]*) as first CIDR and ([^"]*) as second CIDR and call failed$`, iCallTheOverlapCLIFailed)
	ctx.Step(`^I call the overlap CLI with too much arguments$`, iCallTheOverlapCLIWithTooMuchArguments)
	ctx.Step(`^output result should be "([^"]*)"$`, outputResultShouldBe)
	ctx.Step(`^error message should be "([^"]*)"$`, errorMessageShouldBe)
}

func iCallTheOverlapCLIFailed(cidr1, cidr2 string) error {
	// Call CLI with given args
	output = ""
	cmd := exec.Command("./overlap", cidr1, cidr2)
	outputs, err := cmd.CombinedOutput()
	if err == nil {
		return errors.New("call should failed")
	}
	output = string(outputs)

	return nil
}

func iCallTheOverlapCLI(cidr1, cidr2 string) error {
	// Call CLI with given args
	output = ""
	cmd := exec.Command("./overlap", cidr1, cidr2)
	outputs, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("unable to get output: %w", err)
	}
	output = string(outputs)

	return nil
}

func iCallTheOverlapCLIWithTooMuchArguments() error {
	// Call CLI with given args
	output = ""
	cmd := exec.Command("./overlap", "foo", "bar", "hello")

	outputs, err := cmd.CombinedOutput()
	if err == nil {
		return errors.New("call should failed")
	}
	output = string(outputs)

	return nil
}

func outputResultShouldBe(result string) error {
	// Check output
	if result != output {
		return fmt.Errorf("unexpected result: exp=%s ; real=%s", result, output)
	}

	return nil
}

func errorMessageShouldBe(err string) error {
	// Check output
	if strings.Contains(err, output) {
		return fmt.Errorf("unexpected result: exp=%s ; real=%s", err, output)
	}

	return nil
}
