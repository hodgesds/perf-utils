package perf

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMSRPaths(t *testing.T) {
	_, err := MSRPaths()
	require.Nil(t, err)
}

func TestMSR(t *testing.T) {
	msrs, err := MSRPaths()
	require.Nil(t, err)

	msr, err := NewMSR(msrs[0])
	require.Nil(t, err)

	// TODO: This may only work on certain architectures :(
	_, err = msr.Read(0x00)
	require.Nil(t, err)

	require.Nil(t, msr.Close())

}

func TestMSRs(t *testing.T) {
	MSRs(func(err error) {
		require.Nil(t, err)
	})
}
