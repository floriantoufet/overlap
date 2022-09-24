package handler_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/floriantoufet/overlap/internal/handler"
	"github.com/floriantoufet/overlap/internal/processor"
)

func TestHandler_GetOverlapRelation(t *testing.T) {
	t.Parallel()

	var (
		cidrA = "10.0.2.0/24"
		cidrB = "10.0.2.10/24"
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		p := NewMockProcessor(ctrl)
		p.EXPECT().GetOverlapRelation(cidrA, cidrB).Return(processor.SameRelation, nil)

		h := handler.NewHandler(p)

		var output bytes.Buffer
		assert.Nil(t, h.PrintOverlapRelation([]string{cidrA, cidrB}, &output))
		assert.Equal(t, string(processor.SameRelation), output.String())
	})

	t.Run("fail when invalid number of argument", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		p := NewMockProcessor(ctrl)

		h := handler.NewHandler(p)

		var output bytes.Buffer
		assert.EqualError(t, h.PrintOverlapRelation([]string{cidrA, cidrB, ""}, &output), "invalid number of arguments")
	})

	t.Run("panic if get an error from processor", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		p := NewMockProcessor(ctrl)
		p.EXPECT().GetOverlapRelation(cidrA, cidrB).Return(processor.NetworkRelation(""), errors.New("some error"))

		h := handler.NewHandler(p)

		var output bytes.Buffer
		assert.EqualError(t, h.PrintOverlapRelation([]string{cidrA, cidrB}, &output), "some error")
	})
}
