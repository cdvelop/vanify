// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package systeminfo

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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo(in *jlexer.Lexer, out *VideoEncodeAcceleratorCapability) {
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
		case "profile":
			out.Profile = string(in.String())
		case "maxResolution":
			if in.IsNull() {
				in.Skip()
				out.MaxResolution = nil
			} else {
				if out.MaxResolution == nil {
					out.MaxResolution = new(Size)
				}
				(*out.MaxResolution).UnmarshalEasyJSON(in)
			}
		case "maxFramerateNumerator":
			out.MaxFramerateNumerator = int64(in.Int64())
		case "maxFramerateDenominator":
			out.MaxFramerateDenominator = int64(in.Int64())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo(out *jwriter.Writer, in VideoEncodeAcceleratorCapability) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"profile\":"
		out.RawString(prefix[1:])
		out.String(string(in.Profile))
	}
	{
		const prefix string = ",\"maxResolution\":"
		out.RawString(prefix)
		if in.MaxResolution == nil {
			out.RawString("null")
		} else {
			(*in.MaxResolution).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"maxFramerateNumerator\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxFramerateNumerator))
	}
	{
		const prefix string = ",\"maxFramerateDenominator\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxFramerateDenominator))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VideoEncodeAcceleratorCapability) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VideoEncodeAcceleratorCapability) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VideoEncodeAcceleratorCapability) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VideoEncodeAcceleratorCapability) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo1(in *jlexer.Lexer, out *VideoDecodeAcceleratorCapability) {
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
		case "profile":
			out.Profile = string(in.String())
		case "maxResolution":
			if in.IsNull() {
				in.Skip()
				out.MaxResolution = nil
			} else {
				if out.MaxResolution == nil {
					out.MaxResolution = new(Size)
				}
				(*out.MaxResolution).UnmarshalEasyJSON(in)
			}
		case "minResolution":
			if in.IsNull() {
				in.Skip()
				out.MinResolution = nil
			} else {
				if out.MinResolution == nil {
					out.MinResolution = new(Size)
				}
				(*out.MinResolution).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo1(out *jwriter.Writer, in VideoDecodeAcceleratorCapability) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"profile\":"
		out.RawString(prefix[1:])
		out.String(string(in.Profile))
	}
	{
		const prefix string = ",\"maxResolution\":"
		out.RawString(prefix)
		if in.MaxResolution == nil {
			out.RawString("null")
		} else {
			(*in.MaxResolution).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"minResolution\":"
		out.RawString(prefix)
		if in.MinResolution == nil {
			out.RawString("null")
		} else {
			(*in.MinResolution).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VideoDecodeAcceleratorCapability) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VideoDecodeAcceleratorCapability) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VideoDecodeAcceleratorCapability) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VideoDecodeAcceleratorCapability) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo2(in *jlexer.Lexer, out *Size) {
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
		case "width":
			out.Width = int64(in.Int64())
		case "height":
			out.Height = int64(in.Int64())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo2(out *jwriter.Writer, in Size) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"width\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Width))
	}
	{
		const prefix string = ",\"height\":"
		out.RawString(prefix)
		out.Int64(int64(in.Height))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Size) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Size) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Size) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Size) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo3(in *jlexer.Lexer, out *ProcessInfo) {
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
		case "type":
			out.Type = string(in.String())
		case "id":
			out.ID = int64(in.Int64())
		case "cpuTime":
			out.CPUTime = float64(in.Float64())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo3(out *jwriter.Writer, in ProcessInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"cpuTime\":"
		out.RawString(prefix)
		out.Float64(float64(in.CPUTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProcessInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProcessInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProcessInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProcessInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo4(in *jlexer.Lexer, out *ImageDecodeAcceleratorCapability) {
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
		case "imageType":
			(out.ImageType).UnmarshalEasyJSON(in)
		case "maxDimensions":
			if in.IsNull() {
				in.Skip()
				out.MaxDimensions = nil
			} else {
				if out.MaxDimensions == nil {
					out.MaxDimensions = new(Size)
				}
				(*out.MaxDimensions).UnmarshalEasyJSON(in)
			}
		case "minDimensions":
			if in.IsNull() {
				in.Skip()
				out.MinDimensions = nil
			} else {
				if out.MinDimensions == nil {
					out.MinDimensions = new(Size)
				}
				(*out.MinDimensions).UnmarshalEasyJSON(in)
			}
		case "subsamplings":
			if in.IsNull() {
				in.Skip()
				out.Subsamplings = nil
			} else {
				in.Delim('[')
				if out.Subsamplings == nil {
					if !in.IsDelim(']') {
						out.Subsamplings = make([]SubsamplingFormat, 0, 4)
					} else {
						out.Subsamplings = []SubsamplingFormat{}
					}
				} else {
					out.Subsamplings = (out.Subsamplings)[:0]
				}
				for !in.IsDelim(']') {
					var v1 SubsamplingFormat
					(v1).UnmarshalEasyJSON(in)
					out.Subsamplings = append(out.Subsamplings, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo4(out *jwriter.Writer, in ImageDecodeAcceleratorCapability) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"imageType\":"
		out.RawString(prefix[1:])
		(in.ImageType).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"maxDimensions\":"
		out.RawString(prefix)
		if in.MaxDimensions == nil {
			out.RawString("null")
		} else {
			(*in.MaxDimensions).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"minDimensions\":"
		out.RawString(prefix)
		if in.MinDimensions == nil {
			out.RawString("null")
		} else {
			(*in.MinDimensions).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"subsamplings\":"
		out.RawString(prefix)
		if in.Subsamplings == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Subsamplings {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ImageDecodeAcceleratorCapability) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ImageDecodeAcceleratorCapability) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ImageDecodeAcceleratorCapability) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ImageDecodeAcceleratorCapability) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo5(in *jlexer.Lexer, out *GetProcessInfoReturns) {
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
		case "processInfo":
			if in.IsNull() {
				in.Skip()
				out.ProcessInfo = nil
			} else {
				in.Delim('[')
				if out.ProcessInfo == nil {
					if !in.IsDelim(']') {
						out.ProcessInfo = make([]*ProcessInfo, 0, 8)
					} else {
						out.ProcessInfo = []*ProcessInfo{}
					}
				} else {
					out.ProcessInfo = (out.ProcessInfo)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *ProcessInfo
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(ProcessInfo)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.ProcessInfo = append(out.ProcessInfo, v4)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo5(out *jwriter.Writer, in GetProcessInfoReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.ProcessInfo) != 0 {
		const prefix string = ",\"processInfo\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v5, v6 := range in.ProcessInfo {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetProcessInfoReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetProcessInfoReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetProcessInfoReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetProcessInfoReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo6(in *jlexer.Lexer, out *GetProcessInfoParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo6(out *jwriter.Writer, in GetProcessInfoParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetProcessInfoParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetProcessInfoParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetProcessInfoParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetProcessInfoParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo7(in *jlexer.Lexer, out *GetInfoReturns) {
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
		case "gpu":
			if in.IsNull() {
				in.Skip()
				out.Gpu = nil
			} else {
				if out.Gpu == nil {
					out.Gpu = new(GPUInfo)
				}
				(*out.Gpu).UnmarshalEasyJSON(in)
			}
		case "modelName":
			out.ModelName = string(in.String())
		case "modelVersion":
			out.ModelVersion = string(in.String())
		case "commandLine":
			out.CommandLine = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo7(out *jwriter.Writer, in GetInfoReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Gpu != nil {
		const prefix string = ",\"gpu\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Gpu).MarshalEasyJSON(out)
	}
	if in.ModelName != "" {
		const prefix string = ",\"modelName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ModelName))
	}
	if in.ModelVersion != "" {
		const prefix string = ",\"modelVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ModelVersion))
	}
	if in.CommandLine != "" {
		const prefix string = ",\"commandLine\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CommandLine))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetInfoReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetInfoReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetInfoReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetInfoReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo7(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo8(in *jlexer.Lexer, out *GetInfoParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo8(out *jwriter.Writer, in GetInfoParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetInfoParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetInfoParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetInfoParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetInfoParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo8(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo9(in *jlexer.Lexer, out *GetFeatureStateReturns) {
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
		case "featureEnabled":
			out.FeatureEnabled = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo9(out *jwriter.Writer, in GetFeatureStateReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FeatureEnabled {
		const prefix string = ",\"featureEnabled\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.FeatureEnabled))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFeatureStateReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFeatureStateReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFeatureStateReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFeatureStateReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo9(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo10(in *jlexer.Lexer, out *GetFeatureStateParams) {
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
		case "featureState":
			out.FeatureState = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo10(out *jwriter.Writer, in GetFeatureStateParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"featureState\":"
		out.RawString(prefix[1:])
		out.String(string(in.FeatureState))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFeatureStateParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFeatureStateParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFeatureStateParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFeatureStateParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo10(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo11(in *jlexer.Lexer, out *GPUInfo) {
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
		case "devices":
			if in.IsNull() {
				in.Skip()
				out.Devices = nil
			} else {
				in.Delim('[')
				if out.Devices == nil {
					if !in.IsDelim(']') {
						out.Devices = make([]*GPUDevice, 0, 8)
					} else {
						out.Devices = []*GPUDevice{}
					}
				} else {
					out.Devices = (out.Devices)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *GPUDevice
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(GPUDevice)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.Devices = append(out.Devices, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "auxAttributes":
			(out.AuxAttributes).UnmarshalEasyJSON(in)
		case "featureStatus":
			(out.FeatureStatus).UnmarshalEasyJSON(in)
		case "driverBugWorkarounds":
			if in.IsNull() {
				in.Skip()
				out.DriverBugWorkarounds = nil
			} else {
				in.Delim('[')
				if out.DriverBugWorkarounds == nil {
					if !in.IsDelim(']') {
						out.DriverBugWorkarounds = make([]string, 0, 4)
					} else {
						out.DriverBugWorkarounds = []string{}
					}
				} else {
					out.DriverBugWorkarounds = (out.DriverBugWorkarounds)[:0]
				}
				for !in.IsDelim(']') {
					var v8 string
					v8 = string(in.String())
					out.DriverBugWorkarounds = append(out.DriverBugWorkarounds, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "videoDecoding":
			if in.IsNull() {
				in.Skip()
				out.VideoDecoding = nil
			} else {
				in.Delim('[')
				if out.VideoDecoding == nil {
					if !in.IsDelim(']') {
						out.VideoDecoding = make([]*VideoDecodeAcceleratorCapability, 0, 8)
					} else {
						out.VideoDecoding = []*VideoDecodeAcceleratorCapability{}
					}
				} else {
					out.VideoDecoding = (out.VideoDecoding)[:0]
				}
				for !in.IsDelim(']') {
					var v9 *VideoDecodeAcceleratorCapability
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(VideoDecodeAcceleratorCapability)
						}
						(*v9).UnmarshalEasyJSON(in)
					}
					out.VideoDecoding = append(out.VideoDecoding, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "videoEncoding":
			if in.IsNull() {
				in.Skip()
				out.VideoEncoding = nil
			} else {
				in.Delim('[')
				if out.VideoEncoding == nil {
					if !in.IsDelim(']') {
						out.VideoEncoding = make([]*VideoEncodeAcceleratorCapability, 0, 8)
					} else {
						out.VideoEncoding = []*VideoEncodeAcceleratorCapability{}
					}
				} else {
					out.VideoEncoding = (out.VideoEncoding)[:0]
				}
				for !in.IsDelim(']') {
					var v10 *VideoEncodeAcceleratorCapability
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(VideoEncodeAcceleratorCapability)
						}
						(*v10).UnmarshalEasyJSON(in)
					}
					out.VideoEncoding = append(out.VideoEncoding, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "imageDecoding":
			if in.IsNull() {
				in.Skip()
				out.ImageDecoding = nil
			} else {
				in.Delim('[')
				if out.ImageDecoding == nil {
					if !in.IsDelim(']') {
						out.ImageDecoding = make([]*ImageDecodeAcceleratorCapability, 0, 8)
					} else {
						out.ImageDecoding = []*ImageDecodeAcceleratorCapability{}
					}
				} else {
					out.ImageDecoding = (out.ImageDecoding)[:0]
				}
				for !in.IsDelim(']') {
					var v11 *ImageDecodeAcceleratorCapability
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(ImageDecodeAcceleratorCapability)
						}
						(*v11).UnmarshalEasyJSON(in)
					}
					out.ImageDecoding = append(out.ImageDecoding, v11)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo11(out *jwriter.Writer, in GPUInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"devices\":"
		out.RawString(prefix[1:])
		if in.Devices == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Devices {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					(*v13).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if (in.AuxAttributes).IsDefined() {
		const prefix string = ",\"auxAttributes\":"
		out.RawString(prefix)
		(in.AuxAttributes).MarshalEasyJSON(out)
	}
	if (in.FeatureStatus).IsDefined() {
		const prefix string = ",\"featureStatus\":"
		out.RawString(prefix)
		(in.FeatureStatus).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"driverBugWorkarounds\":"
		out.RawString(prefix)
		if in.DriverBugWorkarounds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.DriverBugWorkarounds {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"videoDecoding\":"
		out.RawString(prefix)
		if in.VideoDecoding == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.VideoDecoding {
				if v16 > 0 {
					out.RawByte(',')
				}
				if v17 == nil {
					out.RawString("null")
				} else {
					(*v17).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"videoEncoding\":"
		out.RawString(prefix)
		if in.VideoEncoding == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.VideoEncoding {
				if v18 > 0 {
					out.RawByte(',')
				}
				if v19 == nil {
					out.RawString("null")
				} else {
					(*v19).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"imageDecoding\":"
		out.RawString(prefix)
		if in.ImageDecoding == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.ImageDecoding {
				if v20 > 0 {
					out.RawByte(',')
				}
				if v21 == nil {
					out.RawString("null")
				} else {
					(*v21).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GPUInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GPUInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GPUInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GPUInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo11(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo12(in *jlexer.Lexer, out *GPUDevice) {
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
		case "vendorId":
			out.VendorID = float64(in.Float64())
		case "deviceId":
			out.DeviceID = float64(in.Float64())
		case "subSysId":
			out.SubSysID = float64(in.Float64())
		case "revision":
			out.Revision = float64(in.Float64())
		case "vendorString":
			out.VendorString = string(in.String())
		case "deviceString":
			out.DeviceString = string(in.String())
		case "driverVendor":
			out.DriverVendor = string(in.String())
		case "driverVersion":
			out.DriverVersion = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo12(out *jwriter.Writer, in GPUDevice) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"vendorId\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.VendorID))
	}
	{
		const prefix string = ",\"deviceId\":"
		out.RawString(prefix)
		out.Float64(float64(in.DeviceID))
	}
	if in.SubSysID != 0 {
		const prefix string = ",\"subSysId\":"
		out.RawString(prefix)
		out.Float64(float64(in.SubSysID))
	}
	if in.Revision != 0 {
		const prefix string = ",\"revision\":"
		out.RawString(prefix)
		out.Float64(float64(in.Revision))
	}
	{
		const prefix string = ",\"vendorString\":"
		out.RawString(prefix)
		out.String(string(in.VendorString))
	}
	{
		const prefix string = ",\"deviceString\":"
		out.RawString(prefix)
		out.String(string(in.DeviceString))
	}
	{
		const prefix string = ",\"driverVendor\":"
		out.RawString(prefix)
		out.String(string(in.DriverVendor))
	}
	{
		const prefix string = ",\"driverVersion\":"
		out.RawString(prefix)
		out.String(string(in.DriverVersion))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GPUDevice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GPUDevice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSysteminfo12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GPUDevice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GPUDevice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSysteminfo12(l, v)
}
