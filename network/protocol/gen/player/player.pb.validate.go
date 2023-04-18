// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/player.proto

package player

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

// Validate checks the field values on CSLogin with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CSLogin) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CSLogin with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CSLoginMultiError, or nil if none found.
func (m *CSLogin) ValidateAll() error {
	return m.validate(true)
}

func (m *CSLogin) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UId

	if len(errors) > 0 {
		return CSLoginMultiError(errors)
	}

	return nil
}

// CSLoginMultiError is an error wrapping multiple validation errors returned
// by CSLogin.ValidateAll() if the designated constraints aren't met.
type CSLoginMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CSLoginMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CSLoginMultiError) AllErrors() []error { return m }

// CSLoginValidationError is the validation error returned by CSLogin.Validate
// if the designated constraints aren't met.
type CSLoginValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CSLoginValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CSLoginValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CSLoginValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CSLoginValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CSLoginValidationError) ErrorName() string { return "CSLoginValidationError" }

// Error satisfies the builtin error interface
func (e CSLoginValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCSLogin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CSLoginValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CSLoginValidationError{}

// Validate checks the field values on SCLogin with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SCLogin) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SCLogin with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SCLoginMultiError, or nil if none found.
func (m *SCLogin) ValidateAll() error {
	return m.validate(true)
}

func (m *SCLogin) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return SCLoginMultiError(errors)
	}

	return nil
}

// SCLoginMultiError is an error wrapping multiple validation errors returned
// by SCLogin.ValidateAll() if the designated constraints aren't met.
type SCLoginMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SCLoginMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SCLoginMultiError) AllErrors() []error { return m }

// SCLoginValidationError is the validation error returned by SCLogin.Validate
// if the designated constraints aren't met.
type SCLoginValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SCLoginValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SCLoginValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SCLoginValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SCLoginValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SCLoginValidationError) ErrorName() string { return "SCLoginValidationError" }

// Error satisfies the builtin error interface
func (e SCLoginValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSCLogin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SCLoginValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SCLoginValidationError{}

// Validate checks the field values on CSAddFriend with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CSAddFriend) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CSAddFriend with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CSAddFriendMultiError, or
// nil if none found.
func (m *CSAddFriend) ValidateAll() error {
	return m.validate(true)
}

func (m *CSAddFriend) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UId

	if len(errors) > 0 {
		return CSAddFriendMultiError(errors)
	}

	return nil
}

// CSAddFriendMultiError is an error wrapping multiple validation errors
// returned by CSAddFriend.ValidateAll() if the designated constraints aren't met.
type CSAddFriendMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CSAddFriendMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CSAddFriendMultiError) AllErrors() []error { return m }

// CSAddFriendValidationError is the validation error returned by
// CSAddFriend.Validate if the designated constraints aren't met.
type CSAddFriendValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CSAddFriendValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CSAddFriendValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CSAddFriendValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CSAddFriendValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CSAddFriendValidationError) ErrorName() string { return "CSAddFriendValidationError" }

// Error satisfies the builtin error interface
func (e CSAddFriendValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCSAddFriend.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CSAddFriendValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CSAddFriendValidationError{}

// Validate checks the field values on SCAddFriend with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SCAddFriend) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SCAddFriend with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SCAddFriendMultiError, or
// nil if none found.
func (m *SCAddFriend) ValidateAll() error {
	return m.validate(true)
}

func (m *SCAddFriend) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Desc

	if len(errors) > 0 {
		return SCAddFriendMultiError(errors)
	}

	return nil
}

// SCAddFriendMultiError is an error wrapping multiple validation errors
// returned by SCAddFriend.ValidateAll() if the designated constraints aren't met.
type SCAddFriendMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SCAddFriendMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SCAddFriendMultiError) AllErrors() []error { return m }

// SCAddFriendValidationError is the validation error returned by
// SCAddFriend.Validate if the designated constraints aren't met.
type SCAddFriendValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SCAddFriendValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SCAddFriendValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SCAddFriendValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SCAddFriendValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SCAddFriendValidationError) ErrorName() string { return "SCAddFriendValidationError" }

// Error satisfies the builtin error interface
func (e SCAddFriendValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSCAddFriend.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SCAddFriendValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SCAddFriendValidationError{}

// Validate checks the field values on CSDelFriend with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CSDelFriend) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CSDelFriend with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CSDelFriendMultiError, or
// nil if none found.
func (m *CSDelFriend) ValidateAll() error {
	return m.validate(true)
}

func (m *CSDelFriend) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UId

	if len(errors) > 0 {
		return CSDelFriendMultiError(errors)
	}

	return nil
}

// CSDelFriendMultiError is an error wrapping multiple validation errors
// returned by CSDelFriend.ValidateAll() if the designated constraints aren't met.
type CSDelFriendMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CSDelFriendMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CSDelFriendMultiError) AllErrors() []error { return m }

// CSDelFriendValidationError is the validation error returned by
// CSDelFriend.Validate if the designated constraints aren't met.
type CSDelFriendValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CSDelFriendValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CSDelFriendValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CSDelFriendValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CSDelFriendValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CSDelFriendValidationError) ErrorName() string { return "CSDelFriendValidationError" }

// Error satisfies the builtin error interface
func (e CSDelFriendValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCSDelFriend.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CSDelFriendValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CSDelFriendValidationError{}

// Validate checks the field values on SCDelFriend with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SCDelFriend) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SCDelFriend with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SCDelFriendMultiError, or
// nil if none found.
func (m *SCDelFriend) ValidateAll() error {
	return m.validate(true)
}

func (m *SCDelFriend) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Desc

	if len(errors) > 0 {
		return SCDelFriendMultiError(errors)
	}

	return nil
}

// SCDelFriendMultiError is an error wrapping multiple validation errors
// returned by SCDelFriend.ValidateAll() if the designated constraints aren't met.
type SCDelFriendMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SCDelFriendMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SCDelFriendMultiError) AllErrors() []error { return m }

// SCDelFriendValidationError is the validation error returned by
// SCDelFriend.Validate if the designated constraints aren't met.
type SCDelFriendValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SCDelFriendValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SCDelFriendValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SCDelFriendValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SCDelFriendValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SCDelFriendValidationError) ErrorName() string { return "SCDelFriendValidationError" }

// Error satisfies the builtin error interface
func (e SCDelFriendValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSCDelFriend.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SCDelFriendValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SCDelFriendValidationError{}

// Validate checks the field values on CSSendChatMsg with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CSSendChatMsg) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CSSendChatMsg with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CSSendChatMsgMultiError, or
// nil if none found.
func (m *CSSendChatMsg) ValidateAll() error {
	return m.validate(true)
}

func (m *CSSendChatMsg) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UId

	if all {
		switch v := interface{}(m.GetMsg()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CSSendChatMsgValidationError{
					field:  "Msg",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CSSendChatMsgValidationError{
					field:  "Msg",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMsg()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CSSendChatMsgValidationError{
				field:  "Msg",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Category

	if len(errors) > 0 {
		return CSSendChatMsgMultiError(errors)
	}

	return nil
}

// CSSendChatMsgMultiError is an error wrapping multiple validation errors
// returned by CSSendChatMsg.ValidateAll() if the designated constraints
// aren't met.
type CSSendChatMsgMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CSSendChatMsgMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CSSendChatMsgMultiError) AllErrors() []error { return m }

// CSSendChatMsgValidationError is the validation error returned by
// CSSendChatMsg.Validate if the designated constraints aren't met.
type CSSendChatMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CSSendChatMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CSSendChatMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CSSendChatMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CSSendChatMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CSSendChatMsgValidationError) ErrorName() string { return "CSSendChatMsgValidationError" }

// Error satisfies the builtin error interface
func (e CSSendChatMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCSSendChatMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CSSendChatMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CSSendChatMsgValidationError{}

// Validate checks the field values on SCSendChatMsg with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SCSendChatMsg) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SCSendChatMsg with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SCSendChatMsgMultiError, or
// nil if none found.
func (m *SCSendChatMsg) ValidateAll() error {
	return m.validate(true)
}

func (m *SCSendChatMsg) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SCSendChatMsgMultiError(errors)
	}

	return nil
}

// SCSendChatMsgMultiError is an error wrapping multiple validation errors
// returned by SCSendChatMsg.ValidateAll() if the designated constraints
// aren't met.
type SCSendChatMsgMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SCSendChatMsgMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SCSendChatMsgMultiError) AllErrors() []error { return m }

// SCSendChatMsgValidationError is the validation error returned by
// SCSendChatMsg.Validate if the designated constraints aren't met.
type SCSendChatMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SCSendChatMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SCSendChatMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SCSendChatMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SCSendChatMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SCSendChatMsgValidationError) ErrorName() string { return "SCSendChatMsgValidationError" }

// Error satisfies the builtin error interface
func (e SCSendChatMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSCSendChatMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SCSendChatMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SCSendChatMsgValidationError{}

// Validate checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ChatMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ChatMessageMultiError, or
// nil if none found.
func (m *ChatMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Content

	if len(errors) > 0 {
		return ChatMessageMultiError(errors)
	}

	return nil
}

// ChatMessageMultiError is an error wrapping multiple validation errors
// returned by ChatMessage.ValidateAll() if the designated constraints aren't met.
type ChatMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatMessageMultiError) AllErrors() []error { return m }

// ChatMessageValidationError is the validation error returned by
// ChatMessage.Validate if the designated constraints aren't met.
type ChatMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatMessageValidationError) ErrorName() string { return "ChatMessageValidationError" }

// Error satisfies the builtin error interface
func (e ChatMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatMessageValidationError{}
