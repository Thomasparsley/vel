package vel

func requestDataParser[T any](parser func(out any) error) (T, error) {
	var data T

	err := parser(data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Parse query and store to locals
func Query[Q any]() func(ctx *Ctx) error {
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
func Body[B any]() func(ctx *Ctx) error {
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
func Params[P any]() func(ctx *Ctx) error {
	return func(ctx *Ctx) error {
		params, err := requestDataParser[P](ctx.ParamsParser)
		if err != nil {
			return err
		}

		ctx.Locals("params", params)
		return ctx.Next()
	}
}

func Redirect(location string, status ...int) func(ctx *Ctx) error {
	return func(ctx *Ctx) error {
		err := ctx.Next()

		if err != nil {
			return err
		}

		return ctx.Redirect(location, status...)
	}
}
