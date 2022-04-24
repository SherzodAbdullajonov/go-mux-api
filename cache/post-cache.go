package cache

import (
	"github.com/SherzodAbullajonov/go-mux-api/entity"
)

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
