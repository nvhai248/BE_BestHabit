package statisticalbiz

import (
	"bestHabit/modules/statistical/statisticalmodel"
	"fmt"
)

type CountUserStorage interface {
	CountUserByTimeCreated(time string) (int, error)
}

type CountHabitStorage interface {
	CountHabitByTimeCreated(time string) (int, error)
}

type CountTaskStorage interface {
	CountTaskByTimeCreated(time string) (int, error)
}

type CountChallengeStorage interface {
	CountChallengesByTimeCreated(time string) (int, error)
}

type statisticalBiz struct {
	countUserStorage      CountUserStorage
	countHabitStorage     CountHabitStorage
	countTaskStorage      CountTaskStorage
	countChallengeStorage CountChallengeStorage
}

func NewStatisticalBiz(countUserStorage CountUserStorage,
	countHabitStorage CountHabitStorage,
	countTaskStorage CountTaskStorage,
	countChallengeStorage CountChallengeStorage) *statisticalBiz {
	return &statisticalBiz{
		countUserStorage:      countUserStorage,
		countHabitStorage:     countHabitStorage,
		countTaskStorage:      countTaskStorage,
		countChallengeStorage: countChallengeStorage,
	}
}

func (b *statisticalBiz) GetStatistical(time string) (*statisticalmodel.Statistical, error) {
	// get main statistical
	cc, err := b.countChallengeStorage.CountChallengesByTimeCreated("")
	if err != nil {
		return nil, err
	}

	fmt.Println(cc)

	uc, err := b.countUserStorage.CountUserByTimeCreated("")
	if err != nil {
		return nil, err
	}

	hc, err := b.countHabitStorage.CountHabitByTimeCreated("")
	if err != nil {
		return nil, err
	}

	tc, err := b.countTaskStorage.CountTaskByTimeCreated("")
	if err != nil {
		return nil, err
	}

	// get element statistical
	_cc, err := b.countChallengeStorage.CountChallengesByTimeCreated(time)
	if err != nil {
		return nil, err
	}

	_uc, err := b.countUserStorage.CountUserByTimeCreated(time)
	if err != nil {
		return nil, err
	}

	_hc, err := b.countHabitStorage.CountHabitByTimeCreated(time)
	if err != nil {
		return nil, err
	}

	_tc, err := b.countTaskStorage.CountTaskByTimeCreated(time)
	if err != nil {
		return nil, err
	}

	return statisticalmodel.NewStatistical(tc, hc, uc, cc, time,
		statisticalmodel.NewStatisticalElement(_tc, _hc, _uc, _cc)), nil
}
