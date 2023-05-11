import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/registry/etcd"
)

func main() {
	// Initialize the etcd registry
	r := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"localhost:2379"}
	})

	// Register a service
	err := r.Register(context.Background(), &registry.ServiceInstance{
		ID:       "service-1",
		Name:     "service",
		Version:  "v1.0.0",
		Metadata: map[string]string{"framework": "kratos"},
		Endpoints: []*registry.Endpoint{
			{
				Zone:     "",
				Protocol: "grpc",
				Address:  "127.0.0.1:9000",
			},
		},
	})
	if err != nil {
		// Handle error
	}

	// Discover a service
	instances, err := r.Fetch(context.Background(), "service")
	if err != nil {
		// Handle error
	}
	// Use the instances to communicate with the service
}

