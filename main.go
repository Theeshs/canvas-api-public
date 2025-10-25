package main

import (
	"context"
	"log"

	"api/config"
	"api/ent"
	"api/services/educations"
	"api/services/experiences"
	"api/services/projects"
	"api/services/public"
	"api/services/skills"
	"api/services/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	// Load config
	config.MustLoad()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Connect to DB
	client, err := ent.Open(dialect.Postgres, config.AppConfig.BuildDSN())
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(
		context.Background(),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	registerRoutes(app, client)
	log.Fatal(app.Listen(":" + config.AppConfig.AppPort))
}

func registerRoutes(app *fiber.App, client *ent.Client) {
	user.RegisterRoutes(app, client)
	educations.RegisterRoutes(app, client)
	experiences.RegisterRoutes(app, client)
	skills.RegisterRoutes(app, client)
	public.RegisterRoutes(app, client)
	projects.RegisterRoutes(app, client)
}
