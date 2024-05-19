package repository

import (
	"errors"
	"fmt"
	"time"

	"database/sql"

	"github.com/lionelstefan/go-dating-app/model"
	"github.com/lionelstefan/go-dating-app/shared"
	"github.com/google/uuid"
)

func AddNewSwipe(req model.Swipes) error {
	db := shared.CreateDBConnection()

	canSwipe, err := CanSwipe(req.UserId)

	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	if !canSwipe {
		return errors.New("limit swipe")
	}


	t := time.Now()
	formattedDate := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	var swipeCount int
	checkSwipeCount := db.Get(&swipeCount, "SELECT count FROM swipes_count WHERE user_id = $1 AND DATE(created_at) = $2", req.UserId, formattedDate)

	if checkSwipeCount != nil {
		if checkSwipeCount == sql.ErrNoRows {
			swipeCount = 1
		}
	} else {
		swipeCount += 1
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO swipes (user_id, target_id, is_like) VALUES ($1, $2, $3)", req.UserId, req.TargetId, req.IsLike)
	if checkSwipeCount != nil {
		if checkSwipeCount == sql.ErrNoRows {
			tx.MustExec("INSERT INTO swipes_count (user_id, count) VALUES ($1, $2)", req.UserId, swipeCount)
		}
	} else {
		tx.MustExec("UPDATE swipes_count SET count = $1 WHERE user_id = $2 AND DATE(created_at) = $3", swipeCount, req.UserId, formattedDate)
	}
	tx.Commit()

	return nil
}

func CanSwipe(user_id uuid.UUID) (bool, error) {
	db := shared.CreateDBConnection()

	var swipeCount int

	t := time.Now()
	formattedDate := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())

	err := db.Get(&swipeCount, "SELECT count FROM swipes_count WHERE user_id = $1 AND DATE(created_at) = $2", user_id, formattedDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		} else {
			return false, err
		}
	}

	if swipeCount >= 10 {
		return false, nil
	}

	return true, nil
}
