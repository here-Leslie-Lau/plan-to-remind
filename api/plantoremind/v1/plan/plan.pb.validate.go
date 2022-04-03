// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/plantoremind/v1/plan/plan.proto

package plan

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

// Validate checks the field values on CreatePlanRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreatePlanRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePlanRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePlanRequestMultiError, or nil if none found.
func (m *CreatePlanRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePlanRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreatePlanRequestMultiError(errors)
	}

	return nil
}

// CreatePlanRequestMultiError is an error wrapping multiple validation errors
// returned by CreatePlanRequest.ValidateAll() if the designated constraints
// aren't met.
type CreatePlanRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePlanRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePlanRequestMultiError) AllErrors() []error { return m }

// CreatePlanRequestValidationError is the validation error returned by
// CreatePlanRequest.Validate if the designated constraints aren't met.
type CreatePlanRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePlanRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePlanRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePlanRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePlanRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePlanRequestValidationError) ErrorName() string {
	return "CreatePlanRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePlanRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePlanRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePlanRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePlanRequestValidationError{}

// Validate checks the field values on UpdatePlanRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdatePlanRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePlanRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePlanRequestMultiError, or nil if none found.
func (m *UpdatePlanRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePlanRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdatePlanRequestMultiError(errors)
	}

	return nil
}

// UpdatePlanRequestMultiError is an error wrapping multiple validation errors
// returned by UpdatePlanRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdatePlanRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePlanRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePlanRequestMultiError) AllErrors() []error { return m }

// UpdatePlanRequestValidationError is the validation error returned by
// UpdatePlanRequest.Validate if the designated constraints aren't met.
type UpdatePlanRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePlanRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePlanRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePlanRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePlanRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePlanRequestValidationError) ErrorName() string {
	return "UpdatePlanRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePlanRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePlanRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePlanRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePlanRequestValidationError{}

// Validate checks the field values on UpdatePlanReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdatePlanReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePlanReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePlanReplyMultiError, or nil if none found.
func (m *UpdatePlanReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePlanReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdatePlanReplyMultiError(errors)
	}

	return nil
}

// UpdatePlanReplyMultiError is an error wrapping multiple validation errors
// returned by UpdatePlanReply.ValidateAll() if the designated constraints
// aren't met.
type UpdatePlanReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePlanReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePlanReplyMultiError) AllErrors() []error { return m }

// UpdatePlanReplyValidationError is the validation error returned by
// UpdatePlanReply.Validate if the designated constraints aren't met.
type UpdatePlanReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePlanReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePlanReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePlanReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePlanReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePlanReplyValidationError) ErrorName() string { return "UpdatePlanReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdatePlanReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePlanReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePlanReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePlanReplyValidationError{}

// Validate checks the field values on DeletePlanRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeletePlanRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePlanRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePlanRequestMultiError, or nil if none found.
func (m *DeletePlanRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePlanRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeletePlanRequestMultiError(errors)
	}

	return nil
}

// DeletePlanRequestMultiError is an error wrapping multiple validation errors
// returned by DeletePlanRequest.ValidateAll() if the designated constraints
// aren't met.
type DeletePlanRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePlanRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePlanRequestMultiError) AllErrors() []error { return m }

// DeletePlanRequestValidationError is the validation error returned by
// DeletePlanRequest.Validate if the designated constraints aren't met.
type DeletePlanRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePlanRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePlanRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePlanRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePlanRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePlanRequestValidationError) ErrorName() string {
	return "DeletePlanRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePlanRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePlanRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePlanRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePlanRequestValidationError{}

// Validate checks the field values on DeletePlanReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeletePlanReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePlanReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePlanReplyMultiError, or nil if none found.
func (m *DeletePlanReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePlanReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeletePlanReplyMultiError(errors)
	}

	return nil
}

// DeletePlanReplyMultiError is an error wrapping multiple validation errors
// returned by DeletePlanReply.ValidateAll() if the designated constraints
// aren't met.
type DeletePlanReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePlanReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePlanReplyMultiError) AllErrors() []error { return m }

// DeletePlanReplyValidationError is the validation error returned by
// DeletePlanReply.Validate if the designated constraints aren't met.
type DeletePlanReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePlanReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePlanReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePlanReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePlanReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePlanReplyValidationError) ErrorName() string { return "DeletePlanReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeletePlanReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePlanReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePlanReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePlanReplyValidationError{}

// Validate checks the field values on PlanData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PlanData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlanData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PlanDataMultiError, or nil
// if none found.
func (m *PlanData) ValidateAll() error {
	return m.validate(true)
}

func (m *PlanData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for State

	// no validation rules for Level

	// no validation rules for CronId

	// no validation rules for DeadTime

	// no validation rules for Name

	// no validation rules for CronDesc

	if len(errors) > 0 {
		return PlanDataMultiError(errors)
	}

	return nil
}

// PlanDataMultiError is an error wrapping multiple validation errors returned
// by PlanData.ValidateAll() if the designated constraints aren't met.
type PlanDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlanDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlanDataMultiError) AllErrors() []error { return m }

// PlanDataValidationError is the validation error returned by
// PlanData.Validate if the designated constraints aren't met.
type PlanDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlanDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlanDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlanDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlanDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlanDataValidationError) ErrorName() string { return "PlanDataValidationError" }

// Error satisfies the builtin error interface
func (e PlanDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlanData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlanDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlanDataValidationError{}

// Validate checks the field values on GetPlanRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetPlanRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPlanRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetPlanRequestMultiError,
// or nil if none found.
func (m *GetPlanRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPlanRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetPlanRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetPlanRequestMultiError(errors)
	}

	return nil
}

// GetPlanRequestMultiError is an error wrapping multiple validation errors
// returned by GetPlanRequest.ValidateAll() if the designated constraints
// aren't met.
type GetPlanRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPlanRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPlanRequestMultiError) AllErrors() []error { return m }

// GetPlanRequestValidationError is the validation error returned by
// GetPlanRequest.Validate if the designated constraints aren't met.
type GetPlanRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPlanRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPlanRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPlanRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPlanRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPlanRequestValidationError) ErrorName() string { return "GetPlanRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetPlanRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPlanRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPlanRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPlanRequestValidationError{}

// Validate checks the field values on GetPlanReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetPlanReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPlanReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetPlanReplyMultiError, or
// nil if none found.
func (m *GetPlanReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPlanReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPlan()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetPlanReplyValidationError{
					field:  "Plan",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetPlanReplyValidationError{
					field:  "Plan",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPlan()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPlanReplyValidationError{
				field:  "Plan",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetPlanReplyMultiError(errors)
	}

	return nil
}

// GetPlanReplyMultiError is an error wrapping multiple validation errors
// returned by GetPlanReply.ValidateAll() if the designated constraints aren't met.
type GetPlanReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPlanReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPlanReplyMultiError) AllErrors() []error { return m }

// GetPlanReplyValidationError is the validation error returned by
// GetPlanReply.Validate if the designated constraints aren't met.
type GetPlanReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPlanReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPlanReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPlanReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPlanReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPlanReplyValidationError) ErrorName() string { return "GetPlanReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetPlanReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPlanReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPlanReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPlanReplyValidationError{}

// Validate checks the field values on ListPlanRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListPlanRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPlanRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPlanRequestMultiError, or nil if none found.
func (m *ListPlanRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPlanRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListPlanRequestMultiError(errors)
	}

	return nil
}

// ListPlanRequestMultiError is an error wrapping multiple validation errors
// returned by ListPlanRequest.ValidateAll() if the designated constraints
// aren't met.
type ListPlanRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPlanRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPlanRequestMultiError) AllErrors() []error { return m }

// ListPlanRequestValidationError is the validation error returned by
// ListPlanRequest.Validate if the designated constraints aren't met.
type ListPlanRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPlanRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPlanRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPlanRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPlanRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPlanRequestValidationError) ErrorName() string { return "ListPlanRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListPlanRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPlanRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPlanRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPlanRequestValidationError{}

// Validate checks the field values on ListPlanReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListPlanReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPlanReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListPlanReplyMultiError, or
// nil if none found.
func (m *ListPlanReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPlanReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListPlanReplyMultiError(errors)
	}

	return nil
}

// ListPlanReplyMultiError is an error wrapping multiple validation errors
// returned by ListPlanReply.ValidateAll() if the designated constraints
// aren't met.
type ListPlanReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPlanReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPlanReplyMultiError) AllErrors() []error { return m }

// ListPlanReplyValidationError is the validation error returned by
// ListPlanReply.Validate if the designated constraints aren't met.
type ListPlanReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPlanReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPlanReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPlanReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPlanReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPlanReplyValidationError) ErrorName() string { return "ListPlanReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListPlanReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPlanReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPlanReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPlanReplyValidationError{}
