package network

var ()

//TODO fix tests
// func TestConnect(t *testing.T) {

// 	A := NewLocalTransport("A")
// 	B := NewLocalTransport("B")

// 	A.Connect(B)
// 	B.Connect(A)

// 	assert.Equal(t, A., B)
// 	assert.Equal(t, B.peers[A.addr], A)
// }

// func TestSendMessage(t *testing.T) {

// 	A := NewLocalTransport("A")
// 	B := NewLocalTransport("B")

// 	A.Connect(B)
// 	B.Connect(A)

// 	msg := []byte("Hello world")
// 	A.SendMessage(B.addr, msg)

// 	assert.Nil(t, A.SendMessage(B.addr, msg))

// 	rpc := <-B.Consume()

// 	assert.Equal(t, rpc.Payload, msg)
// 	assert.Equal(t, rpc.From, A.addr)
// }
