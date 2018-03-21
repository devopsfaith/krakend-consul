package consul

import (
	"context"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/sd/dnssrv"
	"github.com/hashicorp/consul/api"
)

func Register(ctx context.Context, e config.ExtraConfig, port int) error {
	cfg, err := parse(e, port)
	if err != nil {
		return err
	}
	return register(ctx, cfg)
}

func register(ctx context.Context, cfg Config) error {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = cfg.Address
	c, err := api.NewClient(consulConfig)
	if err != nil {
		return err
	}

	service := &api.AgentServiceRegistration{
		//ID:   cfg.Name + time.Now().String(),
		Name: cfg.Name,
		Port: cfg.Port,
		Tags: cfg.Tags,
	}

	if err := c.Agent().ServiceRegister(service); err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		c.Agent().ServiceDeregister(service.ID)
	}()

	return dnssrv.Register()
}
