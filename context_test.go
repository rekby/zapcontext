package zc

import (
	"context"
	"go.uber.org/zap"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("L", func(t *testing.T) {
		t.Parallel()

		if L(context.Background()) != nil {
			t.Error()
		}

		l := zap.NewNop()
		ctx := WithContext(context.Background(), l)
		l2 := L(ctx)

		if l != l2 {
			t.Error()
		}

		ctxSugar := WithContextSugar(context.Background(), zap.NewNop().Sugar())
		if L(ctxSugar) == nil {
			t.Error()
		}
	})

	t.Run("LNop", func(t *testing.T) {
		t.Parallel()

		if LNop(context.Background()) == nil {
			t.Error()
		}

		l := zap.NewNop()
		ctx := WithContext(context.Background(), l)
		l2 := LNop(ctx)

		if l != l2 {
			t.Error()
		}

		ctxSugar := WithContextSugar(context.Background(), zap.NewNop().Sugar())
		if LNop(ctxSugar) == nil {
			t.Error()
		}
	})

	t.Run("S", func(t *testing.T) {
		t.Parallel()

		if S(context.Background()) != nil {
			t.Error()
		}

		l := zap.NewNop().Sugar()
		ctx := WithContextSugar(context.Background(), l)
		l2 := S(ctx)

		if l != l2 {
			t.Error()
		}

		ctxSugar := WithContext(context.Background(), zap.NewNop())
		if S(ctxSugar) == nil {
			t.Error()
		}
	})

	t.Run("SNop", func(t *testing.T) {
		t.Parallel()

		if SNop(context.Background()) == nil {
			t.Error()
		}

		l := zap.NewNop().Sugar()
		ctx := WithContextSugar(context.Background(), l)
		l2 := SNop(ctx)

		if l != l2 {
			t.Error()
		}

		ctxSugar := WithContext(context.Background(), zap.NewNop())
		if SNop(ctxSugar) == nil {
			t.Error()
		}
	})

}
