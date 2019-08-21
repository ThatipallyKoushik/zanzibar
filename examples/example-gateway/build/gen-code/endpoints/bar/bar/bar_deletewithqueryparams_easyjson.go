// Code generated by zanzibar
// @generated
// Checksum : w4tJ9F0//opMGSTzPlQTsw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson883c5671DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams(in *jlexer.Lexer, out *Bar_DeleteWithQueryParams_Result) {
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
		key := in.UnsafeString()
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
func easyjson883c5671EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams(out *jwriter.Writer, in Bar_DeleteWithQueryParams_Result) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_DeleteWithQueryParams_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson883c5671EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_DeleteWithQueryParams_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson883c5671EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_DeleteWithQueryParams_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson883c5671DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_DeleteWithQueryParams_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson883c5671DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams(l, v)
}
func easyjson883c5671DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams1(in *jlexer.Lexer, out *Bar_DeleteWithQueryParams_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var RequestSet bool
	var FilterSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(BarRequest)
				}
				easyjson883c5671DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(in, out.Request)
			}
			RequestSet = true
		case "filter":
			out.Filter = string(in.String())
			FilterSet = true
		case "count":
			if in.IsNull() {
				in.Skip()
				out.Count = nil
			} else {
				if out.Count == nil {
					out.Count = new(int32)
				}
				*out.Count = int32(in.Int32())
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
	if !RequestSet {
		in.AddError(fmt.Errorf("key 'request' is required"))
	}
	if !FilterSet {
		in.AddError(fmt.Errorf("key 'filter' is required"))
	}
}
func easyjson883c5671EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams1(out *jwriter.Writer, in Bar_DeleteWithQueryParams_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"request\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Request == nil {
			out.RawString("null")
		} else {
			easyjson883c5671EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(out, *in.Request)
		}
	}
	{
		const prefix string = ",\"filter\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Filter))
	}
	if in.Count != nil {
		const prefix string = ",\"count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.Count))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_DeleteWithQueryParams_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson883c5671EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_DeleteWithQueryParams_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson883c5671EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_DeleteWithQueryParams_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson883c5671DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_DeleteWithQueryParams_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson883c5671DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarDeleteWithQueryParams1(l, v)
}
func easyjson883c5671DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(in *jlexer.Lexer, out *BarRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var BoolFieldSet bool
	var BinaryFieldSet bool
	var TimestampSet bool
	var EnumFieldSet bool
	var LongFieldSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "boolField":
			out.BoolField = bool(in.Bool())
			BoolFieldSet = true
		case "binaryField":
			if in.IsNull() {
				in.Skip()
				out.BinaryField = nil
			} else {
				out.BinaryField = in.Bytes()
			}
			BinaryFieldSet = true
		case "timestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
			TimestampSet = true
		case "enumField":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EnumField).UnmarshalJSON(data))
			}
			EnumFieldSet = true
		case "longField":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LongField).UnmarshalJSON(data))
			}
			LongFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !BoolFieldSet {
		in.AddError(fmt.Errorf("key 'boolField' is required"))
	}
	if !BinaryFieldSet {
		in.AddError(fmt.Errorf("key 'binaryField' is required"))
	}
	if !TimestampSet {
		in.AddError(fmt.Errorf("key 'timestamp' is required"))
	}
	if !EnumFieldSet {
		in.AddError(fmt.Errorf("key 'enumField' is required"))
	}
	if !LongFieldSet {
		in.AddError(fmt.Errorf("key 'longField' is required"))
	}
}
func easyjson883c5671EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(out *jwriter.Writer, in BarRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"stringField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.StringField))
	}
	{
		const prefix string = ",\"boolField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.BoolField))
	}
	{
		const prefix string = ",\"binaryField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.BinaryField)
	}
	{
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"enumField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.EnumField).MarshalJSON())
	}
	{
		const prefix string = ",\"longField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LongField).MarshalJSON())
	}
	out.RawByte('}')
}
