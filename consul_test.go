package consul

import (
	"context"
	"fmt"
	"testing"
	"time"

	gologging "github.com/devopsfaith/krakend-gologging"
	"github.com/luraproject/lura/config"
)

func TestRegister_ok(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	extraCfg := config.ExtraConfig{
		Namespace: map[string]interface{}{
			"address": "127.0.0.1:8500",
			"tags": []interface{}{
				"1224", "2233",
			},
		},
	}

	logger, err := gologging.NewLogger(config.ExtraConfig{
		gologging.Namespace: map[string]interface{}{
			"level":  "DEBUG",
			"stdout": true,
		},
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = Register(ctx, extraCfg, 111, "test", logger)
	if err != nil {
		t.Errorf("error %s", err.Error())
		return
	}

	fmt.Println("test registered")

	<-time.After(10 * time.Second)
	cancel()
	<-time.After(10 * time.Second)
	fmt.Println("test deregistered")

}
