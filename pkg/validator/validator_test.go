package validator_test

import (
	"testing"

	"github.com/atmxlab/proxychecker/pkg/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidator_Empty(t *testing.T) {
	t.Parallel()

	v := validator.New()
	require.NoError(t, v.Err())
}

func TestValidator_Failed(t *testing.T) {
	t.Parallel()

	t.Run("nonempty", func(t *testing.T) {
		t.Parallel()

		v := validator.New()
		v.Failed("msg")
		require.ErrorContains(t, v.Err(), "msg")
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		v := validator.New()
		v.Failed("")
		require.NoError(t, v.Err())
	})
}

func TestValidator_Failedf(t *testing.T) {
	t.Parallel()

	t.Run("nonempty", func(t *testing.T) {
		t.Parallel()

		v := validator.New()
		v.Failedf("msg: %d", 1)
		require.ErrorContains(t, v.Err(), "msg")
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		v := validator.New()
		v.Failedf("")
		require.NoError(t, v.Err())
	})
}

func TestValidator_AddErr(t *testing.T) {
	t.Parallel()

	t.Run("nonnil", func(t *testing.T) {
		t.Parallel()

		v := validator.New()
		v.AddErr(assert.AnError)
		require.ErrorContains(t, v.Err(), assert.AnError.Error())
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		v := validator.New()
		v.AddErr(nil)
		require.NoError(t, v.Err())
	})
}

func TestValidator_WrapErr(t *testing.T) {
	t.Parallel()

	t.Run("nonnil", func(t *testing.T) {
		t.Parallel()

		v := validator.New()
		v.WrapErr(assert.AnError, "msg")
		require.ErrorContains(t, v.Err(), assert.AnError.Error())
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		v := validator.New()
		v.WrapErr(nil, "msg")
		require.NoError(t, v.Err())
	})
}
