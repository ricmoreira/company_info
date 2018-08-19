// Inspiration to create dependcy injection came from this post: https://blog.drewolson.org/dependency-injection-in-go/

package containers

import (
	"company_info/config"
	controllers "company_info/controllers/v1"
	"company_info/handlers"
	"company_info/repositories"
	"company_info/server"
	"company_info/services"

	"go.uber.org/dig"
)

// BuildContainer returns a container with all app dependencies built in
func BuildContainer() *dig.Container {
	container := dig.New()

	// config
	err := container.Provide(config.NewConfig)
	if err != nil {
		panic(err)
	}

	// persistance layer
	err = container.Provide(repositories.NewDBCollections)
	if err != nil {
		panic(err)
	}
	err = container.Provide(repositories.NewHeaderRepository)
	if err != nil {
		panic(err)
	}

	// services
	err = container.Provide(services.NewHeaderService)
	if err != nil {
		panic(err)
	}
	err = container.Provide(services.NewKafkaConsumer)
	if err != nil {
		panic(err)
	}

	// controllers
	err = container.Provide(controllers.NewHeaderController)
	if err != nil {
		panic(err)
	}

	// generic http layer
	err = container.Provide(handlers.NewHttpHandlers)
	if err != nil {
		panic(err)
	}

	// server
	err = container.Provide(server.NewServer)
	if err != nil {
		panic(err)
	}

	return container
}
