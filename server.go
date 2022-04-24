package main

import (
	"os"

	"github.com/SherzodAbullajonov/go-mux-api/cache"
	"github.com/SherzodAbullajonov/go-mux-api/controller"
	"github.com/SherzodAbullajonov/go-mux-api/repository"
	router "github.com/SherzodAbullajonov/go-mux-api/router"
	"github.com/SherzodAbullajonov/go-mux-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postController controller.PostController = controller.NewPostController(postService, postCache)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
