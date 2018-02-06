package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"microservices/restapi/operations"
	"microservices/restapi/operations/comment"
	"microservices/restapi/operations/like"
	"microservices/internal/comment_like"
	"microservices/models"
	"fmt"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../api/comment_like/comment_like.yaml

func configureFlags(api *operations.CommentLikeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CommentLikeAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.LikeCountLikeHandler = like.CountLikeHandlerFunc(func(params like.CountLikeParams) middleware.Responder {
		count, err := comment_like.CountLike(params.Title)
		if err != nil {
			return like.NewCountLikeBadRequest()
		}
		return like.NewCountLikeOK().WithPayload(&models.CountLikeModel{
			Count: &count,
			Title: &params.Title,
		})
	})
	api.CommentDoCommentHandler = comment.DoCommentHandlerFunc(func(params comment.DoCommentParams) middleware.Responder {
		//comment_like.DoComment(params.HTTPRequest.p)
		ip := params.HTTPRequest.Header.Get("RemoteAddr")
		fmt.Println(ip)
		fmt.Printf("%#v\n", params.HTTPRequest)
		fmt.Printf("%+v\n", params.HTTPRequest)
		return middleware.NotImplemented("operation comment.DoComment has not yet been implemented")
	})
	api.LikeDoLikeHandler = like.DoLikeHandlerFunc(func(params like.DoLikeParams) middleware.Responder {
		return middleware.NotImplemented("operation like.DoLike has not yet been implemented")
	})
	api.LikeDoUnlikeHandler = like.DoUnlikeHandlerFunc(func(params like.DoUnlikeParams) middleware.Responder {
		return middleware.NotImplemented("operation like.DoUnlike has not yet been implemented")
	})
	api.CommentShowCommentHandler = comment.ShowCommentHandlerFunc(func(params comment.ShowCommentParams) middleware.Responder {
		return middleware.NotImplemented("operation comment.ShowComment has not yet been implemented")
	})
	api.LikeShowLikeHandler = like.ShowLikeHandlerFunc(func(params like.ShowLikeParams) middleware.Responder {
		return middleware.NotImplemented("operation like.ShowLike has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
