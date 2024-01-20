package habitstorage

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"context"
	"fmt"
	"strings"
)

func replacePlaceholders(query string, args []interface{}) string {
	for _, arg := range args {
		strVal := fmt.Sprintf("'%v'", arg)
		query = strings.Replace(query, "?", strVal, 1)
	}

	return query
}

func (s *sqlStore) ListHabitByConditions(ctx context.Context,
	filter *habitmodel.HabitFilter,
	paging *common.Paging,
	conditions map[string]interface{}) ([]habitmodel.Habit, error) {
	db := s.db

	args := []interface{}{}
	query := "SELECT * FROM habits"
	countQuery := "SELECT COUNT(*) FROM habits"

	var conditionsAndMore string

	// add conditions
	// user_id = ?? or deadline = ??
	check := false
	if len(conditions) > 0 {
		conditionsAndMore += " WHERE "

		for key, value := range conditions {
			if check {
				conditionsAndMore += " AND "
			}
			conditionsAndMore += key + " = ? "
			check = true

			args = append(args, value)
		}
	}

	// add filter conditions
	if v := filter.Name; v != "" {
		if len(conditions) > 0 {
			conditionsAndMore += " AND "
		} else {
			conditionsAndMore += " WHERE "
		}
		conditionsAndMore += "name LIKE " + "'%" + v + "%'"
	}

	// add status
	if len(conditions) > 0 || filter.Name != "" {
		conditionsAndMore += " AND status in (1)"
	} else {
		conditionsAndMore += " WHERE status in (1)"
	}

	var habits []habitmodel.Habit
	limit := paging.Limit

	// count paging
	var total int64
	countQuery = db.Rebind(countQuery + conditionsAndMore)
	countQuery = replacePlaceholders(countQuery, args)

	if err := db.Get(&total, countQuery); err != nil {
		return nil, common.ErrDB(err)
	}

	paging.Total = total

	// update paging
	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			conditionsAndMore = conditionsAndMore + fmt.Sprintf(" AND id < %d ", int(uid.GetLocalID())) + "ORDER BY id DESC LIMIT ?"
			args = append(args, limit)
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit

		conditionsAndMore = conditionsAndMore + " ORDER BY id DESC LIMIT ? OFFSET ?"
		args = append(args, limit, offset)
	}

	query = db.Rebind(query + conditionsAndMore)
	if err := db.Select(&habits, query, args...); err != nil {
		return nil, common.ErrDB(err)
	}

	return habits, nil
}
