package domain

type Configs []Config
type Config struct {
	ID   uint64 `json:"id" validate:"required" type:"autoIncrement"`
	Name string `json:"name" validate:"required"`
}
