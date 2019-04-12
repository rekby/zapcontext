package zc

import (
	"context"
	"testing"

	"go.uber.org/zap"
)

func TestAll(t *testing.T) {

	defaultTestLogger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}
	SetDefaultLogger(defaultTestLogger)

	if defaultSugaredLogger == nil {
		t.Fatal(err)
	}

	t.Run("L", func(t *testing.T) {
		if defaultLogger == nil {
			t.Error()
		}

		if L(nil) != defaultLogger {
			t.Error()
		}

		if L(context.Background()) != defaultLogger {
			t.Error()
		}

		l := zap.NewNop()
		ctx := WithLogger(context.Background(), l)
		l2 := L(ctx)

		if l != l2 {
			t.Error()
		}

		ctxSugar := WithSugarLogger(context.Background(), zap.NewNop().Sugar())
		if L(ctxSugar) == defaultLogger || L(ctxSugar) == nil {
			t.Error()
		}
	})

	t.Run("LNop", func(t *testing.T) {
		if nopLogger == nil {
			t.Error()
		}

		if LNop(nil) != nopLogger {
			t.Error()
		}

		if LNop(context.Background()) != nopLogger {
			t.Error()
		}

		l := zap.NewNop()
		ctx := WithLogger(context.Background(), l)
		l2 := LNop(ctx)

		if l != l2 {
			t.Error()
		}

		ctxSugar := WithSugarLogger(context.Background(), zap.NewNop().Sugar())
		if LNop(ctxSugar) == nopLogger || LNop(ctxSugar) == nil {
			t.Error()
		}
	})

	t.Run("S", func(t *testing.T) {
		if S(nil) != defaultSugaredLogger {
			t.Error()
		}

		if S(context.Background()) != defaultSugaredLogger {
			t.Error()
		}

		l := zap.NewNop().Sugar()
		ctx := WithSugarLogger(context.Background(), l)
		l2 := S(ctx)

		if l != l2 {
			t.Error()
		}

		ctxSugar := WithLogger(context.Background(), zap.NewNop())
		if S(ctxSugar) == nil || S(ctxSugar) == defaultSugaredLogger {
			t.Error()
		}
	})

	t.Run("SNop", func(t *testing.T) {
		if SNop(nil) == nil {
			t.Error()
		}

		if SNop(context.Background()) == nil {
			t.Error()
		}

		l := zap.NewNop().Sugar()
		ctx := WithSugarLogger(context.Background(), l)
		l2 := SNop(ctx)

		if l != l2 {
			t.Error()
		}

		ctxSugar := WithLogger(context.Background(), zap.NewNop())
		if SNop(ctxSugar) == nil || SNop(ctxSugar) == defaultSugaredLogger {
			t.Error()
		}
	})

}
