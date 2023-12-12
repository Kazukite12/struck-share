package api

import (
	"fmt"
	"net/http"

	db "github.com/Kazukite12/StruckShare/db/sqlc"
	"github.com/Kazukite12/StruckShare/token"
	"github.com/Kazukite12/StruckShare/util"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*", "Origin", "authorization"},
	})

	handler := c.Handler(server.router)

	return http.ListenAndServe(address, handler)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20

	router.POST("/user/register", server.createUser)
	router.POST("/user/profile", server.profileUpload)
	router.POST("/user/login", server.loginUser)
	router.Static("/uploads/profile", "./assets/uploads/profilePicture/")
	router.GET("/validate", server.validate)

	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))

	authRoutes.POST("/post", server.createPost)
	authRoutes.POST("/post/likes", server.createPostLikes)
	authRoutes.GET("/post/likes", server.getPostLikes)
	authRoutes.GET("/post/:id", server.getPost)
	authRoutes.GET("/posts", server.listPost)
	authRoutes.POST("/comment", server.createComment)
	authRoutes.GET("/comments", server.listComment)

	authRoutes.POST("/follow", server.createFollower)
	authRoutes.GET("/follow/:id", server.listFollower)
	authRoutes.GET("/follow/count/:id", server.countUserTotalFollower)

	server.router = router

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
