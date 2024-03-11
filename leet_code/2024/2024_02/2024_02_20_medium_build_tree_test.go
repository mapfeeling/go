package _024_02

import (
	"encoding/json"
	"fmt"
	"testing"

	"git.n.xiaomi.com/mitv-media/define/pb"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

// 从前序与中序遍历构造二叉树
/*
给定两个整数数组 preorder 和 inorder ,其中 preorder 是二叉树的先序遍历,inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点
*/

type Tree20240220 struct {
	val       int
	LeftNode  *Tree20240220
	RightNode *Tree20240220
}

//func buildTree(preorder []int, inorder []int) *Tree20240220 {
//
//}

func NLR(root *Tree20240220) {
	if root == nil {
		return
	}
	fmt.Println("当前节点:", root.val)
	NLR(root.LeftNode)
	NLR(root.RightNode)
}

func LNR(root *Tree20240220) {
	if root == nil {
		return
	}
	LNR(root.LeftNode)
	fmt.Println("当前节点:", root.val)
	LNR(root.RightNode)
}

func LRN(root *Tree20240220) {
	if root == nil {
		return
	}
	LRN(root.LeftNode)
	LRN(root.RightNode)
	fmt.Println("当前节点:", root.val)
}

type BlockNew struct {
	*pb.Block
}

func Test20240220(t *testing.T) {
	qq()
	a()
	c()
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		var m = map[string]interface{}{
			"a": 1,
			"b": 2.0,
			"c": false,
		}

		var b = &BlockNew{
			&pb.Block{},
		}

		var p = &pb.Poster{}

		if myStruct, err := structpb.NewStruct(m); err == nil {
			fmt.Println("<===444", myStruct)
			fmt.Println("<===444", myStruct.AsMap())
			b.Filters = myStruct
		}

		//p.Id, _ = pbany(1)
		//p.Id, _ = interfaceToAny(1)
		p.Id, _ = d(false)

		fmt.Println(b)
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
			"msg":     b,
			"msg_2":   p,
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}

func pbany(v interface{}) (*anypb.Any, error) {
	pv, ok := v.(proto.Message)
	if !ok {
		return &anypb.Any{}, fmt.Errorf("%v is not proto.Message", pv)
	}
	return anypb.New(pv)
}

func ConvertInterfaceToAny(v interface{}) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	bytes, _ := sonic.Marshal(v)
	bytesValue := &wrappers.BytesValue{
		Value: bytes,
	}

	err := anypb.MarshalFrom(anyValue, bytesValue, proto.MarshalOptions{})
	return anyValue, err
}

func a() {
	var res protoregistry.MessageTypeResolver = protoregistry.GlobalTypes
	typeUrl := "type.googleapis.com/google.protobuf.StringValue"
	fmt.Println(protoreflect.FullName(typeUrl))
	msgType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName("google.protobuf.StringValue"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msgType)

	msgType, err = res.FindMessageByURL(typeUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msgType)
}

func interfaceToAny(v interface{}) (*anypb.Any, error) {
	bytes, err := sonic.Marshal(v)
	if err != nil {
		println("error json.Marshal interfaceToAny")
		return nil, err
	}
	m := pb.Bytes{B: bytes}
	return anypb.New(&m)
}

type Dog struct {
	Id   int
	Name string
}

func c() {
	data := map[string]interface{}{
		"id":   1,
		"name": "kitty",
	}

	s, _ := structpb.NewValue(data)
	newData := s.GetStructValue()

	m := pb.Poster{Id: structpb.NewStructValue(newData)}

	var dog Dog
	unmarshal(m.GetId(), &dog)
	fmt.Println(dog)

	var mm map[string]interface{}
	unmarshal(m.GetId(), &mm)
	fmt.Println(mm)

}

func unmarshal(p *structpb.Value, o interface{}) error {
	byt, _ := p.MarshalJSON()
	return json.Unmarshal(byt, o)
}

func d(o interface{}) (*structpb.Value, error) {
	s, err := structpb.NewValue(o)
	return s, err
}

//*structpb.Value

func qq() {
	fmt.Println("i am ") // 1
	var p = &pb.Poster{
		Url: "",
		Md5: "",
	}
	fmt.Println("pp", p == nil)
}
