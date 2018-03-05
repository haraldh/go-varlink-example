package orgexamplemore

//go:generate $GOPATH/bin/varlink-generator ./org.example.more.varlink

import (
	"github.com/varlink/go/varlink"
)

type Interface struct {
	varlink.InterfaceDefinition
	Service *varlink.Service
}

func NewInterface() Interface {
	return Interface{InterfaceDefinition: NewInterfaceDefinition()}
}

func (intf *Interface) TestMore(call varlink.Call) error {
	var in TestMore_In
	err := call.GetParameters(&in)
	if err != nil {
		return call.ReplyError("org.varlink.service.InvalidParameter", varlink.InvalidParameter_Error{Parameter: "parameters"})
	}

	// FIXME: Fill me in
	return call.ReplyError("org.varlink.service.MethodNotImplemented", varlink.MethodNotImplemented_Error{Method: "TestMore"})

	return call.Reply(&varlink.ServiceOut{
		Parameters: TestMore_Out{
		// FIXME: Fill me in
		},
	})
}

func (intf *Interface) StopServing(call varlink.Call) error {
	if intf.Service != nil {
		intf.Service.Stop()
	}
	return call.Reply(&varlink.ServiceOut{})
}

func (intf *Interface) Ping(call varlink.Call) error {
	var in Ping_In

	err := call.GetParameters(&in)
	if err != nil {
		return call.ReplyError("org.varlink.service.InvalidParameter", varlink.InvalidParameter_Error{Parameter: "parameters"})
	}

	return call.Reply(&varlink.ServiceOut{
		Parameters: Ping_Out{
			in.Ping,
		},
	})
}
