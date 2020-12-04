package _map

type AddLink struct {
	Url string `json:"url" form:"url" validate:"required" label:"url"`
}
