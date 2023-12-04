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
	"0zJzVom2PdDkycmjycf/HwAA///173AkdvcAAA==",
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
