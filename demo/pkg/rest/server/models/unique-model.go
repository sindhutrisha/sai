package models

type Unique struct {
	Id int64 `json:"id,omitempty"`

	Age int8 `json:"age,omitempty"`

	Star string `json:"star,omitempty"`

	Valid bool `json:"valid,omitempty"`
}
