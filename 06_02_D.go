import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func main() {
	// Initialize the file configuration source
	source := file.NewSource("config.yaml")

	// Create a configuration object
	cfg, err := config.New(
		config.WithSource(source),
		config.WithDecoder(file.NewDecoder()),
	)
	if err != nil {
		// Handle error
	}

	// Get a configuration value
	port := cfg.Value("server.port").Int(8080)
}
