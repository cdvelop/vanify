// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package headlessexperimental

import (
	json "encoding/json"
	easyjson "github.com/cdvelop/vanify/pkg/easyjson"
	jlexer "github.com/cdvelop/vanify/pkg/easyjson/jlexer"
	jwriter "github.com/cdvelop/vanify/pkg/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoHeadlessexperimental(in *jlexer.Lexer, out *ScreenshotParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "format":
			(out.Format).UnmarshalEasyJSON(in)
		case "quality":
			out.Quality = int64(in.Int64())
		case "optimizeForSpeed":
			out.OptimizeForSpeed = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoHeadlessexperimental(out *jwriter.Writer, in ScreenshotParams) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Format != "" {
		const prefix string = ",\"format\":"
		first = false
		out.RawString(prefix[1:])
		(in.Format).MarshalEasyJSON(out)
	}
	if in.Quality != 0 {
		const prefix string = ",\"quality\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Quality))
	}
	if in.OptimizeForSpeed {
		const prefix string = ",\"optimizeForSpeed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.OptimizeForSpeed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScreenshotParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoHeadlessexperimental(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScreenshotParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoHeadlessexperimental(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScreenshotParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoHeadlessexperimental(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScreenshotParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoHeadlessexperimental(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoHeadlessexperimental1(in *jlexer.Lexer, out *BeginFrameReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "hasDamage":
			out.HasDamage = bool(in.Bool())
		case "screenshotData":
			out.ScreenshotData = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoHeadlessexperimental1(out *jwriter.Writer, in BeginFrameReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.HasDamage {
		const prefix string = ",\"hasDamage\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.HasDamage))
	}
	if in.ScreenshotData != "" {
		const prefix string = ",\"screenshotData\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ScreenshotData))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BeginFrameReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoHeadlessexperimental1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BeginFrameReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoHeadlessexperimental1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BeginFrameReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoHeadlessexperimental1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BeginFrameReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoHeadlessexperimental1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoHeadlessexperimental2(in *jlexer.Lexer, out *BeginFrameParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "frameTimeTicks":
			out.FrameTimeTicks = float64(in.Float64())
		case "interval":
			out.Interval = float64(in.Float64())
		case "noDisplayUpdates":
			out.NoDisplayUpdates = bool(in.Bool())
		case "screenshot":
			if in.IsNull() {
				in.Skip()
				out.Screenshot = nil
			} else {
				if out.Screenshot == nil {
					out.Screenshot = new(ScreenshotParams)
				}
				(*out.Screenshot).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoHeadlessexperimental2(out *jwriter.Writer, in BeginFrameParams) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FrameTimeTicks != 0 {
		const prefix string = ",\"frameTimeTicks\":"
		first = false
		out.RawString(prefix[1:])
		out.Float64(float64(in.FrameTimeTicks))
	}
	if in.Interval != 0 {
		const prefix string = ",\"interval\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Interval))
	}
	if in.NoDisplayUpdates {
		const prefix string = ",\"noDisplayUpdates\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.NoDisplayUpdates))
	}
	if in.Screenshot != nil {
		const prefix string = ",\"screenshot\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Screenshot).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BeginFrameParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoHeadlessexperimental2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BeginFrameParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoHeadlessexperimental2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BeginFrameParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoHeadlessexperimental2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BeginFrameParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoHeadlessexperimental2(l, v)
}
