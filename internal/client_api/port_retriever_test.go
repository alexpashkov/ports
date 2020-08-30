package client_api

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestPortRetriever(t *testing.T) {
	retriever, err := buildPortRetriever(openPortsDecoder(t))
	require.NoError(t, err)
	received := readIntoSliceFromRetriever(t, retriever)
	require.NoError(t, err)
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

func readIntoSliceFromRetriever(t testing.TB, retrievePort portRetriever) []Port {
	t.Helper()
	var (
		ports []Port
		p     Port
		err   error
	)
	for p, err = retrievePort(); err == nil; p, err = retrievePort() {
		ports = append(ports, p)
	}
	require.Equal(t, err, io.EOF)
	return ports
}

func readIntoSlice(t testing.TB, dec *json.Decoder) []Port {
	t.Helper()
	m := make(map[string]Port)
	require.NoError(t, dec.Decode(&m))
	var res []Port
	for _, p := range m {
		res = append(res, p)
	}
	return res
}

func sortPorts(p []Port) {
	sort.Sort(Ports(p))
}

type Ports []Port

func (p Ports) Len() int {
	return len(p)
}

func (p Ports) Less(i, j int) bool {
	return p[i].ID() < p[j].ID()
}

func (p Ports) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
