package firsttry

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/liangzuobin/gomocktest/firsttry/mock_firsttry"
	"github.com/stretchr/testify/assert"
)

func TestFoolish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFace := mock_firsttry.NewMockFace(ctrl)

	mockFace.EXPECT().Bar().Return("hello world", nil)

	mockFace.EXPECT().Baz().Return("u r fooled", nil)

	err := Foolish(mockFace)

	assert.Nil(t, err, "what a paceful world")
}
