// Code generated by protoc-gen-validate
// source: github.com/infobloxopen/atlas-contacts-app/proto/contacts.proto
// DO NOT EDIT!!!

package contacts

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Contact with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Contact) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	if err := m._validateEmail(m.GetEmailAddress()); err != nil {
		return ContactValidationError{
			Field:  "EmailAddress",
			Reason: "value must be a valid email address",
			Cause:  err,
		}
	}

	return nil
}

func (m *Contact) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *Contact) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// ContactValidationError is the validation error returned by Contact.Validate
// if the designated constraints aren't met.
type ContactValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ContactValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContact.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ContactValidationError{}

// Validate checks the field values on TagCondition with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TagCondition) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Tag

	// no validation rules for Op

	// no validation rules for Value

	return nil
}

// TagConditionValidationError is the validation error returned by
// TagCondition.Validate if the designated constraints aren't met.
type TagConditionValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TagConditionValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTagCondition.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TagConditionValidationError{}

// Validate checks the field values on TagQuery with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TagQuery) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetConditions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return TagQueryValidationError{
					Field:  fmt.Sprintf("Conditions[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// TagQueryValidationError is the validation error returned by
// TagQuery.Validate if the designated constraints aren't met.
type TagQueryValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TagQueryValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTagQuery.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TagQueryValidationError{}

// Validate checks the field values on PagingOptions with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PagingOptions) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetLimit(); val < 0 || val >= 1000 {
		return PagingOptionsValidationError{
			Field:  "Limit",
			Reason: "value must be inside range [0, 1000)",
		}
	}

	// no validation rules for Start

	// no validation rules for Offset

	// no validation rules for Total

	return nil
}

// PagingOptionsValidationError is the validation error returned by
// PagingOptions.Validate if the designated constraints aren't met.
type PagingOptionsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e PagingOptionsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagingOptions.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = PagingOptionsValidationError{}

// Validate checks the field values on SearchCriteria with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchCriteria) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	// no validation rules for EmailAddress

	return nil
}

// SearchCriteriaValidationError is the validation error returned by
// SearchCriteria.Validate if the designated constraints aren't met.
type SearchCriteriaValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e SearchCriteriaValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchCriteria.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = SearchCriteriaValidationError{}

// Validate checks the field values on SearchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCriteria()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return SearchRequestValidationError{
				Field:  "Criteria",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPaging()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return SearchRequestValidationError{
				Field:  "Paging",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTags()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return SearchRequestValidationError{
				Field:  "Tags",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// SearchRequestValidationError is the validation error returned by
// SearchRequest.Validate if the designated constraints aren't met.
type SearchRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e SearchRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = SearchRequestValidationError{}

// Validate checks the field values on ContactPage with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ContactPage) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPaging()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return ContactPageValidationError{
				Field:  "Paging",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetContacts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return ContactPageValidationError{
					Field:  fmt.Sprintf("Contacts[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// ContactPageValidationError is the validation error returned by
// ContactPage.Validate if the designated constraints aren't met.
type ContactPageValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ContactPageValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContactPage.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ContactPageValidationError{}
