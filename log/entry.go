package log

import (
	"context"
	"time"

	"github.com/goravel/framework/contracts/log"
)

type Entry struct {
	code       string
	ctx        context.Context
	data       log.Data
	domain     string
	hint       string
	level      log.Level
	message    string
	owner      any
	request    map[string]any
	response   map[string]any
	stacktrace map[string]any
	tags       []string
	time       time.Time
	user       any
	with       map[string]any
}

func (r *Entry) Code() string {
	return r.code
}

func (r *Entry) Context() context.Context {
	return r.ctx
}

func (r *Entry) Data() log.Data {
	return r.data
}

func (r *Entry) Domain() string {
	return r.domain
}

func (r *Entry) Hint() string {
	return r.hint
}

func (r *Entry) Level() log.Level {
	return r.level
}

func (r *Entry) Message() string {
	return r.message
}

func (r *Entry) Owner() any {
	return r.owner
}

func (r *Entry) Request() map[string]any {
	return r.request
}

func (r *Entry) Response() map[string]any {
	return r.response
}

func (r *Entry) Tags() []string {
	return r.tags
}

func (r *Entry) Time() time.Time {
	return r.time
}

func (r *Entry) Trace() map[string]any {
	return r.stacktrace
}

func (r *Entry) User() any {
	return r.user
}

func (r *Entry) With() map[string]any {
	return r.with
}
