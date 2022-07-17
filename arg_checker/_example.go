package example

import (
	"context"
	ctxPkg "context"
)

func OKFunc1(ctx context.Context, id int, email string) {}

func OKFunc2(ctx ctxPkg.Context, attrs map[string]interface{}) {}

func OKFunc3(id int, email string, ctx context.Context) {}

func InvalidFunc(c context.Context) {}
