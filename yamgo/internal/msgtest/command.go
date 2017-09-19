package msgtest

import (
	"bytes"

	"github.com/10gen/mongo-go-driver/bson"
	"github.com/10gen/mongo-go-driver/yamgo/private/msg"
)

// CreateCommandReply creates a msg.Reply from the BSON response from the server.
func CreateCommandReply(cmd interface{}) *msg.Reply {
	doc, _ := bson.Marshal(cmd)
	reply := &msg.Reply{
		NumberReturned: 1,
		DocumentsBytes: doc,
	}

	// encode it, then decode it to handle the internal workings of msg.Reply
	codec := msg.NewWireProtocolCodec()
	var b bytes.Buffer
	err := codec.Encode(&b, reply)
	if err != nil {
		panic(err)
	}
	resp, err := codec.Decode(&b)
	if err != nil {
		panic(err)
	}

	return resp.(*msg.Reply)
}