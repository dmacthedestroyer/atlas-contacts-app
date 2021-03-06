// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/infobloxopen/atlas-contacts-app/pkg/pb/contacts.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	github.com/infobloxopen/atlas-contacts-app/pkg/pb/contacts.proto

It has these top-level messages:
	Contact
	Email
	CreateContactRequest
	CreateContactResponse
	ReadContactRequest
	ReadContactResponse
	UpdateContactRequest
	UpdateContactResponse
	DeleteContactRequest
	ListContactsResponse
	SMSRequest
*/
package pb

import context "context"
import errors "errors"

import auth "github.com/infobloxopen/atlas-app-toolkit/auth"
import gorm "github.com/jinzhu/gorm"
import ops "github.com/infobloxopen/atlas-app-toolkit/gorm"

import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

// ContactORM no comment was provided for message type
type ContactORM struct {
	AccountID  string
	Id         uint64
	FirstName  string
	MiddleName string
	LastName   string
	// Skipping field from proto option: PrimaryEmail
	Emails []*EmailORM `gorm:"foreignkey:ContactId"`
}

// TableName overrides the default tablename generated by GORM
func (ContactORM) TableName() string {
	return "contacts"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Contact) ToORM(ctx context.Context) (ContactORM, error) {
	to := ContactORM{}
	var err error
	if prehook, ok := interface{}(m).(ContactWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.FirstName = m.FirstName
	to.MiddleName = m.MiddleName
	to.LastName = m.LastName
	// Skipping field: PrimaryEmail
	for _, v := range m.Emails {
		if v != nil {
			if tempEmails, cErr := v.ToORM(ctx); cErr == nil {
				to.Emails = append(to.Emails, &tempEmails)
			} else {
				return to, cErr
			}
		} else {
			to.Emails = append(to.Emails, nil)
		}
	}
	if posthook, ok := interface{}(m).(ContactWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ContactORM) ToPB(ctx context.Context) (Contact, error) {
	to := Contact{}
	var err error
	if prehook, ok := interface{}(m).(ContactWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.FirstName = m.FirstName
	to.MiddleName = m.MiddleName
	to.LastName = m.LastName
	// Skipping field: PrimaryEmail
	for _, v := range m.Emails {
		if v != nil {
			if tempEmails, cErr := v.ToPB(ctx); cErr == nil {
				to.Emails = append(to.Emails, &tempEmails)
			} else {
				return to, cErr
			}
		} else {
			to.Emails = append(to.Emails, nil)
		}
	}
	if posthook, ok := interface{}(m).(ContactWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Contact the arg will be the target, the caller the one being converted from

// ContactBeforeToORM called before default ToORM code
type ContactWithBeforeToORM interface {
	BeforeToORM(context.Context, *ContactORM) error
}

// ContactAfterToORM called after default ToORM code
type ContactWithAfterToORM interface {
	AfterToORM(context.Context, *ContactORM) error
}

// ContactBeforeToPB called before default ToPB code
type ContactWithBeforeToPB interface {
	BeforeToPB(context.Context, *Contact) error
}

// ContactAfterToPB called after default ToPB code
type ContactWithAfterToPB interface {
	AfterToPB(context.Context, *Contact) error
}

// EmailORM no comment was provided for message type
type EmailORM struct {
	IsPrimary bool
	ContactId uint64
	Id        uint64
	Address   string `gorm:"UNIQUE"`
}

// TableName overrides the default tablename generated by GORM
func (EmailORM) TableName() string {
	return "emails"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Email) ToORM(ctx context.Context) (EmailORM, error) {
	to := EmailORM{}
	var err error
	if prehook, ok := interface{}(m).(EmailWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Address = m.Address
	if posthook, ok := interface{}(m).(EmailWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *EmailORM) ToPB(ctx context.Context) (Email, error) {
	to := Email{}
	var err error
	if prehook, ok := interface{}(m).(EmailWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Address = m.Address
	if posthook, ok := interface{}(m).(EmailWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Email the arg will be the target, the caller the one being converted from

// EmailBeforeToORM called before default ToORM code
type EmailWithBeforeToORM interface {
	BeforeToORM(context.Context, *EmailORM) error
}

// EmailAfterToORM called after default ToORM code
type EmailWithAfterToORM interface {
	AfterToORM(context.Context, *EmailORM) error
}

// EmailBeforeToPB called before default ToPB code
type EmailWithBeforeToPB interface {
	BeforeToPB(context.Context, *Email) error
}

// EmailAfterToPB called after default ToPB code
type EmailWithAfterToPB interface {
	AfterToPB(context.Context, *Email) error
}

////////////////////////// CURDL for objects
// DefaultCreateContact executes a basic gorm create call
func DefaultCreateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateContact")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	accountID, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	ormObj.AccountID = accountID
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadContact executes a basic gorm read call
func DefaultReadContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadContact")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	accountID, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	ormParams.AccountID = accountID
	ormResponse := ContactORM{}
	if err = db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateContact executes a basic gorm update call
func DefaultUpdateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateContact")
	}
	if exists, err := DefaultReadContact(ctx, &Contact{Id: in.GetId()}, db); err != nil {
		return nil, err
	} else if exists == nil {
		return nil, errors.New("Contact not found")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteContact(ctx context.Context, in *Contact, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteContact")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	accountID, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return err
	}
	ormObj.AccountID = accountID
	err = db.Where(&ormObj).Delete(&ContactORM{}).Error
	return err
}

// DefaultListContact executes a gorm list call
func DefaultListContact(ctx context.Context, db *gorm.DB) ([]*Contact, error) {
	ormResponse := []ContactORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	accountID, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	db = db.Where(&ContactORM{AccountID: accountID})
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Contact{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultStrictUpdateContact clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateContact")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	filterObjEmail := EmailORM{}
	if ormObj.Id == 0 {
		return nil, errors.New("Can't do overwriting update with no 'Id' value for FK of field 'Emails'")
	}
	filterObjEmail.ContactId = ormObj.Id
	if err = db.Where(filterObjEmail).Delete(Email{}).Error; err != nil {
		return nil, err
	}
	accountID, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	db = db.Where(&ContactORM{AccountID: accountID})
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultCreateEmail executes a basic gorm create call
func DefaultCreateEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadEmail executes a basic gorm read call
func DefaultReadEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadEmail")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := EmailORM{}
	if err = db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateEmail executes a basic gorm update call
func DefaultUpdateEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteEmail(ctx context.Context, in *Email, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	err = db.Where(&ormObj).Delete(&EmailORM{}).Error
	return err
}

// DefaultListEmail executes a gorm list call
func DefaultListEmail(ctx context.Context, db *gorm.DB) ([]*Email, error) {
	ormResponse := []EmailORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Email{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultStrictUpdateEmail clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

type ContactsDefaultServer struct {
	DB *gorm.DB
}
type ContactsCreateCustomHandler interface {
	CustomCreate(context.Context, *CreateContactRequest) (*CreateContactResponse, error)
}

// Create ...
func (m *ContactsDefaultServer) Create(ctx context.Context, in *CreateContactRequest) (*CreateContactResponse, error) {
	if custom, ok := interface{}(m).(ContactsCreateCustomHandler); ok {
		return custom.CustomCreate(ctx, in)
	}
	res, err := DefaultCreateContact(ctx, in.GetPayload(), m.DB)
	if err != nil {
		return nil, err
	}
	return &CreateContactResponse{Result: res}, nil
}

type ContactsReadCustomHandler interface {
	CustomRead(context.Context, *ReadContactRequest) (*ReadContactResponse, error)
}

// Read ...
func (m *ContactsDefaultServer) Read(ctx context.Context, in *ReadContactRequest) (*ReadContactResponse, error) {
	if custom, ok := interface{}(m).(ContactsReadCustomHandler); ok {
		return custom.CustomRead(ctx, in)
	}
	res, err := DefaultReadContact(ctx, &Contact{Id: in.GetId()}, m.DB)
	if err != nil {
		return nil, err
	}
	return &ReadContactResponse{Result: res}, nil
}

type ContactsUpdateCustomHandler interface {
	CustomUpdate(context.Context, *UpdateContactRequest) (*UpdateContactResponse, error)
}

// Update ...
func (m *ContactsDefaultServer) Update(ctx context.Context, in *UpdateContactRequest) (*UpdateContactResponse, error) {
	if custom, ok := interface{}(m).(ContactsUpdateCustomHandler); ok {
		return custom.CustomUpdate(ctx, in)
	}
	res, err := DefaultStrictUpdateContact(ctx, in.GetPayload(), m.DB)
	if err != nil {
		return nil, err
	}
	return &UpdateContactResponse{Result: res}, nil
}

type ContactsDeleteCustomHandler interface {
	CustomDelete(context.Context, *DeleteContactRequest) (*google_protobuf.Empty, error)
}

// Delete ...
func (m *ContactsDefaultServer) Delete(ctx context.Context, in *DeleteContactRequest) (*google_protobuf.Empty, error) {
	if custom, ok := interface{}(m).(ContactsDeleteCustomHandler); ok {
		return custom.CustomDelete(ctx, in)
	}
	return &google_protobuf.Empty{}, DefaultDeleteContact(ctx, &Contact{Id: in.GetId()}, m.DB)
}

type ContactsListCustomHandler interface {
	CustomList(context.Context, *google_protobuf.Empty) (*ListContactsResponse, error)
}

// List ...
func (m *ContactsDefaultServer) List(ctx context.Context, in *google_protobuf.Empty) (*ListContactsResponse, error) {
	if custom, ok := interface{}(m).(ContactsListCustomHandler); ok {
		return custom.CustomList(ctx, in)
	}
	res, err := DefaultListContact(ctx, m.DB)
	if err != nil {
		return nil, err
	}
	return &ListContactsResponse{Results: res}, nil
}

type ContactsSendSMSCustomHandler interface {
	CustomSendSMS(context.Context, *SMSRequest) (*google_protobuf.Empty, error)
}

// SendSMS ...
func (m *ContactsDefaultServer) SendSMS(ctx context.Context, in *SMSRequest) (*google_protobuf.Empty, error) {
	if custom, ok := interface{}(m).(ContactsSendSMSCustomHandler); ok {
		return custom.CustomSendSMS(ctx, in)
	}
	return &google_protobuf.Empty{}, nil
}
