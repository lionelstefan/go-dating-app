package repository

import (
	"github.com/lionelstefan/go-dating-app/model"
	"github.com/lionelstefan/go-dating-app/shared"
)

func Purchase(req model.PremiumPurchase) error {
	db := shared.CreateDBConnection()

	db.MustExec("INSERT INTO premium_purchase (user_id, premium_type) VALUES ($1, $2)", req.UserId, req.PremiumType)

	if req.PremiumType == "no_swipe_limit" {
		db.MustExec("UPDATE users SET no_swipe_limit = true WHERE id = $1", req.UserId)
	}

	if req.PremiumType == "verified" {
		db.MustExec("UPDATE users SET is_verified = true WHERE id = $1", req.UserId)
	}

	return nil
}