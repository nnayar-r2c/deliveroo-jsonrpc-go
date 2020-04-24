// Package jsonrpc implements a microframework for writing JSON-RPC web
// applications.
//
// Methods are defined as either:
//     func(ctx context.Context, params T) (interface{}, error) // params can be any JSON unmarshalable type
//     func(ctx context.Context) (interface{}, error)           // or, params can be omitted.
//
// If a method returns a value along with a nil error, the value will be
// rendered to the client as JSON.
//
// If an error is returned, it will be sanitized and returned to the client as
// json. Errors generated by a call to `jsonrpc.Error` will be rendered as-is to
// the client. Any other errors will be obfuscated to the caller (unless
// `DumpErrors` is enabled).
//
// Example:
//
//	var logger = log.New(os.Stderr, "server: ", 0)
//
//	func main() {
//		server := jsonrpc.New()
//		server.Use(LoggingMiddleware(logger))
//		server.Register(jsonrpc.Methods{
//			"Hello": hello,
//		})
//
//		http.ListenAndServe(":80", server)
//	}
//
//	type helloParams struct {
//		Name string `json:"name"`
//	}
//
//	func hello(ctx context.Context, params *helloParams) (interface{}, error) {
//		return jsonrpc.M{"message": fmt.Sprintf("Hello, %s", params.Name)}, nil
//	}
//
//	func LoggingMiddleware(logger *logger.Logger) Middleware {
//		return func (next jsonrpc.Next) jsonrpc.Next {
//			return func(ctx context.Context, params interface{}) (interface{}, error) {
//				method := jsonrpc.MethodFromContext(ctx)
//				start := time.Now()
//				defer func() {
//					logger.Printf("%s (%v)\n", method, time.Since(start))
//				}()
//				return next(ctx, params)
//			}
//		}
//	}
package jsonrpc
