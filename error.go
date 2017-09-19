package wyre

import (
	"fmt"
)

type APIError struct {
	Language      string          `json:"language"`
	ExceptionID   string          `json:"exceptionId"`
	CompositeType string          `json:"compositeType"`
	SubType       APIErrorSubType `json:"subType"`
	Message       string          `json:"message"`
	Type          APIErrorType    `json:"type"`
	Transient     bool            `json:"transient"`
}

func (e *APIError) Error() string {
	if e.Type == ERR_TYPE_VALIDATION {
		return fmt.Sprintf("[%s] %s(%s): %s", e.ExceptionID, e.Type, e.SubType, e.Message)
	} else {
		return fmt.Sprintf("[%s] %s: %s", e.ExceptionID, e.Type, e.Message)
	}
}

type APIErrorType string

const (
	ERR_TYPE_VALIDATION         APIErrorType = "ValidationException"
	ERR_TYPE_UNKNOWN            APIErrorType = "UnknownException"
	ERR_TYPE_INSUFFICIENT_FUNDS APIErrorType = "InsufficientFundsException"
	ERR_TYPE_RATE_LIMIT         APIErrorType = "RateLimitException"
	ERR_TYPE_ACCESS_DENIED      APIErrorType = "AccessDeniedException"
	ERR_TYPE_TRANSFER           APIErrorType = "TransferException"
	ERR_TYPE_NOT_FOUND          APIErrorType = "NotFoundException"
	ERR_TYPE_CUSTOMER_SUPPORT   APIErrorType = "CustomerSupportException"
	ERR_TYPE_MFA_REQUIRED       APIErrorType = "MFARequiredException"
)

type APIErrorSubType string

const (
	ERR_SUBTYPE_FIELD_REQUIRED                                          APIErrorSubType = "FIELD_REQUIRED"
	ERR_SUBTYPE_INVALID_VALUE                                           APIErrorSubType = "INVALID_VALUE"
	ERR_SUBTYPE_TRANSACTION_AMOUNT_TOO_SMALL                            APIErrorSubType = "TRANSACTION_AMOUNT_TOO_SMALL"
	ERR_SUBTYPE_UNSUPPORTED_SOURCE_CURRENCY                             APIErrorSubType = "UNSUPPORTED_SOURCE_CURRENCY"
	ERR_SUBTYPE_SENDER_PROVIDED_ID_IN_USE                               APIErrorSubType = "SENDER_PROVIDED_ID_IN_USE"
	ERR_SUBTYPE_CANNOT_SEND_SELF_FUNDS                                  APIErrorSubType = "CANNOT_SEND_SELF_FUNDS"
	ERR_SUBTYPE_INVALID_PAYMENT_METHOD                                  APIErrorSubType = "INVALID_PAYMENT_METHOD"
	ERR_SUBTYPE_PAYMENT_METHOD_INACTIVE                                 APIErrorSubType = "PAYMENT_METHOD_INACTIVE"
	ERR_SUBTYPE_PAYMENT_METHOD_UNSUPPORTED_CHARGE_CURRENCY              APIErrorSubType = "PAYMENT_METHOD_UNSUPPORTED_CHARGE_CURRENCY"
	ERR_SUBTYPE_PAYMENT_METHOD_UNCHARGEABLE                             APIErrorSubType = "PAYMENT_METHOD_UNCHARGEABLE"
	ERR_SUBTYPE_PAYMENT_METHOD_UNSUPPORTED_DEPOSIT_CURRENCY             APIErrorSubType = "PAYMENT_METHOD_UNSUPPORTED_DEPOSIT_CURRENCY"
	ERR_SUBTYPE_PAYMENT_METHOD_UNDEPOSITABLE                            APIErrorSubType = "PAYMENT_METHOD_UNDEPOSITABLE"
	ERR_SUBTYPE_PAYMENT_METHOD_DOESNT_SUPPORT_FOLLOWUPS                 APIErrorSubType = "PAYMENT_METHOD_DOESNT_SUPPORT_FOLLOWUPS"
	ERR_SUBTYPE_PAYMENT_METHOD_DOESNT_SUPPORT_MICRODEPOSIT_VERIFICATION APIErrorSubType = "PAYMENT_METHOD_DOESNT_SUPPORT_MICRODEPOSIT_VERIFICATION"
)
