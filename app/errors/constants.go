package errors

import (
	"errors"
)

// App Error Definition
var (
	ErrNotFound            = errors.New("not found")
	ErrConflictResource    = errors.New("conflict resource")
	ErrForbidden           = errors.New("forbidden")
	ErrLockedResource      = errors.New("locked resource")
	ErrContextCancelled    = errors.New("context is canceled")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
)

//nolint:stylecheck
var (
	ErrSecom1000001 = errors.New("Processing error: Input [Send date and time] is invalid")
	ErrSecom1000011 = errors.New("Processing error: Input [Sequence number] is invalid")
	ErrSecom1000021 = errors.New("Processing error: Input [Request expiration date] is invalid")
	ErrSecom1000030 = errors.New("Processing error: Account authentication error")
	ErrSecom1000031 = errors.New("Processing error: Incorrect input [Organization ID]")
	ErrSecom1000041 = errors.New("Processing error: Input [API group ID] is invalid")
	ErrSecom1000051 = errors.New("Processing error: Input [API usage ID] is invalid")
	ErrSecom1000061 = errors.New("Processing error: Input [API password] is invalid")
	ErrSecom1000071 = errors.New("Processing error: Input [Number of processes] is invalid")
	ErrSecom1000079 = errors.New("Processing error: Certificate for input [User ID, User Certificate Number] does not exist")
	ErrSecom1000080 = errors.New("Processing error: Invalid certificate for input [User ID, User Certificate Number]")
	ErrSecom1000081 = errors.New("Processing error: Input [User ID] is invalid")
	ErrSecom1000091 = errors.New("Processing error: Input [User certificate number] is incorrect")
	ErrSecom1000099 = errors.New("Processing error: Certificate of input [CN, subdomain, CN serial number] has been issued")
	ErrSecom1000100 = errors.New("Processing error: Incorrect combination of input [CN, subdomain, CN serial number]")
	ErrSecom1000101 = errors.New("Processing error: Input [CN] is invalid")
	ErrSecom1000111 = errors.New("Processing error: Input [subdomain] is invalid")
	ErrSecom1000121 = errors.New("Processing error: Input [CN serial number] is invalid")
	ErrSecom1000131 = errors.New("Processing error: Input [Certificate validity period end date] is invalid")
	ErrSecom1000141 = errors.New("Processing error: Input [Certificate validity period] is invalid")
	ErrSecom1000151 = errors.New("Processing error: Input [subjectAltName key] is invalid")
	ErrSecom1000161 = errors.New("Processing error: Input [subjectAltName value] is invalid")
	ErrSecom1000171 = errors.New("Processing error: Input [otherName key] is invalid")
	ErrSecom1000181 = errors.New("Processing error: Input [otherName value] is invalid")
	ErrSecom1000191 = errors.New("Processing error: Input [extended area key] is invalid")
	ErrSecom1000201 = errors.New("Processing error: Input [extended area value] is invalid")
	ErrSecom1000211 = errors.New("Processing error: Input [Reason for revocation] is invalid")
	ErrSecom1000221 = errors.New("Processing error: Input [comment] is invalid")
	ErrSecom1000231 = errors.New("Processing error: Input [Initial PIN code] is invalid")
	ErrSecom1000241 = errors.New("Processing error: Input [old PIN code] is invalid")
	ErrSecom1000251 = errors.New("Processing error: Input [new PIN code] is invalid")
	ErrSecom1000261 = errors.New("Processing error: Input [PIN code] is invalid")
	ErrSecom1000262 = errors.New("Processing error: Initial PIN code has not been changed")
	ErrSecom1000271 = errors.New("Processing error: Input [Signature method] is invalid")
	ErrSecom1000281 = errors.New("Processing error: Input [Signature generation type] is invalid")
	ErrSecom1000291 = errors.New("Processing error: Input [input mode] is invalid")
	ErrSecom1000299 = errors.New("Processing error: Input [input file or input data] is an unsupported format")
	ErrSecom1000300 = errors.New("Processing error: Input [input file or input data] exceeds file size limit")
	ErrSecom1000301 = errors.New("Processing error: Input [input file path] is invalid")
	ErrSecom1000311 = errors.New("Processing error: Input [input data] is invalid")
	ErrSecom1000321 = errors.New("Processing error: Input [output mode] is invalid")
	ErrSecom1000331 = errors.New("Processing error: Input [output file path] is invalid")
	ErrSecom1000340 = errors.New("Processing error: Input [Visible signature information] is invalid")
	ErrSecom1000341 = errors.New("Processing error: Input [coordinate X] is invalid")
	ErrSecom1000351 = errors.New("Processing error: Input [coordinate Y] is invalid")
	ErrSecom1000361 = errors.New("Processing error: Input [width] is invalid")
	ErrSecom1000371 = errors.New("Processing error: Input [height] is incorrect")
	ErrSecom1000381 = errors.New("Processing error: Input [Page number information] is invalid")
	ErrSecom1000391 = errors.New("Processing error: Input [imprint No] is invalid")
	ErrSecom1000401 = errors.New("Processing error: Input [image data] is invalid")
	ErrSecom1000411 = errors.New("Processing error: Input [visible signature character] is invalid")
	ErrSecom1000421 = errors.New("Processing error: Input [Display position] is invalid")
	ErrSecom1000431 = errors.New("Processing error: Input [font used] is invalid")
	ErrSecom1000441 = errors.New("Processing error: Input [font size] is incorrect")
	ErrSecom1000451 = errors.New("Processing error: Input [font color] is incorrect")
	ErrSecom1000461 = errors.New("Processing error: Input [Input mode (signature verification)] is invalid")
	ErrSecom1000471 = errors.New("Processing error: Input [Create verification result type] is invalid")
	ErrSecom1000481 = errors.New("Processing error: Input [Delivery return method type] is invalid")
	ErrSecom1000491 = errors.New("Processing error: Input [Data to be verified] is invalid")
	ErrSecom1000501 = errors.New("Processing error: Input [Signature data] is invalid")
	ErrSecom1000511 = errors.New("Processing error: Input [Attachment information List] is invalid")
	ErrSecom1000520 = errors.New("Processing error: Total of input [input file, attachment] exceeds file size limit")
	ErrSecom1000521 = errors.New("Processing error: Input [Attachment path] is invalid")
	ErrSecom1000531 = errors.New("Processing error: Input [Attachment data] is invalid")
	ErrSecom1000541 = errors.New("Processing error: Input [Attachment name] is invalid")
	ErrSecom1000551 = errors.New("Processing error: Input [Attachment description] is invalid")
	ErrSecom1000561 = errors.New("Processing error: Input [Attachment creation date] is invalid")
	ErrSecom1000571 = errors.New("Processing error: Input [Attachment update date] is invalid")
	ErrSecom1000581 = errors.New("Processing error: Input [Attachment additional item name] is invalid")
	ErrSecom1000591 = errors.New("Processing error: Input [Attachment additional item value (type)] is invalid")
	ErrSecom1000601 = errors.New("Processing error: Input [Attachment additional item value (character)] is invalid")
	ErrSecom1000611 = errors.New("Processing error: Input [Attachment additional item value (numeric value)] is invalid")
	ErrSecom1000621 = errors.New("Processing error: Input [Attachment additional item value (date and time)] is invalid")
	ErrSecom1000631 = errors.New("Processing error: Input [Attachment appendix name prefix] is invalid")
	ErrSecom1100001 = errors.New("Processing error: The request has been processed in the past")
	ErrSecom1200001 = errors.New("Processing completed: Verification result error")
	ErrSecom1200002 = errors.New("Process completed: Input [input file or input data] has already been signed")
	ErrSecom2000001 = errors.New("Communication error")
	ErrSecom2000011 = errors.New("Timeout error")
	ErrSecom2000021 = errors.New("Certificate with the same public key has been issued")
	ErrSecom2000031 = errors.New("Possible hardware failure")
	ErrSecom3000001 = errors.New("System error")
)

var (
	ErrSalesforceAppNotInstalled  = errors.New("app is not installed")
	ErrSalesforceInvalidSessionID = errors.New("invalid session id")
	ErrSalesforceResourceNotFound = errors.New("resource not found")
)