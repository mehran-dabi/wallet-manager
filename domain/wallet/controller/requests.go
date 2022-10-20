package controller

type createRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}

type addFundRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Fund uint64 `json:"fund" binding:"required"`
}

type subtractFundRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Fund uint64 `json:"fund" binding:"required"`
}
