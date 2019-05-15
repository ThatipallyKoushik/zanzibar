// Code generated by zanzibar
// @generated
// Checksum : hiSSI9tb3wASTCls+Pp53w==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

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

func easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall(in *jlexer.Lexer, out *SimpleService_Call_Result) {
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
		case "authErr":
			if in.IsNull() {
				in.Skip()
				out.AuthErr = nil
			} else {
				if out.AuthErr == nil {
					out.AuthErr = new(AuthErr)
				}
				easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(in, out.AuthErr)
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
func easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall(out *jwriter.Writer, in SimpleService_Call_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AuthErr != nil {
		const prefix string = ",\"authErr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(out, *in.AuthErr)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_Call_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_Call_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_Call_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_Call_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall(l, v)
}
func easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(in *jlexer.Lexer, out *AuthErr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var MessageSet bool
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
		case "message":
			out.Message = string(in.String())
			MessageSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !MessageSet {
		in.AddError(fmt.Errorf("key 'message' is required"))
	}
}
func easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(out *jwriter.Writer, in AuthErr) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}
func easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall1(in *jlexer.Lexer, out *SimpleService_Call_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var ArgSet bool
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
		case "arg":
			if in.IsNull() {
				in.Skip()
				out.Arg = nil
			} else {
				if out.Arg == nil {
					out.Arg = new(BazRequest)
				}
				easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz1(in, out.Arg)
			}
			ArgSet = true
		case "i64Optional":
			if in.IsNull() {
				in.Skip()
				out.I64Optional = nil
			} else {
				if out.I64Optional == nil {
					out.I64Optional = new(int64)
				}
				*out.I64Optional = int64(in.Int64())
			}
		case "testUUID":
			if in.IsNull() {
				in.Skip()
				out.TestUUID = nil
			} else {
				if out.TestUUID == nil {
					out.TestUUID = new(UUID)
				}
				*out.TestUUID = UUID(in.String())
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
	if !ArgSet {
		in.AddError(fmt.Errorf("key 'arg' is required"))
	}
}
func easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall1(out *jwriter.Writer, in SimpleService_Call_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"arg\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Arg == nil {
			out.RawString("null")
		} else {
			easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz1(out, *in.Arg)
		}
	}
	if in.I64Optional != nil {
		const prefix string = ",\"i64Optional\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.I64Optional))
	}
	if in.TestUUID != nil {
		const prefix string = ",\"testUUID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.TestUUID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_Call_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_Call_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_Call_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_Call_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceCall1(l, v)
}
func easyjsonA954e906DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz1(in *jlexer.Lexer, out *BazRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var B1Set bool
	var S2Set bool
	var I3Set bool
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
		case "b1":
			out.B1 = bool(in.Bool())
			B1Set = true
		case "s2":
			out.S2 = string(in.String())
			S2Set = true
		case "i3":
			out.I3 = int32(in.Int32())
			I3Set = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !B1Set {
		in.AddError(fmt.Errorf("key 'b1' is required"))
	}
	if !S2Set {
		in.AddError(fmt.Errorf("key 's2' is required"))
	}
	if !I3Set {
		in.AddError(fmt.Errorf("key 'i3' is required"))
	}
}
func easyjsonA954e906EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz1(out *jwriter.Writer, in BazRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"b1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.B1))
	}
	{
		const prefix string = ",\"s2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.S2))
	}
	{
		const prefix string = ",\"i3\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.I3))
	}
	out.RawByte('}')
}
