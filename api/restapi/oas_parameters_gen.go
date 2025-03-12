// Code generated by ogen, DO NOT EDIT.

package restapi

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// APITaskIDGetParams is parameters of GET /api/task/{id} operation.
type APITaskIDGetParams struct {
	ID uint64
}

func unpackAPITaskIDGetParams(packed middleware.Parameters) (params APITaskIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uint64)
	}
	return params
}

func decodeAPITaskIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params APITaskIDGetParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUint64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// APITaskIDResolveGetParams is parameters of GET /api/task/{id}/resolve operation.
type APITaskIDResolveGetParams struct {
	ID uint64
}

func unpackAPITaskIDResolveGetParams(packed middleware.Parameters) (params APITaskIDResolveGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uint64)
	}
	return params
}

func decodeAPITaskIDResolveGetParams(args [1]string, argsEscaped bool, r *http.Request) (params APITaskIDResolveGetParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUint64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// APITaskIDSolveGetParams is parameters of GET /api/task/{id}/solve operation.
type APITaskIDSolveGetParams struct {
	ID         uint64
	SolutionID uint64
}

func unpackAPITaskIDSolveGetParams(packed middleware.Parameters) (params APITaskIDSolveGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uint64)
	}
	{
		key := middleware.ParameterKey{
			Name: "solutionID",
			In:   "query",
		}
		params.SolutionID = packed[key].(uint64)
	}
	return params
}

func decodeAPITaskIDSolveGetParams(args [1]string, argsEscaped bool, r *http.Request) (params APITaskIDSolveGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUint64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: solutionID.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "solutionID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUint64(val)
				if err != nil {
					return err
				}

				params.SolutionID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return err
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "solutionID",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
