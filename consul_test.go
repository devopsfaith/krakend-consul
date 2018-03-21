package consul

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/devopsfaith/krakend/config"
)

func TestRegister_ok(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	extraCfg := config.ExtraConfig{
		Namespace: map[string]interface{}{
			"address": "127.0.0.1:8500",
			"name":    "test",
			"tags": []interface{}{
				"1224", "2233",
			},
		},
	}

	err := Register(ctx, extraCfg, 111)
	if err != nil {
		t.Errorf("error %s", err.Error())
		return
	}

	fmt.Println("registered")

	<-time.After(30 * time.Second)
	cancel()
	fmt.Println("deregistered")

}
