// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/interface/v1/comment.proto

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

	if v, ok := interface{}(m.GetUserId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommentAddReqValidationError{
				field:  "UserId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFoodId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommentAddReqValidationError{
				field:  "FoodId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetComment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommentAddReqValidationError{
				field:  "Comment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

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

	// no validation rules for CommentId

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

// Validate checks the field values on ListCommentResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListCommentResp) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCommentRespValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListCommentRespValidationError is the validation error returned by
// ListCommentResp.Validate if the designated constraints aren't met.
type ListCommentRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCommentRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCommentRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCommentRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCommentRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCommentRespValidationError) ErrorName() string { return "ListCommentRespValidationError" }

// Error satisfies the builtin error interface
func (e ListCommentRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCommentResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCommentRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCommentRespValidationError{}

// Validate checks the field values on ListCommentResp_ListCommentItem with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCommentResp_ListCommentItem) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	// no validation rules for UserName

	// no validation rules for FoodId

	// no validation rules for FoodName

	// no validation rules for Comment

	return nil
}

// ListCommentResp_ListCommentItemValidationError is the validation error
// returned by ListCommentResp_ListCommentItem.Validate if the designated
// constraints aren't met.
type ListCommentResp_ListCommentItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCommentResp_ListCommentItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCommentResp_ListCommentItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCommentResp_ListCommentItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCommentResp_ListCommentItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCommentResp_ListCommentItemValidationError) ErrorName() string {
	return "ListCommentResp_ListCommentItemValidationError"
}

// Error satisfies the builtin error interface
func (e ListCommentResp_ListCommentItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCommentResp_ListCommentItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCommentResp_ListCommentItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCommentResp_ListCommentItemValidationError{}
