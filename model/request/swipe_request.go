package request

type SwipeRequest struct {
	TargetId string `json:targetId`
	IsLike   string `json:isLike`
}