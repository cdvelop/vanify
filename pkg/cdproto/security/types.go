package security

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"fmt"

	"github.com/cdvelop/vanify/pkg/cdproto/cdp"
	"github.com/cdvelop/vanify/pkg/easyjson"
	"github.com/cdvelop/vanify/pkg/easyjson/jlexer"
	"github.com/cdvelop/vanify/pkg/easyjson/jwriter"
)

// CertificateID an internal certificate ID value.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#type-CertificateId
type CertificateID int64

// Int64 returns the CertificateID as int64 value.
func (t CertificateID) Int64() int64 {
	return int64(t)
}

// MixedContentType a description of mixed content (HTTP resources on HTTPS
// pages), as defined by https://www.w3.org/TR/mixed-content/#categories.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#type-MixedContentType
type MixedContentType string

// String returns the MixedContentType as string value.
func (t MixedContentType) String() string {
	return string(t)
}

// MixedContentType values.
const (
	MixedContentTypeBlockable           MixedContentType = "blockable"
	MixedContentTypeOptionallyBlockable MixedContentType = "optionally-blockable"
	MixedContentTypeNone                MixedContentType = "none"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t MixedContentType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t MixedContentType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *MixedContentType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch MixedContentType(v) {
	case MixedContentTypeBlockable:
		*t = MixedContentTypeBlockable
	case MixedContentTypeOptionallyBlockable:
		*t = MixedContentTypeOptionallyBlockable
	case MixedContentTypeNone:
		*t = MixedContentTypeNone

	default:
		in.AddError(fmt.Errorf("unknown MixedContentType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *MixedContentType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// State the security level of a page or resource.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#type-SecurityState
type State string

// String returns the State as string value.
func (t State) String() string {
	return string(t)
}

// State values.
const (
	StateUnknown        State = "unknown"
	StateNeutral        State = "neutral"
	StateInsecure       State = "insecure"
	StateSecure         State = "secure"
	StateInfo           State = "info"
	StateInsecureBroken State = "insecure-broken"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t State) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t State) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *State) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch State(v) {
	case StateUnknown:
		*t = StateUnknown
	case StateNeutral:
		*t = StateNeutral
	case StateInsecure:
		*t = StateInsecure
	case StateSecure:
		*t = StateSecure
	case StateInfo:
		*t = StateInfo
	case StateInsecureBroken:
		*t = StateInsecureBroken

	default:
		in.AddError(fmt.Errorf("unknown State value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *State) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// CertificateSecurityState details about the security state of the page
// certificate.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#type-CertificateSecurityState
type CertificateSecurityState struct {
	Protocol                    string              `json:"protocol"`                          // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                 string              `json:"keyExchange"`                       // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup            string              `json:"keyExchangeGroup,omitempty"`        // (EC)DH group used by the connection, if applicable.
	Cipher                      string              `json:"cipher"`                            // Cipher name.
	Mac                         string              `json:"mac,omitempty"`                     // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Certificate                 []string            `json:"certificate"`                       // Page certificate.
	SubjectName                 string              `json:"subjectName"`                       // Certificate subject name.
	Issuer                      string              `json:"issuer"`                            // Name of the issuing CA.
	ValidFrom                   *cdp.TimeSinceEpoch `json:"validFrom"`                         // Certificate valid from date.
	ValidTo                     *cdp.TimeSinceEpoch `json:"validTo"`                           // Certificate valid to (expiration) date
	CertificateNetworkError     string              `json:"certificateNetworkError,omitempty"` // The highest priority network error code, if the certificate has an error.
	CertificateHasWeakSignature bool                `json:"certificateHasWeakSignature"`       // True if the certificate uses a weak signature algorithm.
	CertificateHasSha1signature bool                `json:"certificateHasSha1Signature"`       // True if the certificate has a SHA1 signature in the chain.
	ModernSSL                   bool                `json:"modernSSL"`                         // True if modern SSL
	ObsoleteSslProtocol         bool                `json:"obsoleteSslProtocol"`               // True if the connection is using an obsolete SSL protocol.
	ObsoleteSslKeyExchange      bool                `json:"obsoleteSslKeyExchange"`            // True if the connection is using an obsolete SSL key exchange.
	ObsoleteSslCipher           bool                `json:"obsoleteSslCipher"`                 // True if the connection is using an obsolete SSL cipher.
	ObsoleteSslSignature        bool                `json:"obsoleteSslSignature"`              // True if the connection is using an obsolete SSL signature.
}

// SafetyTipStatus [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#type-SafetyTipStatus
type SafetyTipStatus string

// String returns the SafetyTipStatus as string value.
func (t SafetyTipStatus) String() string {
	return string(t)
}

// SafetyTipStatus values.
const (
	SafetyTipStatusBadReputation SafetyTipStatus = "badReputation"
	SafetyTipStatusLookalike     SafetyTipStatus = "lookalike"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SafetyTipStatus) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SafetyTipStatus) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SafetyTipStatus) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch SafetyTipStatus(v) {
	case SafetyTipStatusBadReputation:
		*t = SafetyTipStatusBadReputation
	case SafetyTipStatusLookalike:
		*t = SafetyTipStatusLookalike

	default:
		in.AddError(fmt.Errorf("unknown SafetyTipStatus value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SafetyTipStatus) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SafetyTipInfo [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#type-SafetyTipInfo
type SafetyTipInfo struct {
	SafetyTipStatus SafetyTipStatus `json:"safetyTipStatus"`   // Describes whether the page triggers any safety tips or reputation warnings. Default is unknown.
	SafeURL         string          `json:"safeUrl,omitempty"` // The URL the safety tip suggested ("Did you mean?"). Only filled in for lookalike matches.
}

// VisibleSecurityState security state information about the page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#type-VisibleSecurityState
type VisibleSecurityState struct {
	SecurityState            State                     `json:"securityState"`                      // The security level of the page.
	CertificateSecurityState *CertificateSecurityState `json:"certificateSecurityState,omitempty"` // Security state details about the page certificate.
	SafetyTipInfo            *SafetyTipInfo            `json:"safetyTipInfo,omitempty"`            // The type of Safety Tip triggered on the page. Note that this field will be set even if the Safety Tip UI was not actually shown.
	SecurityStateIssueIDs    []string                  `json:"securityStateIssueIds"`              // Array of security state issues ids.
}

// StateExplanation an explanation of an factor contributing to the security
// state.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#type-SecurityStateExplanation
type StateExplanation struct {
	SecurityState    State            `json:"securityState"`             // Security state representing the severity of the factor being explained.
	Title            string           `json:"title"`                     // Title describing the type of factor.
	Summary          string           `json:"summary"`                   // Short phrase describing the type of factor.
	Description      string           `json:"description"`               // Full text explanation of the factor.
	MixedContentType MixedContentType `json:"mixedContentType"`          // The type of mixed content described by the explanation.
	Certificate      []string         `json:"certificate"`               // Page certificate.
	Recommendations  []string         `json:"recommendations,omitempty"` // Recommendations to fix any issues.
}

// CertificateErrorAction the action to take when a certificate error occurs.
// continue will continue processing the request and cancel will cancel the
// request.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#type-CertificateErrorAction
type CertificateErrorAction string

// String returns the CertificateErrorAction as string value.
func (t CertificateErrorAction) String() string {
	return string(t)
}

// CertificateErrorAction values.
const (
	CertificateErrorActionContinue CertificateErrorAction = "continue"
	CertificateErrorActionCancel   CertificateErrorAction = "cancel"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CertificateErrorAction) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CertificateErrorAction) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CertificateErrorAction) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch CertificateErrorAction(v) {
	case CertificateErrorActionContinue:
		*t = CertificateErrorActionContinue
	case CertificateErrorActionCancel:
		*t = CertificateErrorActionCancel

	default:
		in.AddError(fmt.Errorf("unknown CertificateErrorAction value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *CertificateErrorAction) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
