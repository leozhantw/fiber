package session

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

// go test -run TestStore_getSessionID
func TestStore_getSessionID(t *testing.T) {
	expectedID := "test-session-id"

	// fiber instance
	app := fiber.New()

	t.Run("from cookie", func(t *testing.T) {
		// session store
		store := New()
		// fiber context
		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(ctx)
		// set cookie
		ctx.Request().Header.SetCookie(store.CookieName, expectedID)

		utils.AssertEqual(t, expectedID, store.getSessionID(ctx))
	})

	t.Run("from header", func(t *testing.T) {
		// session store
		store := New(Config{
			Source: SourceHeader,
		})
		// fiber context
		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(ctx)
		// set header
		ctx.Request().Header.Set(store.CookieName, expectedID)

		utils.AssertEqual(t, expectedID, store.getSessionID(ctx))
	})

	t.Run("from url query", func(t *testing.T) {
		// session store
		store := New(Config{
			Source: SourceURLQuery,
		})
		// fiber context
		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(ctx)
		// set url parameter
		ctx.Request().SetRequestURI(fmt.Sprintf("/path?%s=%s", store.CookieName, expectedID))

		utils.AssertEqual(t, expectedID, store.getSessionID(ctx))
	})
}
