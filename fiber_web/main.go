package main

import (
	"log"

	"go-web-lab/fiber_web/handlers"
	"go-web-lab/fiber_web/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRouter() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(middleware.LoggerMiddleware())

	app.Static("/static", "./fiber_web/static")

	app.Get("/", handlers.Home)

	userRoutes := app.Group("/users")
	{
		userRoutes.Get("", handlers.GetUsers)
		userRoutes.Get("/:id", handlers.GetUserByID)
		userRoutes.Post("", handlers.CreateUser)
		userRoutes.Put("/:id", handlers.UpdateUser)
		userRoutes.Delete("/:id", handlers.DeleteUser)
	}

	app.Get("/search", handlers.Search)

	api := app.Group("/api")
	{
		api.Get("/products", handlers.GetProducts)
		api.Post("/products", handlers.CreateProduct)
	}

	protected := app.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.Get("/data", handlers.ProtectedData)
		protected.Get("/user", handlers.ProtectedWithUser)
	}

	uploadRoutes := app.Group("/upload")
	{
		uploadRoutes.Post("/single", handlers.UploadSingleFile)
		uploadRoutes.Post("/multiple", handlers.UploadMultipleFiles)
		uploadRoutes.Post("/save", handlers.SaveUploadedFile)
	}

	app.Post("/form", handlers.ProcessForm)

	cookieRoutes := app.Group("/cookie")
	{
		cookieRoutes.Get("/set", handlers.SetCookie)
		cookieRoutes.Get("/get", handlers.GetCookie)
		cookieRoutes.Delete("/delete", handlers.DeleteCookie)
		cookieRoutes.Get("/multiple", handlers.MultipleCookies)
	}

	contextRoutes := app.Group("/context")
	{
		contextRoutes.Post("/set", handlers.SetContext)
		contextRoutes.Get("/get", handlers.GetContext)
		contextRoutes.Get("/chain", handlers.ContextChain)
	}

	responseRoutes := app.Group("/response")
	{
		responseRoutes.Get("/xml", handlers.XMLResponse)
		responseRoutes.Get("/yaml", handlers.YAMLResponse)
		responseRoutes.Get("/string", handlers.StringResponse)
		responseRoutes.Get("/data", handlers.DataResponse)
		responseRoutes.Get("/header", handlers.HeaderOnlyResponse)
	}

	advancedRoutes := app.Group("/advanced")
	{
		advancedRoutes.Get("/path/:id/:name", handlers.PathParams)
		advancedRoutes.Get("/file/*", handlers.WildcardParams)
		advancedRoutes.Get("/query-array", handlers.QueryArray)
		advancedRoutes.Get("/query-map", handlers.QueryMap)
		advancedRoutes.Post("/post-array", handlers.PostArray)
		advancedRoutes.Post("/post-map", handlers.PostMap)
		advancedRoutes.Get("/client-ip", handlers.ClientIP)
		advancedRoutes.Get("/request-info", handlers.RequestInfo)
		advancedRoutes.Get("/bind-query", handlers.BindQuery)
		advancedRoutes.Post("/bind-form", handlers.BindForm)
		advancedRoutes.Get("/bind-uri/:id/:name", handlers.BindURI)
		advancedRoutes.Get("/bind-header", handlers.BindHeader)
	}

	errorRoutes := app.Group("/error")
	{
		errorRoutes.Get("/trigger", handlers.TriggerError)
		errorRoutes.Get("/custom", handlers.CustomError)
	}

	app.Get("/async", handlers.AsyncHandler)
	app.Get("/delayed", handlers.DelayedResponse)

	app.Use(handlers.NoRoute)

	return app
}

func main() {
	app := setupRouter()

	if err := app.Listen(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
