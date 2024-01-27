package statisticalbiz

import (
	"bestHabit/modules/statistical/statisticalmodel"
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

	cc, err := b.countChallengeStorage.CountChallengesByTimeCreated(time)
	if err != nil {
		return nil, err
	}

	uc, err := b.countUserStorage.CountUserByTimeCreated(time)
	if err != nil {
		return nil, err
	}

	hc, err := b.countHabitStorage.CountHabitByTimeCreated(time)
	if err != nil {
		return nil, err
	}

	tc, err := b.countTaskStorage.CountTaskByTimeCreated(time)
	if err != nil {
		return nil, err
	}

	var statisticalElements []statisticalmodel.StatisticalElement
	templateMonths := []string{"-01", "-02", "-03", "-04", "-05", "-06", "-07",
		"-08", "-09", "-10", "-11", "-12"}
	for i := 0; i < 12; i++ {
		// get element statistical
		_cc, err := b.countChallengeStorage.CountChallengesByTimeCreated(time + templateMonths[i])
		if err != nil {
			return nil, err
		}

		_uc, err := b.countUserStorage.CountUserByTimeCreated(time + templateMonths[i])
		if err != nil {
			return nil, err
		}

		_hc, err := b.countHabitStorage.CountHabitByTimeCreated(time + templateMonths[i])
		if err != nil {
			return nil, err
		}

		_tc, err := b.countTaskStorage.CountTaskByTimeCreated(time + templateMonths[i])
		if err != nil {
			return nil, err
		}

		statisticalElements = append(statisticalElements, *statisticalmodel.NewStatisticalElement(_tc, _hc, _uc, _cc))
	}

	return statisticalmodel.NewStatistical(tc, hc, uc, cc, time, statisticalElements), nil
}
