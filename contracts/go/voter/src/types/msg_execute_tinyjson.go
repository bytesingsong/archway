// Code generated by tinyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

import (
	types "github.com/CosmWasm/cosmwasm-go/std/types"
	tinyjson "github.com/CosmWasm/tinyjson"
	jlexer "github.com/CosmWasm/tinyjson/jlexer"
	jwriter "github.com/CosmWasm/tinyjson/jwriter"
	custom "github.com/archway-network/voter/src/pkg/archway/custom"
)

// suppress unused package warning
var (
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ tinyjson.Marshaler
)

func tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes(in *jlexer.Lexer, out *VoteRequest) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "option":
			out.Option = string(in.String())
		case "vote":
			out.Vote = string(in.String())
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
func tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes(out *jwriter.Writer, in VoteRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"option\":"
		out.RawString(prefix)
		out.String(string(in.Option))
	}
	{
		const prefix string = ",\"vote\":"
		out.RawString(prefix)
		out.String(string(in.Vote))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VoteRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v VoteRequest) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VoteRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *VoteRequest) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes(l, v)
}
func tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes1(in *jlexer.Lexer, out *SendIBCVoteRequest) {
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
		case "channel_id":
			out.ChannelID = string(in.String())
		case "id":
			out.ID = uint64(in.Uint64())
		case "option":
			out.Option = string(in.String())
		case "vote":
			out.Vote = string(in.String())
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
func tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes1(out *jwriter.Writer, in SendIBCVoteRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"channel_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ChannelID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"option\":"
		out.RawString(prefix)
		out.String(string(in.Option))
	}
	{
		const prefix string = ",\"vote\":"
		out.RawString(prefix)
		out.String(string(in.Vote))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SendIBCVoteRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v SendIBCVoteRequest) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SendIBCVoteRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes1(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *SendIBCVoteRequest) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes1(l, v)
}
func tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes2(in *jlexer.Lexer, out *ReleaseResponse) {
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
		case "released_amount":
			if in.IsNull() {
				in.Skip()
				out.ReleasedAmount = nil
			} else {
				in.Delim('[')
				if out.ReleasedAmount == nil {
					if !in.IsDelim(']') {
						out.ReleasedAmount = make([]types.Coin, 0, 2)
					} else {
						out.ReleasedAmount = []types.Coin{}
					}
				} else {
					out.ReleasedAmount = (out.ReleasedAmount)[:0]
				}
				for !in.IsDelim(']') {
					var v1 types.Coin
					(v1).UnmarshalTinyJSON(in)
					out.ReleasedAmount = append(out.ReleasedAmount, v1)
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
func tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes2(out *jwriter.Writer, in ReleaseResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"released_amount\":"
		out.RawString(prefix[1:])
		if in.ReleasedAmount == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.ReleasedAmount {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalTinyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReleaseResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v ReleaseResponse) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReleaseResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes2(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *ReleaseResponse) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes2(l, v)
}
func tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes3(in *jlexer.Lexer, out *NewVotingResponse) {
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
		case "voting_id":
			out.VotingID = uint64(in.Uint64())
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
func tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes3(out *jwriter.Writer, in NewVotingResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"voting_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.VotingID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NewVotingResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v NewVotingResponse) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NewVotingResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes3(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *NewVotingResponse) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes3(l, v)
}
func tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes4(in *jlexer.Lexer, out *NewVotingRequest) {
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
		case "name":
			out.Name = string(in.String())
		case "vote_options":
			if in.IsNull() {
				in.Skip()
				out.VoteOptions = nil
			} else {
				in.Delim('[')
				if out.VoteOptions == nil {
					if !in.IsDelim(']') {
						out.VoteOptions = make([]string, 0, 4)
					} else {
						out.VoteOptions = []string{}
					}
				} else {
					out.VoteOptions = (out.VoteOptions)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.VoteOptions = append(out.VoteOptions, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "duration":
			out.Duration = uint64(in.Uint64())
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
func tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes4(out *jwriter.Writer, in NewVotingRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"vote_options\":"
		out.RawString(prefix)
		if in.VoteOptions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.VoteOptions {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"duration\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Duration))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NewVotingRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v NewVotingRequest) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NewVotingRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes4(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *NewVotingRequest) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes4(l, v)
}
func tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes5(in *jlexer.Lexer, out *MsgExecute) {
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
		case "release":
			if in.IsNull() {
				in.Skip()
				out.Release = nil
			} else {
				if out.Release == nil {
					out.Release = new(struct{})
				}
				tinyjson8d55ca79Decode(in, out.Release)
			}
		case "new_voting":
			if in.IsNull() {
				in.Skip()
				out.NewVoting = nil
			} else {
				if out.NewVoting == nil {
					out.NewVoting = new(NewVotingRequest)
				}
				(*out.NewVoting).UnmarshalTinyJSON(in)
			}
		case "vote":
			if in.IsNull() {
				in.Skip()
				out.Vote = nil
			} else {
				if out.Vote == nil {
					out.Vote = new(VoteRequest)
				}
				(*out.Vote).UnmarshalTinyJSON(in)
			}
		case "send_ibc_vote":
			if in.IsNull() {
				in.Skip()
				out.SendIBCVote = nil
			} else {
				if out.SendIBCVote == nil {
					out.SendIBCVote = new(SendIBCVoteRequest)
				}
				(*out.SendIBCVote).UnmarshalTinyJSON(in)
			}
		case "custom_custom":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CustomCustom).UnmarshalJSON(data))
			}
		case "custom_update_metadata":
			if in.IsNull() {
				in.Skip()
				out.CustomUpdateMetadata = nil
			} else {
				if out.CustomUpdateMetadata == nil {
					out.CustomUpdateMetadata = new(custom.UpdateContractMetadataRequest)
				}
				(*out.CustomUpdateMetadata).UnmarshalTinyJSON(in)
			}
		case "custom_withdraw_rewards":
			if in.IsNull() {
				in.Skip()
				out.CustomWithdrawRewards = nil
			} else {
				if out.CustomWithdrawRewards == nil {
					out.CustomWithdrawRewards = new(custom.WithdrawRewardsRequest)
				}
				(*out.CustomWithdrawRewards).UnmarshalTinyJSON(in)
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
func tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes5(out *jwriter.Writer, in MsgExecute) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Release != nil {
		const prefix string = ",\"release\":"
		first = false
		out.RawString(prefix[1:])
		tinyjson8d55ca79Encode(out, *in.Release)
	}
	if in.NewVoting != nil {
		const prefix string = ",\"new_voting\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.NewVoting).MarshalTinyJSON(out)
	}
	if in.Vote != nil {
		const prefix string = ",\"vote\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Vote).MarshalTinyJSON(out)
	}
	if in.SendIBCVote != nil {
		const prefix string = ",\"send_ibc_vote\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.SendIBCVote).MarshalTinyJSON(out)
	}
	if len(in.CustomCustom) != 0 {
		const prefix string = ",\"custom_custom\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CustomCustom).MarshalJSON())
	}
	if in.CustomUpdateMetadata != nil {
		const prefix string = ",\"custom_update_metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.CustomUpdateMetadata).MarshalTinyJSON(out)
	}
	if in.CustomWithdrawRewards != nil {
		const prefix string = ",\"custom_withdraw_rewards\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.CustomWithdrawRewards).MarshalTinyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MsgExecute) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v MsgExecute) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8d55ca79EncodeGithubComArchwayNetworkVoterSrcTypes5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MsgExecute) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes5(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *MsgExecute) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8d55ca79DecodeGithubComArchwayNetworkVoterSrcTypes5(l, v)
}
func tinyjson8d55ca79Decode(in *jlexer.Lexer, out *struct{}) {
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
func tinyjson8d55ca79Encode(out *jwriter.Writer, in struct{}) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}