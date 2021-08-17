package setup

import (
	"gin-admin/global"
	"gin-admin/pkg/tracer"
)

func setupTracer() error {
	jaegerTrancer, _, err := tracer.NewJaegerTracer("gin-admin", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTrancer
	return nil
}
