// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/plantoremind/v1/cron/cron.proto

package cron

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on CreateCronRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateCronRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCronRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCronRequestMultiError, or nil if none found.
func (m *CreateCronRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCronRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDesc()) < 1 {
		err := CreateCronRequestValidationError{
			field:  "Desc",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetExpression()) < 1 {
		err := CreateCronRequestValidationError{
			field:  "Expression",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateCronRequestMultiError(errors)
	}

	return nil
}

// CreateCronRequestMultiError is an error wrapping multiple validation errors
// returned by CreateCronRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateCronRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCronRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCronRequestMultiError) AllErrors() []error { return m }

// CreateCronRequestValidationError is the validation error returned by
// CreateCronRequest.Validate if the designated constraints aren't met.
type CreateCronRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCronRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCronRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCronRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCronRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCronRequestValidationError) ErrorName() string {
	return "CreateCronRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCronRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCronRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCronRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCronRequestValidationError{}

// Validate checks the field values on UpdateCronRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateCronRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCronRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCronRequestMultiError, or nil if none found.
func (m *UpdateCronRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCronRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateCronRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Desc

	// no validation rules for Expression

	if len(errors) > 0 {
		return UpdateCronRequestMultiError(errors)
	}

	return nil
}

// UpdateCronRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateCronRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateCronRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCronRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCronRequestMultiError) AllErrors() []error { return m }

// UpdateCronRequestValidationError is the validation error returned by
// UpdateCronRequest.Validate if the designated constraints aren't met.
type UpdateCronRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCronRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCronRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCronRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCronRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCronRequestValidationError) ErrorName() string {
	return "UpdateCronRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCronRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCronRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCronRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCronRequestValidationError{}

// Validate checks the field values on DeleteCronRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteCronRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCronRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCronRequestMultiError, or nil if none found.
func (m *DeleteCronRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCronRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteCronRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteCronRequestMultiError(errors)
	}

	return nil
}

// DeleteCronRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteCronRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteCronRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCronRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCronRequestMultiError) AllErrors() []error { return m }

// DeleteCronRequestValidationError is the validation error returned by
// DeleteCronRequest.Validate if the designated constraints aren't met.
type DeleteCronRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCronRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCronRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCronRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCronRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCronRequestValidationError) ErrorName() string {
	return "DeleteCronRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCronRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCronRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCronRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCronRequestValidationError{}

// Validate checks the field values on CronData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CronData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CronData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CronDataMultiError, or nil
// if none found.
func (m *CronData) ValidateAll() error {
	return m.validate(true)
}

func (m *CronData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Desc

	// no validation rules for Expression

	if len(errors) > 0 {
		return CronDataMultiError(errors)
	}

	return nil
}

// CronDataMultiError is an error wrapping multiple validation errors returned
// by CronData.ValidateAll() if the designated constraints aren't met.
type CronDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CronDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CronDataMultiError) AllErrors() []error { return m }

// CronDataValidationError is the validation error returned by
// CronData.Validate if the designated constraints aren't met.
type CronDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CronDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CronDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CronDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CronDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CronDataValidationError) ErrorName() string { return "CronDataValidationError" }

// Error satisfies the builtin error interface
func (e CronDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCronData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CronDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CronDataValidationError{}

// Validate checks the field values on GetCronRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCronRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCronRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCronRequestMultiError,
// or nil if none found.
func (m *GetCronRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCronRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetCronRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetCronRequestMultiError(errors)
	}

	return nil
}

// GetCronRequestMultiError is an error wrapping multiple validation errors
// returned by GetCronRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCronRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCronRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCronRequestMultiError) AllErrors() []error { return m }

// GetCronRequestValidationError is the validation error returned by
// GetCronRequest.Validate if the designated constraints aren't met.
type GetCronRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCronRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCronRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCronRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCronRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCronRequestValidationError) ErrorName() string { return "GetCronRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetCronRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCronRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCronRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCronRequestValidationError{}

// Validate checks the field values on GetCronReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCronReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCronReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCronReplyMultiError, or
// nil if none found.
func (m *GetCronReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCronReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetCronReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetCronReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCronReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetCronReplyMultiError(errors)
	}

	return nil
}

// GetCronReplyMultiError is an error wrapping multiple validation errors
// returned by GetCronReply.ValidateAll() if the designated constraints aren't met.
type GetCronReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCronReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCronReplyMultiError) AllErrors() []error { return m }

// GetCronReplyValidationError is the validation error returned by
// GetCronReply.Validate if the designated constraints aren't met.
type GetCronReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCronReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCronReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCronReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCronReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCronReplyValidationError) ErrorName() string { return "GetCronReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetCronReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCronReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCronReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCronReplyValidationError{}

// Validate checks the field values on ListCronRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCronRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCronRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCronRequestMultiError, or nil if none found.
func (m *ListCronRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCronRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Offset

	// no validation rules for Limit

	// no validation rules for OrderBy

	if len(errors) > 0 {
		return ListCronRequestMultiError(errors)
	}

	return nil
}

// ListCronRequestMultiError is an error wrapping multiple validation errors
// returned by ListCronRequest.ValidateAll() if the designated constraints
// aren't met.
type ListCronRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCronRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCronRequestMultiError) AllErrors() []error { return m }

// ListCronRequestValidationError is the validation error returned by
// ListCronRequest.Validate if the designated constraints aren't met.
type ListCronRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCronRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCronRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCronRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCronRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCronRequestValidationError) ErrorName() string { return "ListCronRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListCronRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCronRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCronRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCronRequestValidationError{}

// Validate checks the field values on ListCronReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListCronReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCronReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListCronReplyMultiError, or
// nil if none found.
func (m *ListCronReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCronReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCronReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCronReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCronReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListCronReplyMultiError(errors)
	}

	return nil
}

// ListCronReplyMultiError is an error wrapping multiple validation errors
// returned by ListCronReply.ValidateAll() if the designated constraints
// aren't met.
type ListCronReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCronReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCronReplyMultiError) AllErrors() []error { return m }

// ListCronReplyValidationError is the validation error returned by
// ListCronReply.Validate if the designated constraints aren't met.
type ListCronReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCronReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCronReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCronReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCronReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCronReplyValidationError) ErrorName() string { return "ListCronReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListCronReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCronReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCronReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCronReplyValidationError{}