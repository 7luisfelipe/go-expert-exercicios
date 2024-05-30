// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: pedidos.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListarPedidosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListarPedidosRequest) Reset() {
	*x = ListarPedidosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pedidos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListarPedidosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListarPedidosRequest) ProtoMessage() {}

func (x *ListarPedidosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pedidos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListarPedidosRequest.ProtoReflect.Descriptor instead.
func (*ListarPedidosRequest) Descriptor() ([]byte, []int) {
	return file_pedidos_proto_rawDescGZIP(), []int{0}
}

type ListarPedidosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pedidos []*Pedido `protobuf:"bytes,1,rep,name=pedidos,proto3" json:"pedidos,omitempty"`
}

func (x *ListarPedidosResponse) Reset() {
	*x = ListarPedidosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pedidos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListarPedidosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListarPedidosResponse) ProtoMessage() {}

func (x *ListarPedidosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pedidos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListarPedidosResponse.ProtoReflect.Descriptor instead.
func (*ListarPedidosResponse) Descriptor() ([]byte, []int) {
	return file_pedidos_proto_rawDescGZIP(), []int{1}
}

func (x *ListarPedidosResponse) GetPedidos() []*Pedido {
	if x != nil {
		return x.Pedidos
	}
	return nil
}

type Pedido struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NumeroPedido  int32   `protobuf:"varint,2,opt,name=numeroPedido,proto3" json:"numeroPedido,omitempty"`
	NomeProduto   string  `protobuf:"bytes,3,opt,name=nomeProduto,proto3" json:"nomeProduto,omitempty"`
	Quantidade    int32   `protobuf:"varint,4,opt,name=quantidade,proto3" json:"quantidade,omitempty"`
	PrecoUnitario float32 `protobuf:"fixed32,5,opt,name=precoUnitario,proto3" json:"precoUnitario,omitempty"`
}

func (x *Pedido) Reset() {
	*x = Pedido{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pedidos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pedido) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pedido) ProtoMessage() {}

func (x *Pedido) ProtoReflect() protoreflect.Message {
	mi := &file_pedidos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pedido.ProtoReflect.Descriptor instead.
func (*Pedido) Descriptor() ([]byte, []int) {
	return file_pedidos_proto_rawDescGZIP(), []int{2}
}

func (x *Pedido) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pedido) GetNumeroPedido() int32 {
	if x != nil {
		return x.NumeroPedido
	}
	return 0
}

func (x *Pedido) GetNomeProduto() string {
	if x != nil {
		return x.NomeProduto
	}
	return ""
}

func (x *Pedido) GetQuantidade() int32 {
	if x != nil {
		return x.Quantidade
	}
	return 0
}

func (x *Pedido) GetPrecoUnitario() float32 {
	if x != nil {
		return x.PrecoUnitario
	}
	return 0
}

var File_pedidos_proto protoreflect.FileDescriptor

var file_pedidos_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x72, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x61,
	0x72, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x07, 0x70, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x52, 0x07, 0x70, 0x65, 0x64, 0x69,
	0x64, 0x6f, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x50, 0x65, 0x64, 0x69,
	0x64, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x6d, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x64, 0x61, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x55, 0x6e, 0x69,
	0x74, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x70, 0x72, 0x65,
	0x63, 0x6f, 0x55, 0x6e, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x32, 0x4f, 0x0a, 0x0d, 0x50, 0x65,
	0x64, 0x69, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x61, 0x72, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x73, 0x12, 0x15, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x61, 0x72, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x72, 0x50, 0x65, 0x64, 0x69,
	0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pedidos_proto_rawDescOnce sync.Once
	file_pedidos_proto_rawDescData = file_pedidos_proto_rawDesc
)

func file_pedidos_proto_rawDescGZIP() []byte {
	file_pedidos_proto_rawDescOnce.Do(func() {
		file_pedidos_proto_rawDescData = protoimpl.X.CompressGZIP(file_pedidos_proto_rawDescData)
	})
	return file_pedidos_proto_rawDescData
}

var file_pedidos_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pedidos_proto_goTypes = []interface{}{
	(*ListarPedidosRequest)(nil),  // 0: ListarPedidosRequest
	(*ListarPedidosResponse)(nil), // 1: ListarPedidosResponse
	(*Pedido)(nil),                // 2: Pedido
}
var file_pedidos_proto_depIdxs = []int32{
	2, // 0: ListarPedidosResponse.pedidos:type_name -> Pedido
	0, // 1: PedidoService.ListarPedidos:input_type -> ListarPedidosRequest
	1, // 2: PedidoService.ListarPedidos:output_type -> ListarPedidosResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pedidos_proto_init() }
func file_pedidos_proto_init() {
	if File_pedidos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pedidos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListarPedidosRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pedidos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListarPedidosResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pedidos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pedido); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pedidos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pedidos_proto_goTypes,
		DependencyIndexes: file_pedidos_proto_depIdxs,
		MessageInfos:      file_pedidos_proto_msgTypes,
	}.Build()
	File_pedidos_proto = out.File
	file_pedidos_proto_rawDesc = nil
	file_pedidos_proto_goTypes = nil
	file_pedidos_proto_depIdxs = nil
}
