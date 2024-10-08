// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DeletePhoto implements deletePhoto operation.
	//
	// Delete a photo from the server.
	//
	// DELETE /v1/photos/{id}
	DeletePhoto(ctx context.Context, params DeletePhotoParams) (DeletePhotoRes, error)
	// GetPhoto implements getPhoto operation.
	//
	// Get a photo from the server by ID.
	//
	// GET /v1/photos/{id}
	GetPhoto(ctx context.Context, params GetPhotoParams) (GetPhotoRes, error)
	// GetPhotos implements getPhotos operation.
	//
	// Get photos from the server with pagination.
	//
	// GET /v1/photos
	GetPhotos(ctx context.Context, params GetPhotosParams) (GetPhotosRes, error)
	// PostPhoto implements postPhoto operation.
	//
	// Upload a new photo to the server.
	//
	// POST /v1/photos
	PostPhoto(ctx context.Context, req *PostPhotoReq) (PostPhotoRes, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
