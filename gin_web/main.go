package main

import (
	"log"

	"go-web-lab/gin_web/handlers"
	"go-web-lab/gin_web/middleware"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())

	r.LoadHTMLGlob("gin_web/templates/*")
	r.Static("/static", "./gin_web/static")
	r.StaticFile("/favicon.ico", "./gin_web/static/favicon.ico")

	r.GET("/", handlers.Home)

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", handlers.GetUsers)
		userRoutes.GET("/:id", handlers.GetUserByID)
		userRoutes.POST("", handlers.CreateUser)
		userRoutes.PUT("/:id", handlers.UpdateUser)
		userRoutes.DELETE("/:id", handlers.DeleteUser)
	}

	r.GET("/search", handlers.Search)

	api := r.Group("/api")
	{
		api.GET("/products", handlers.GetProducts)
		api.POST("/products", handlers.CreateProduct)
	}

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/data", handlers.ProtectedData)
		protected.GET("/user", handlers.ProtectedWithUser)
	}

	uploadRoutes := r.Group("/upload")
	{
		uploadRoutes.POST("/single", handlers.UploadSingleFile)
		uploadRoutes.POST("/multiple", handlers.UploadMultipleFiles)
		uploadRoutes.POST("/save", handlers.SaveUploadedFile)
	}

	r.POST("/form", handlers.ProcessForm)

	templateRoutes := r.Group("/template")
	{
		templateRoutes.GET("/index", handlers.RenderHTML)
		templateRoutes.GET("/user/:id", handlers.RenderUserTemplate)
		templateRoutes.GET("/layout", handlers.RenderWithLayout)
	}

	cookieRoutes := r.Group("/cookie")
	{
		cookieRoutes.GET("/set", handlers.SetCookie)
		cookieRoutes.GET("/get", handlers.GetCookie)
		cookieRoutes.DELETE("/delete", handlers.DeleteCookie)
		cookieRoutes.GET("/multiple", handlers.MultipleCookies)
	}

	redirectRoutes := r.Group("/redirect")
	{
		redirectRoutes.GET("/external", handlers.RedirectToURL)
		redirectRoutes.GET("/internal", handlers.RedirectInternal)
		redirectRoutes.GET("/temporary", handlers.RedirectWithStatus)
		redirectRoutes.GET("/conditional", handlers.ConditionalRedirect)
	}

	contextRoutes := r.Group("/context")
	{
		contextRoutes.POST("/set", handlers.SetContext)
		contextRoutes.GET("/get", handlers.GetContext)
		contextRoutes.GET("/chain", handlers.ContextChain)
	}

	responseRoutes := r.Group("/response")
	{
		responseRoutes.GET("/xml", handlers.XMLResponse)
		responseRoutes.GET("/yaml", handlers.YAMLResponse)
		responseRoutes.GET("/protobuf", handlers.ProtoBufResponse)
		responseRoutes.GET("/string", handlers.StringResponse)
		responseRoutes.GET("/data", handlers.DataResponse)
		responseRoutes.GET("/header", handlers.HeaderOnlyResponse)
	}

	advancedRoutes := r.Group("/advanced")
	{
		advancedRoutes.GET("/path/:id/:name", handlers.PathParams)
		advancedRoutes.GET("/file/*filepath", handlers.WildcardParams)
		advancedRoutes.GET("/query-array", handlers.QueryArray)
		advancedRoutes.GET("/query-map", handlers.QueryMap)
		advancedRoutes.POST("/post-array", handlers.PostArray)
		advancedRoutes.POST("/post-map", handlers.PostMap)
		advancedRoutes.GET("/client-ip", handlers.ClientIP)
		advancedRoutes.GET("/request-info", handlers.RequestInfo)
		advancedRoutes.GET("/bind-query", handlers.BindQuery)
		advancedRoutes.POST("/bind-form", handlers.BindForm)
		advancedRoutes.GET("/bind-uri/:id/:name", handlers.BindURI)
		advancedRoutes.GET("/bind-header", handlers.BindHeader)
	}

	errorRoutes := r.Group("/error")
	{
		errorRoutes.GET("/trigger", handlers.TriggerError)
		errorRoutes.GET("/all", handlers.AllErrors)
		errorRoutes.GET("/custom", handlers.CustomError)
	}

	r.GET("/async", handlers.AsyncHandler)
	r.GET("/delayed", handlers.DelayedResponse)

	r.NoRoute(handlers.NoRoute)

	return r
}

func main() {
	r := setupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
