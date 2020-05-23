package main

import (
	"go.uber.org/zap"
)

func main() {
	sugar := zap.NewExample()
	defer sugar.Sync()
	sugar.Sugar().Infow("deneme", zap.Object("test", "valu"))
}
