package client_api

import (
	"encoding/json"
	"github.com/alexpashkov/ports/internal/port_domain"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestSeedPorts(t *testing.T) {
	var received []port_domain.Port
	require.NoError(t, SeedPorts(openPortsDecoder(t), func(port *port_domain.Port) error {
		received = append(received, *port)
		return nil
	}))
	sortPorts(received)
	expected := readIntoSlice(t, openPortsDecoder(t))
	sortPorts(expected)
	require.Equal(t, len(expected), len(received))
	require.Empty(t, cmp.Diff(received, expected))
}

func openPortsDecoder(t testing.TB) *json.Decoder {
	t.Helper()
	f, err := os.Open(filepath.Join("fixtures", "ports.json"))
	require.NoError(t, err)
	return json.NewDecoder(f)
}

func readIntoSlice(t testing.TB, dec *json.Decoder) []port_domain.Port {
	t.Helper()
	m := make(map[string]port_domain.Port)
	require.NoError(t, dec.Decode(&m))
	var res []port_domain.Port
	for _, p := range m {
		res = append(res, p)
	}
	return res
}

func sortPorts(p []port_domain.Port) {
	sort.Sort(Ports(p))
}

type Ports []port_domain.Port

func (p Ports) Len() int {
	return len(p)
}

func (p Ports) Less(i, j int) bool {
	if len(p[i].Unlocs) < 1 || len(p[j].Unlocs) < 1 {
		return false
	}
	return p[i].Unlocs[0] < p[j].Unlocs[0]
}

func (p Ports) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
