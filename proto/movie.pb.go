// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/movie.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MoviesList struct {
	Movies []*MoviesList_Movie `protobuf:"bytes,1,rep,name=movies" json:"movies,omitempty"`
}

func (m *MoviesList) Reset()                    { *m = MoviesList{} }
func (m *MoviesList) String() string            { return proto1.CompactTextString(m) }
func (*MoviesList) ProtoMessage()               {}
func (*MoviesList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MoviesList) GetMovies() []*MoviesList_Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

type MoviesList_IMDBMeta struct {
	Genre      string  `protobuf:"bytes,1,opt,name=genre" json:"genre,omitempty"`
	MpaaRating string  `protobuf:"bytes,2,opt,name=mpaa_rating,json=mpaaRating" json:"mpaa_rating,omitempty"`
	Score      float32 `protobuf:"fixed32,3,opt,name=score" json:"score,omitempty"`
}

func (m *MoviesList_IMDBMeta) Reset()                    { *m = MoviesList_IMDBMeta{} }
func (m *MoviesList_IMDBMeta) String() string            { return proto1.CompactTextString(m) }
func (*MoviesList_IMDBMeta) ProtoMessage()               {}
func (*MoviesList_IMDBMeta) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *MoviesList_IMDBMeta) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func (m *MoviesList_IMDBMeta) GetMpaaRating() string {
	if m != nil {
		return m.MpaaRating
	}
	return ""
}

func (m *MoviesList_IMDBMeta) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type MoviesList_RottenTomatoesMeta struct {
	TomatoScore        int32  `protobuf:"varint,1,opt,name=tomato_score,json=tomatoScore" json:"tomato_score,omitempty"`
	PopcornScore       int32  `protobuf:"varint,2,opt,name=popcorn_score,json=popcornScore" json:"popcorn_score,omitempty"`
	TheaterReleaseDate string `protobuf:"bytes,3,opt,name=theater_release_date,json=theaterReleaseDate" json:"theater_release_date,omitempty"`
	MpaaRating         string `protobuf:"bytes,4,opt,name=mpaa_rating,json=mpaaRating" json:"mpaa_rating,omitempty"`
	Synopsis           string `protobuf:"bytes,5,opt,name=synopsis" json:"synopsis,omitempty"`
	SynopsisType       string `protobuf:"bytes,6,opt,name=synopsis_type,json=synopsisType" json:"synopsis_type,omitempty"`
	Runtime            string `protobuf:"bytes,7,opt,name=runtime" json:"runtime,omitempty"`
}

func (m *MoviesList_RottenTomatoesMeta) Reset()         { *m = MoviesList_RottenTomatoesMeta{} }
func (m *MoviesList_RottenTomatoesMeta) String() string { return proto1.CompactTextString(m) }
func (*MoviesList_RottenTomatoesMeta) ProtoMessage()    {}
func (*MoviesList_RottenTomatoesMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 1}
}

func (m *MoviesList_RottenTomatoesMeta) GetTomatoScore() int32 {
	if m != nil {
		return m.TomatoScore
	}
	return 0
}

func (m *MoviesList_RottenTomatoesMeta) GetPopcornScore() int32 {
	if m != nil {
		return m.PopcornScore
	}
	return 0
}

func (m *MoviesList_RottenTomatoesMeta) GetTheaterReleaseDate() string {
	if m != nil {
		return m.TheaterReleaseDate
	}
	return ""
}

func (m *MoviesList_RottenTomatoesMeta) GetMpaaRating() string {
	if m != nil {
		return m.MpaaRating
	}
	return ""
}

func (m *MoviesList_RottenTomatoesMeta) GetSynopsis() string {
	if m != nil {
		return m.Synopsis
	}
	return ""
}

func (m *MoviesList_RottenTomatoesMeta) GetSynopsisType() string {
	if m != nil {
		return m.SynopsisType
	}
	return ""
}

func (m *MoviesList_RottenTomatoesMeta) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

type MoviesList_Movie struct {
	Title              string                         `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	ImdbMeta           *MoviesList_IMDBMeta           `protobuf:"bytes,2,opt,name=imdb_meta,json=imdbMeta" json:"imdb_meta,omitempty"`
	RottenTomatoesMeta *MoviesList_RottenTomatoesMeta `protobuf:"bytes,3,opt,name=rotten_tomatoes_meta,json=rottenTomatoesMeta" json:"rotten_tomatoes_meta,omitempty"`
}

func (m *MoviesList_Movie) Reset()                    { *m = MoviesList_Movie{} }
func (m *MoviesList_Movie) String() string            { return proto1.CompactTextString(m) }
func (*MoviesList_Movie) ProtoMessage()               {}
func (*MoviesList_Movie) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 2} }

func (m *MoviesList_Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MoviesList_Movie) GetImdbMeta() *MoviesList_IMDBMeta {
	if m != nil {
		return m.ImdbMeta
	}
	return nil
}

func (m *MoviesList_Movie) GetRottenTomatoesMeta() *MoviesList_RottenTomatoesMeta {
	if m != nil {
		return m.RottenTomatoesMeta
	}
	return nil
}

type PostMoviesResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *PostMoviesResponse) Reset()                    { *m = PostMoviesResponse{} }
func (m *PostMoviesResponse) String() string            { return proto1.CompactTextString(m) }
func (*PostMoviesResponse) ProtoMessage()               {}
func (*PostMoviesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PostMoviesResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type Search struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Search) Reset()                    { *m = Search{} }
func (m *Search) String() string            { return proto1.CompactTextString(m) }
func (*Search) ProtoMessage()               {}
func (*Search) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Search) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto1.RegisterType((*MoviesList)(nil), "proto.MoviesList")
	proto1.RegisterType((*MoviesList_IMDBMeta)(nil), "proto.MoviesList.IMDBMeta")
	proto1.RegisterType((*MoviesList_RottenTomatoesMeta)(nil), "proto.MoviesList.RottenTomatoesMeta")
	proto1.RegisterType((*MoviesList_Movie)(nil), "proto.MoviesList.Movie")
	proto1.RegisterType((*PostMoviesResponse)(nil), "proto.PostMoviesResponse")
	proto1.RegisterType((*Search)(nil), "proto.Search")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MovieService service

type MovieServiceClient interface {
	Autocomplete(ctx context.Context, in *Search, opts ...grpc.CallOption) (*MoviesList, error)
	BulkIndex(ctx context.Context, in *MoviesList, opts ...grpc.CallOption) (*PostMoviesResponse, error)
	Get(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*MoviesList, error)
}

type movieServiceClient struct {
	cc *grpc.ClientConn
}

func NewMovieServiceClient(cc *grpc.ClientConn) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) Autocomplete(ctx context.Context, in *Search, opts ...grpc.CallOption) (*MoviesList, error) {
	out := new(MoviesList)
	err := grpc.Invoke(ctx, "/proto.MovieService/Autocomplete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) BulkIndex(ctx context.Context, in *MoviesList, opts ...grpc.CallOption) (*PostMoviesResponse, error) {
	out := new(PostMoviesResponse)
	err := grpc.Invoke(ctx, "/proto.MovieService/BulkIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) Get(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*MoviesList, error) {
	out := new(MoviesList)
	err := grpc.Invoke(ctx, "/proto.MovieService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieService service

type MovieServiceServer interface {
	Autocomplete(context.Context, *Search) (*MoviesList, error)
	BulkIndex(context.Context, *MoviesList) (*PostMoviesResponse, error)
	Get(context.Context, *google_protobuf1.Empty) (*MoviesList, error)
}

func RegisterMovieServiceServer(s *grpc.Server, srv MovieServiceServer) {
	s.RegisterService(&_MovieService_serviceDesc, srv)
}

func _MovieService_Autocomplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).Autocomplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/Autocomplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).Autocomplete(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_BulkIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoviesList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).BulkIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/BulkIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).BulkIndex(ctx, req.(*MoviesList))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).Get(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Autocomplete",
			Handler:    _MovieService_Autocomplete_Handler,
		},
		{
			MethodName: "BulkIndex",
			Handler:    _MovieService_BulkIndex_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MovieService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/movie.proto",
}

func init() { proto1.RegisterFile("proto/movie.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x26, 0x4d, 0x26, 0xae, 0x50, 0x57, 0x15, 0x18, 0xb7, 0x88, 0xe2, 0x72, 0xa8,
	0x38, 0xd8, 0xa8, 0x1c, 0x90, 0xb8, 0xb5, 0x2a, 0x42, 0x95, 0x88, 0x84, 0x36, 0x05, 0x8e, 0xd6,
	0xc6, 0x19, 0x52, 0x8b, 0xd8, 0x6b, 0xed, 0x4e, 0xaa, 0x46, 0x88, 0x0b, 0x47, 0xae, 0xfc, 0x07,
	0x3f, 0x03, 0x9f, 0xc0, 0x3f, 0x70, 0x45, 0x9e, 0xb5, 0x0b, 0x92, 0x39, 0x79, 0xde, 0x9b, 0x37,
	0xcf, 0x33, 0xb3, 0x1a, 0xd8, 0xad, 0x8c, 0x26, 0x9d, 0x14, 0xfa, 0x3a, 0xc7, 0x98, 0x63, 0x31,
	0xe0, 0x4f, 0x78, 0xb0, 0xd4, 0x7a, 0xb9, 0xc2, 0x44, 0x55, 0x79, 0xa2, 0xca, 0x52, 0x93, 0xa2,
	0x5c, 0x97, 0xd6, 0x89, 0xc2, 0xfd, 0x26, 0xcb, 0x68, 0xbe, 0xfe, 0x90, 0x60, 0x51, 0xd1, 0xc6,
	0x25, 0xa3, 0x9f, 0x5b, 0x00, 0xd3, 0xda, 0xd1, 0xbe, 0xce, 0x2d, 0x89, 0x04, 0x86, 0xec, 0x6f,
	0x03, 0xef, 0xb0, 0x7f, 0x3c, 0x39, 0xb9, 0xe7, 0x64, 0xf1, 0x5f, 0x89, 0x0b, 0x65, 0x23, 0x0b,
	0xdf, 0xc3, 0xe8, 0x62, 0x7a, 0x7e, 0x36, 0x45, 0x52, 0x62, 0x0f, 0x06, 0x4b, 0x2c, 0x0d, 0x06,
	0xde, 0xa1, 0x77, 0x3c, 0x96, 0x0e, 0x88, 0x87, 0x30, 0x29, 0x2a, 0xa5, 0x52, 0xa3, 0x28, 0x2f,
	0x97, 0x41, 0x8f, 0x73, 0x50, 0x53, 0x92, 0x99, 0xba, 0xcc, 0x66, 0xda, 0x60, 0xd0, 0x3f, 0xf4,
	0x8e, 0x7b, 0xd2, 0x81, 0xf0, 0x6b, 0x0f, 0x84, 0xd4, 0x44, 0x58, 0x5e, 0xea, 0x42, 0x91, 0x46,
	0xcb, 0xff, 0x78, 0x04, 0x3e, 0x31, 0x4e, 0x5d, 0x4d, 0xfd, 0xab, 0x81, 0x9c, 0x38, 0x6e, 0x56,
	0x53, 0xe2, 0x08, 0x76, 0x2a, 0x5d, 0x65, 0xda, 0x94, 0x8d, 0xa6, 0xc7, 0x1a, 0xbf, 0x21, 0x9d,
	0xe8, 0x29, 0xec, 0xd1, 0x15, 0x2a, 0x42, 0x93, 0x1a, 0x5c, 0xa1, 0xb2, 0x98, 0x2e, 0x14, 0xb9,
	0x1e, 0xc6, 0x52, 0x34, 0x39, 0xe9, 0x52, 0xe7, 0x8a, 0x3a, 0x73, 0x6c, 0x75, 0xe6, 0x08, 0x61,
	0x64, 0x37, 0xa5, 0xae, 0x6c, 0x6e, 0x83, 0x01, 0x67, 0x6f, 0x71, 0xdd, 0x53, 0x1b, 0xa7, 0xb4,
	0xa9, 0x30, 0x18, 0xb2, 0xc0, 0x6f, 0xc9, 0xcb, 0x4d, 0x85, 0x22, 0x80, 0x6d, 0xb3, 0x2e, 0x29,
	0x2f, 0x30, 0xd8, 0xe6, 0x74, 0x0b, 0xc3, 0xef, 0x1e, 0x0c, 0x78, 0xef, 0xf5, 0xb2, 0x28, 0xa7,
	0xd5, 0xed, 0x8e, 0x19, 0x88, 0xe7, 0x30, 0xce, 0x8b, 0xc5, 0x3c, 0x2d, 0x90, 0x14, 0x8f, 0x3b,
	0x39, 0x09, 0xbb, 0x2f, 0xd7, 0x3e, 0x94, 0x1c, 0xd5, 0x62, 0x5e, 0xe7, 0x3b, 0xd8, 0x33, 0xbc,
	0xe4, 0x94, 0x9a, 0x2d, 0x3b, 0x8f, 0x3e, 0x7b, 0x3c, 0xee, 0x7a, 0x74, 0x9f, 0x44, 0x0a, 0xd3,
	0xe1, 0xa2, 0x18, 0xc4, 0x1b, 0x6d, 0xc9, 0x15, 0x4a, 0xb4, 0x95, 0x2e, 0x2d, 0x0f, 0x68, 0xd7,
	0x59, 0x86, 0xd6, 0x72, 0xfb, 0x23, 0xd9, 0xc2, 0xe8, 0x00, 0x86, 0x33, 0x54, 0x26, 0xbb, 0x12,
	0x02, 0xb6, 0x08, 0x6f, 0xa8, 0x99, 0x8f, 0xe3, 0x93, 0xdf, 0x1e, 0xf8, 0x6c, 0x35, 0x43, 0x73,
	0x9d, 0x67, 0x28, 0xde, 0x82, 0x7f, 0xba, 0x26, 0x9d, 0xe9, 0xa2, 0x5a, 0x21, 0xa1, 0xd8, 0x69,
	0x1a, 0x75, 0x1e, 0xe1, 0x6e, 0xa7, 0xef, 0xe8, 0xe8, 0xcb, 0x8f, 0x5f, 0xdf, 0x7a, 0x0f, 0xc4,
	0xbe, 0xbb, 0x1f, 0x9b, 0xa8, 0x7f, 0xea, 0x93, 0x4f, 0xf5, 0x6f, 0x3e, 0x8b, 0x29, 0x8c, 0xcf,
	0xd6, 0xab, 0x8f, 0x17, 0xe5, 0x02, 0x6f, 0x44, 0xd7, 0x24, 0xbc, 0xdf, 0x50, 0xdd, 0xd1, 0x22,
	0xc1, 0xfe, 0x7e, 0xb4, 0xdd, 0xf8, 0xbf, 0xf0, 0x9e, 0x88, 0x53, 0xe8, 0xbf, 0x42, 0x12, 0x77,
	0x63, 0x77, 0x80, 0x71, 0x7b, 0x80, 0xf1, 0xcb, 0xfa, 0x00, 0xff, 0xd7, 0xe5, 0x1d, 0x76, 0x19,
	0x8b, 0xd6, 0x65, 0x3e, 0x64, 0xc9, 0xb3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0xa6, 0xc3,
	0x1d, 0xfc, 0x03, 0x00, 0x00,
}
