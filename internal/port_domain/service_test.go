package port_domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService(t *testing.T) {
	s := NewService()
	const id = "foo"
	p := newPortWithID(id)
	t.Run("stores the port", func(t *testing.T) {
		require.NoError(t, s.UpsertPort(p))
		require.Equal(t, s.GetPortByID(id), p)
	})
	t.Run("returns nil when there is no such port", func(t *testing.T) {
		require.Nil(t, s.GetPortByID("no such id"))
	})
	t.Run("updates existing port", func(t *testing.T) {
		newPort := newPortWithID(id)
		require.NoError(t, s.UpsertPort(newPort))
		require.Equal(t, s.GetPortByID(id), newPort)
	})
}

func newPortWithID(id string) *Port {
	return &Port{
		Name:     id,
		City:     id,
		Country:  id,
		Province: id,
		Timezone: id,
		Unlocs:   []string{id},
		Code:     id,
	}
}
