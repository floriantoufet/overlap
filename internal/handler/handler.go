// Package handler contains CLI handler.
package handler

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/floriantoufet/overlap/internal/processor"
)

//go:generate mockgen -destination mock_test.go -source $GOFILE -package handler_test

// ErrInvalidNumberArgs when too much or not enough arguments given.
var ErrInvalidNumberArgs = errors.New("invalid number of arguments")

type (
	// Processor used to get relation between two CIDRs.
	Processor interface {
		GetOverlapRelation(cidrA, cidrB string) (processor.NetworkRelation, error)
	}

	// Handler used to get relation between two CIDRs and print result in given output.
	Handler struct {
		processor Processor
	}
)

// NewHandler returns a new implementation of Handler.
func NewHandler(processor Processor) *Handler {
	return &Handler{processor: processor}
}

// PrintOverlapRelation prints overlap relation between given CIDRs and print result in given output
// returns an error if number of argument is more or less than 2 CIDRs
// returns an error if processor returns an error or unable to write into output.
func (h *Handler) PrintOverlapRelation(args []string, output io.Writer) error {
	// Check args.
	if len(args) != 2 {
		return ErrInvalidNumberArgs
	}

	// Call processor
	relation, err := h.processor.GetOverlapRelation(strings.TrimSpace(args[0]), strings.TrimSpace(args[1]))
	if err != nil {
		return err
	}

	// Write relation in given output.
	if _, err := output.Write([]byte(relation)); err != nil {
		return fmt.Errorf("unable to write relation: %w", err)
	}

	// Success
	return nil
}
