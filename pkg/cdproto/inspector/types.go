package inspector

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"fmt"

	"github.com/cdvelop/vanify/pkg/easyjson"
	"github.com/cdvelop/vanify/pkg/easyjson/jlexer"
	"github.com/cdvelop/vanify/pkg/easyjson/jwriter"
)

// DetachReason detach reason.
//
// See: ( -- none -- )
type DetachReason string

// String returns the DetachReason as string value.
func (t DetachReason) String() string {
	return string(t)
}

// DetachReason values.
const (
	DetachReasonTargetClosed         DetachReason = "target_closed"
	DetachReasonCanceledByUser       DetachReason = "canceled_by_user"
	DetachReasonReplacedWithDevtools DetachReason = "replaced_with_devtools"
	DetachReasonRenderProcessGone    DetachReason = "Render process gone."
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DetachReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DetachReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DetachReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch DetachReason(v) {
	case DetachReasonTargetClosed:
		*t = DetachReasonTargetClosed
	case DetachReasonCanceledByUser:
		*t = DetachReasonCanceledByUser
	case DetachReasonReplacedWithDevtools:
		*t = DetachReasonReplacedWithDevtools
	case DetachReasonRenderProcessGone:
		*t = DetachReasonRenderProcessGone

	default:
		in.AddError(fmt.Errorf("unknown DetachReason value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DetachReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
