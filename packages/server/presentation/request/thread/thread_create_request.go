package request

type ThreadCreateRequest struct {
	Title             string  `json:"title" binding:"required,max=50"`
	Description       *string `json:"description" binding:"omitempty,max=255"`
	ThumbnailUrl      *string `json:"thumbnailUrl" binding:"omitempty,max=255"`
	IsNotifyOnComment bool    `json:"isNotifyOnComment"`
}
