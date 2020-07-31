package prmchk_test

import (
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/prmchk"
	"testing"
)

func Test_EmptyStruct(t *testing.T) {
	type (
		X struct{}
	)
	require.False(t, prmchk.Check(X{}))
}

func Test_NonEmptyStruct(t *testing.T) {
	type (
		X struct {
			x int
		}
	)
	require.False(t, prmchk.Check(X{}))
}

func Test_EmbeddedType_NoConflict(t *testing.T) {
	type (
		X struct {
			x int
		}
		Y struct {
			X
			y int
		}
	)
	require.False(t, prmchk.Check(Y{}))
}

func Test_EmbeddedType_Conflict(t *testing.T) {
	type (
		X struct {
			x int
		}
		Y struct {
			X
			x int
		}
	)
	require.True(t, prmchk.Check(Y{}))
}

func Test_DeeplyEmbeddedType_NoConflict(t *testing.T) {
	type (
		X struct {
			x int
		}
		Y struct {
			X
		}
		Z struct {
			Y
			z int
		}
	)
	require.False(t, prmchk.Check(Z{}))
}

func Test_DeeplyEmbeddedType_ConflictZ(t *testing.T) {
	type (
		X struct {
			x int
		}
		Y struct {
			X
		}
		Z struct {
			Y
			x int
		}
	)
	require.True(t, prmchk.Check(Z{}))
}

func Test_DeeplyEmbeddedType_ConflictY(t *testing.T) {
	type (
		X struct {
			x int
		}
		Y struct {
			X
			x int
		}
		Z struct {
			Y
		}
	)
	require.True(t, prmchk.Check(Z{}))
}

func Test_TwoEmbeddedType_NoConflictX(t *testing.T) {
	type (
		X1 struct {
			x int
		}
		X2 struct {
			x int
		}
		Y struct {
			X1
			X2
		}
	)
	require.False(t, prmchk.Check(Y{}))
}

func Test_TwoEmbeddedType_ConflictX1(t *testing.T) {
	type (
		X1 struct {
			x1 int
		}
		X2 struct {
			x2 int
		}
		Y struct {
			X1
			X2
			x1 int
		}
	)
	require.True(t, prmchk.Check(Y{}))
}

func Test_TwoEmbeddedType_ConflictX2(t *testing.T) {
	type (
		X1 struct {
			x1 int
		}
		X2 struct {
			x2 int
		}
		Y struct {
			X1
			X2
			x2 int
		}
	)
	require.True(t, prmchk.Check(Y{}))
}
