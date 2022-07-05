package vel

import "github.com/Thomasparsley/vel/modules/identity"

func requestDataParser[T any](parser func(out any) error) (T, error) {
	var data T

	err := parser(data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Parse query and store to locals
func Query[Q any]() Handler {
	return func(ctx *Ctx) error {
		query, err := requestDataParser[Q](ctx.QueryParser)
		if err != nil {
			return err
		}

		ctx.Locals("query", query)
		return ctx.Next()
	}
}

// Parse body and store to locals
func Body[B any]() Handler {
	return func(ctx *Ctx) error {
		body, err := requestDataParser[B](ctx.BodyParser)
		if err != nil {
			return err
		}

		ctx.Locals("body", body)
		return ctx.Next()
	}
}

// Parse params and store to locals
func Params[P any]() Handler {
	return func(ctx *Ctx) error {
		params, err := requestDataParser[P](ctx.ParamsParser)
		if err != nil {
			return err
		}

		ctx.Locals("params", params)
		return ctx.Next()
	}
}

func Redirect(location string, status ...int) Handler {
	return func(ctx *Ctx) error {
		err := ctx.Next()

		if err != nil {
			return err
		}

		return ctx.Redirect(location, status...)
	}
}

func AuthenticationRequired() Handler {
	return func(ctx *Ctx) error {
		user := ctx.GetLocalUser()
		if user == nil {
			return nil // TODO: return error
		}

		return ctx.Next()
	}
}

func AdminRequired() Handler {
	return func(ctx *Ctx) error {
		err := AuthenticationRequired()(ctx)
		if err != nil {
			return err
		}

		user := ctx.GetLocalUser()
		if !user.IsAdmin() {
			return nil // TODO: Return error
		}

		return ctx.Next()
	}
}

func RoleRequired(name identity.RoleName) Handler {
	return func(ctx *Ctx) error {
		err := AuthenticationRequired()(ctx)
		if err != nil {
			return err
		}

		user := ctx.GetLocalUser()
		if !user.HasRole(name) {
			return nil // TODO: Return error
		}

		return ctx.Next()
	}
}

func PermissionRequired(name identity.PermissionName, permissions identity.Permissions) Handler {
	return func(ctx *Ctx) error {
		err := AuthenticationRequired()(ctx)
		if err != nil {
			return err
		}

		user := ctx.GetLocalUser()
		if !user.HasPermission(name, permissions) {
			return nil // TODO: Return error
		}

		return ctx.Next()
	}
}
