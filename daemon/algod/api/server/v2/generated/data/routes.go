// Package data provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package data

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Removes minimum sync round restriction from the ledger.
	// (DELETE /v2/ledger/sync)
	UnsetSyncRound(ctx echo.Context) error
	// Returns the minimum sync round the ledger is keeping in cache.
	// (GET /v2/ledger/sync)
	GetSyncRound(ctx echo.Context) error
	// Given a round, tells the ledger to keep that round in its cache.
	// (POST /v2/ledger/sync/{round})
	SetSyncRound(ctx echo.Context, round uint64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// UnsetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) UnsetSyncRound(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UnsetSyncRound(ctx)
	return err
}

// GetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) GetSyncRound(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSyncRound(ctx)
	return err
}

// SetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) SetSyncRound(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameterWithLocation("simple", false, "round", runtime.ParamLocationPath, ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SetSyncRound(ctx, round)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/v2/ledger/sync", wrapper.UnsetSyncRound, m...)
	router.GET(baseURL+"/v2/ledger/sync", wrapper.GetSyncRound, m...)
	router.POST(baseURL+"/v2/ledger/sync/:round", wrapper.SetSyncRound, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAAA////ElqSu/oAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
