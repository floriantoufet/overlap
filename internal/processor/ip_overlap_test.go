package processor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/floriantoufet/overlap/internal/processor"
)

func TestGetOverlapRelation(t *testing.T) {
	t.Parallel()

	t.Run("success with subset relation", func(t *testing.T) {
		t.Parallel()

		cidrA := "10.0.0.0/20"
		cidrB := "10.0.2.0/24"

		p := processor.NewProcessor()
		relation, err := p.GetOverlapRelation(cidrA, cidrB)

		assert.Nil(t, err)
		assert.Equal(t, relation, processor.SubsetRelation)
	})

	t.Run("success with superset relation", func(t *testing.T) {
		t.Parallel()

		cidrA := "10.0.2.0/24"
		cidrB := "10.0.0.0/20"

		p := processor.NewProcessor()
		relation, err := p.GetOverlapRelation(cidrA, cidrB)

		assert.Nil(t, err)
		assert.Equal(t, relation, processor.SupersetRelation)
	})

	t.Run("success with different relation", func(t *testing.T) {
		t.Parallel()

		cidrA := "10.0.2.0/24"
		cidrB := "10.0.3.0/24"

		p := processor.NewProcessor()
		relation, err := p.GetOverlapRelation(cidrA, cidrB)

		assert.Nil(t, err)
		assert.Equal(t, relation, processor.DifferentRelation)
	})

	t.Run("success with same relation relation", func(t *testing.T) {
		t.Parallel()

		cidrA := "10.0.2.0/24"
		cidrB := "10.0.2.10/24"

		p := processor.NewProcessor()
		relation, err := p.GetOverlapRelation(cidrA, cidrB)

		assert.Nil(t, err)
		assert.Equal(t, relation, processor.SameRelation)
	})

	t.Run("fail if invalid first CIDR", func(t *testing.T) {
		t.Parallel()

		cidrA := "foo"
		cidrB := "10.0.2.10/24"

		p := processor.NewProcessor()
		_, err := p.GetOverlapRelation(cidrA, cidrB)

		assert.EqualError(t, err, "first CIDR invalid: unable to parse CIDR: netip.ParsePrefix(\"foo\"): no '/'")
	})

	t.Run("fail if invalid second CIDR", func(t *testing.T) {
		t.Parallel()

		cidrA := "10.0.2.10/24"
		cidrB := "foo"

		p := processor.NewProcessor()
		_, err := p.GetOverlapRelation(cidrA, cidrB)

		assert.EqualError(t, err, "second CIDR invalid: unable to parse CIDR: netip.ParsePrefix(\"foo\"): no '/'")
	})

	t.Run("fail if second CIDR is not IPv4", func(t *testing.T) {
		t.Parallel()

		cidrA := "10.0.2.10/24"
		cidrB := "2001:db8::/32"

		p := processor.NewProcessor()
		_, err := p.GetOverlapRelation(cidrA, cidrB)

		assert.EqualError(t, err, "second CIDR invalid: not an IPv4 IP")
	})

	t.Run("fail if first CIDR is not IPv4", func(t *testing.T) {
		t.Parallel()

		cidrA := "2001:db8::/32"
		cidrB := "10.0.2.10/24"

		p := processor.NewProcessor()
		_, err := p.GetOverlapRelation(cidrA, cidrB)

		assert.EqualError(t, err, "first CIDR invalid: not an IPv4 IP")
	})
}
