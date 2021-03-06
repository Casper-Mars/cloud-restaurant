// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/comment/v1/comment.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on CommentAddReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CommentAddReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	// no validation rules for FoodId

	// no validation rules for Comment

	return nil
}

// CommentAddReqValidationError is the validation error returned by
// CommentAddReq.Validate if the designated constraints aren't met.
type CommentAddReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentAddReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentAddReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentAddReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentAddReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentAddReqValidationError) ErrorName() string { return "CommentAddReqValidationError" }

// Error satisfies the builtin error interface
func (e CommentAddReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentAddReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentAddReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentAddReqValidationError{}

// Validate checks the field values on CommentModifyResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CommentModifyResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CommentModifyRespValidationError is the validation error returned by
// CommentModifyResp.Validate if the designated constraints aren't met.
type CommentModifyRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentModifyRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentModifyRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentModifyRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentModifyRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentModifyRespValidationError) ErrorName() string {
	return "CommentModifyRespValidationError"
}

// Error satisfies the builtin error interface
func (e CommentModifyRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentModifyResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentModifyRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentModifyRespValidationError{}

// Validate checks the field values on CommentList with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CommentList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommentListValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CommentListValidationError is the validation error returned by
// CommentList.Validate if the designated constraints aren't met.
type CommentListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentListValidationError) ErrorName() string { return "CommentListValidationError" }

// Error satisfies the builtin error interface
func (e CommentListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentListValidationError{}

// Validate checks the field values on CommentList_CommentListItem with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CommentList_CommentListItem) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for FoodId

	// no validation rules for Comment

	return nil
}

// CommentList_CommentListItemValidationError is the validation error returned
// by CommentList_CommentListItem.Validate if the designated constraints
// aren't met.
type CommentList_CommentListItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentList_CommentListItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentList_CommentListItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentList_CommentListItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentList_CommentListItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentList_CommentListItemValidationError) ErrorName() string {
	return "CommentList_CommentListItemValidationError"
}

// Error satisfies the builtin error interface
func (e CommentList_CommentListItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentList_CommentListItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentList_CommentListItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentList_CommentListItemValidationError{}
