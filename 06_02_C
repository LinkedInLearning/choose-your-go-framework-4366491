import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
)

func main() {
	// Initialize the logger middleware
	logger := middleware.Chain(
		middleware.Recovery(),
		middleware.Trace(),
		middleware.Log(log.DefaultLogger),
	)

	// Create a service with the logger middleware
	app := kratos.New(
		kratos.Name("my-service"),
		kratos.Server(
			http.NewServer(":8000", logger),
		),
	)
	app.Run()
}
