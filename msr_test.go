package perf

import (
	"os"
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

	msr, err := NewMSR(msrs[0], os.O_RDWR, 0660)
	require.Nil(t, err)

	// TODO: This may only work on certain architectures :(
	buf := make([]byte, 8)
	err = msr.Read(0x00, buf)
	require.Nil(t, err)

	require.Nil(t, msr.Close())

}

func TestMSRs(t *testing.T) {
	MSRs(os.O_RDWR, 0660, func(err error) {
		require.Nil(t, err)
	})
}
