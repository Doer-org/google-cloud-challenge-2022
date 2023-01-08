// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddUser implements addUser operation.
	//
	// PATCH /organizations/{id}/addUser
	AddUser(ctx context.Context, params AddUserParams) error
	// CreateComment implements createComment operation.
	//
	// Creates a new Comment and persists it to storage.
	//
	// POST /comments
	CreateComment(ctx context.Context, req *CreateCommentReq) (CreateCommentRes, error)
	// CreateEvent implements createEvent operation.
	//
	// Creates a new Event and persists it to storage.
	//
	// POST /events
	CreateEvent(ctx context.Context, req *CreateEventReq) (CreateEventRes, error)
	// CreateUser implements createUser operation.
	//
	// Creates a new User and persists it to storage.
	//
	// POST /users
	CreateUser(ctx context.Context, req *CreateUserReq) (CreateUserRes, error)
	// DeleteComment implements deleteComment operation.
	//
	// Deletes the Comment with the requested ID.
	//
	// DELETE /comments/{id}
	DeleteComment(ctx context.Context, params DeleteCommentParams) (DeleteCommentRes, error)
	// DeleteEvent implements deleteEvent operation.
	//
	// Deletes the Event with the requested ID.
	//
	// DELETE /events/{id}
	DeleteEvent(ctx context.Context, params DeleteEventParams) (DeleteEventRes, error)
	// DeleteUser implements deleteUser operation.
	//
	// Deletes the User with the requested ID.
	//
	// DELETE /users/{id}
	DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error)
	// ListComment implements listComment operation.
	//
	// List Comments.
	//
	// GET /comments
	ListComment(ctx context.Context, params ListCommentParams) (ListCommentRes, error)
	// ListEvent implements listEvent operation.
	//
	// List Events.
	//
	// GET /events
	ListEvent(ctx context.Context, params ListEventParams) (ListEventRes, error)
	// ListEventUsers implements listEventUsers operation.
	//
	// List attached Users.
	//
	// GET /events/{id}/users
	ListEventUsers(ctx context.Context, params ListEventUsersParams) (ListEventUsersRes, error)
	// ListUser implements listUser operation.
	//
	// List Users.
	//
	// GET /users
	ListUser(ctx context.Context, params ListUserParams) (ListUserRes, error)
	// ListUserEvents implements listUserEvents operation.
	//
	// List attached Events.
	//
	// GET /users/{id}/events
	ListUserEvents(ctx context.Context, params ListUserEventsParams) (ListUserEventsRes, error)
	// ReadComment implements readComment operation.
	//
	// Finds the Comment with the requested ID and returns it.
	//
	// GET /comments/{id}
	ReadComment(ctx context.Context, params ReadCommentParams) (ReadCommentRes, error)
	// ReadCommentEvent implements readCommentEvent operation.
	//
	// Find the attached Event of the Comment with the given ID.
	//
	// GET /comments/{id}/event
	ReadCommentEvent(ctx context.Context, params ReadCommentEventParams) (ReadCommentEventRes, error)
	// ReadCommentUser implements readCommentUser operation.
	//
	// Find the attached User of the Comment with the given ID.
	//
	// GET /comments/{id}/user
	ReadCommentUser(ctx context.Context, params ReadCommentUserParams) (ReadCommentUserRes, error)
	// ReadEvent implements readEvent operation.
	//
	// Finds the Event with the requested ID and returns it.
	//
	// GET /events/{id}
	ReadEvent(ctx context.Context, params ReadEventParams) (ReadEventRes, error)
	// ReadEventAdmin implements readEventAdmin operation.
	//
	// Find the attached User of the Event with the given ID.
	//
	// GET /events/{id}/admin
	ReadEventAdmin(ctx context.Context, params ReadEventAdminParams) (ReadEventAdminRes, error)
	// ReadUser implements readUser operation.
	//
	// Finds the User with the requested ID and returns it.
	//
	// GET /users/{id}
	ReadUser(ctx context.Context, params ReadUserParams) (ReadUserRes, error)
	// UpdateComment implements updateComment operation.
	//
	// Updates a Comment and persists changes to storage.
	//
	// PATCH /comments/{id}
	UpdateComment(ctx context.Context, req *UpdateCommentReq, params UpdateCommentParams) (UpdateCommentRes, error)
	// UpdateEvent implements updateEvent operation.
	//
	// Updates a Event and persists changes to storage.
	//
	// PATCH /events/{id}
	UpdateEvent(ctx context.Context, req *UpdateEventReq, params UpdateEventParams) (UpdateEventRes, error)
	// UpdateUser implements updateUser operation.
	//
	// Updates a User and persists changes to storage.
	//
	// PATCH /users/{id}
	UpdateUser(ctx context.Context, req *UpdateUserReq, params UpdateUserParams) (UpdateUserRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
