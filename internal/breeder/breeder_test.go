package breeder

import (
	"testing"
	"train/internal/animal"
	"train/mock"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestBreeder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mock.NewMockAInterface(mockCtrl)

	t.Run("ペットに足がある時、エラーは起こらない", func(t *testing.T) {
		m.EXPECT().Eat("meat").Times(1)
		m.EXPECT().Walk().Times(1)

		ac := NewBreeder(m)
		result := ac.Breeder()

		require.NoError(t, result)
	})

	t.Run("ペットに足がない場合、エラーを返す", func(t *testing.T) {
		e := animal.AStruct{Name: "Test", HasLegs: false}
		ac := NewBreeder(&e)

		result := ac.Breeder()

		require.Error(t, result)
	})
}
