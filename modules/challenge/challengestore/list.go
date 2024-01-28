package challengestore

import (
	"bestHabit/common"
	"bestHabit/modules/challenge/challengemodel"
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

func (s *sqlStore) ListChallengesByConditions(ctx context.Context,
	filter *challengemodel.ChallengeFilter,
	paging *common.Paging,
	conditions map[string]interface{}) ([]challengemodel.Challenge, error) {
	db := s.db

	args := []interface{}{}
	query := "SELECT * FROM challenges"
	countQuery := "SELECT COUNT(*) FROM challenges"

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

	var challenges []challengemodel.Challenge
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
	if err := db.Select(&challenges, query, args...); err != nil {
		return nil, common.ErrDB(err)
	}

	return challenges, nil
}
