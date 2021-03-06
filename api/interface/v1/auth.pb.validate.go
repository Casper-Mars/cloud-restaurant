// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/interface/v1/auth.proto

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

// Validate checks the field values on RegistryReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RegistryReq) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPhone()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryReqValidationError{
				field:  "Phone",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPassword()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryReqValidationError{
				field:  "Password",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RegistryReqValidationError is the validation error returned by
// RegistryReq.Validate if the designated constraints aren't met.
type RegistryReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistryReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistryReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistryReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistryReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistryReqValidationError) ErrorName() string { return "RegistryReqValidationError" }

// Error satisfies the builtin error interface
func (e RegistryReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistryReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistryReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistryReqValidationError{}

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LoginReq) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPhone()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginReqValidationError{
				field:  "Phone",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPassword()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginReqValidationError{
				field:  "Password",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

// Validate checks the field values on LoginResp with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LoginResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// LoginRespValidationError is the validation error returned by
// LoginResp.Validate if the designated constraints aren't met.
type LoginRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRespValidationError) ErrorName() string { return "LoginRespValidationError" }

// Error satisfies the builtin error interface
func (e LoginRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRespValidationError{}

// Validate checks the field values on SelfInfoResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SelfInfoResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Phone

	// no validation rules for Name

	// no validation rules for Avatar

	return nil
}

// SelfInfoRespValidationError is the validation error returned by
// SelfInfoResp.Validate if the designated constraints aren't met.
type SelfInfoRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SelfInfoRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SelfInfoRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SelfInfoRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SelfInfoRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SelfInfoRespValidationError) ErrorName() string { return "SelfInfoRespValidationError" }

// Error satisfies the builtin error interface
func (e SelfInfoRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSelfInfoResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SelfInfoRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SelfInfoRespValidationError{}
