package storage

import (
	"time"

	"github.com/golang/mock/gomock"

	"github.com/m3db/m3/src/dbnode/storage/index"
)

func MockAggregateIterator(iter *index.MockAggregateIterator) {
	gomock.InOrder(
		iter.EXPECT().Done().Return(false),
		iter.EXPECT().Done().Return(false),
		iter.EXPECT().SearchDuration().Return(time.Second),
		iter.EXPECT().Close().Return(nil).Times(2))
}