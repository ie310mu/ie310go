// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package thrift

import (
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/ie310mu/ie310go/forks/github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type Ie310goThirftService interface {
  // Parameters:
  //  - Data
  Invoke(ctx context.Context, data string) (r string, err error)
}

type Ie310goThirftServiceClient struct {
  c thrift.TClient
}

func NewIe310goThirftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *Ie310goThirftServiceClient {
  return &Ie310goThirftServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewIe310goThirftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *Ie310goThirftServiceClient {
  return &Ie310goThirftServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewIe310goThirftServiceClient(c thrift.TClient) *Ie310goThirftServiceClient {
  return &Ie310goThirftServiceClient{
    c: c,
  }
}

func (p *Ie310goThirftServiceClient) Client_() thrift.TClient {
  return p.c
}
// Parameters:
//  - Data
func (p *Ie310goThirftServiceClient) Invoke(ctx context.Context, data string) (r string, err error) {
  var _args0 Ie310goThirftServiceInvokeArgs
  _args0.Data = data
  var _result1 Ie310goThirftServiceInvokeResult
  if err = p.Client_().Call(ctx, "Invoke", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type Ie310goThirftServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Ie310goThirftService
}

func (p *Ie310goThirftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *Ie310goThirftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *Ie310goThirftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewIe310goThirftServiceProcessor(handler Ie310goThirftService) *Ie310goThirftServiceProcessor {

  self2 := &Ie310goThirftServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["Invoke"] = &ie310goThirftServiceProcessorInvoke{handler:handler}
return self2
}

func (p *Ie310goThirftServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x3

}

type ie310goThirftServiceProcessorInvoke struct {
  handler Ie310goThirftService
}

func (p *ie310goThirftServiceProcessorInvoke) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := Ie310goThirftServiceInvokeArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Invoke", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := Ie310goThirftServiceInvokeResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Invoke(ctx, args.Data); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Invoke: " + err2.Error())
    oprot.WriteMessageBegin("Invoke", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Invoke", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Data
type Ie310goThirftServiceInvokeArgs struct {
  Data string `thrift:"data,1" db:"data" json:"data"`
}

func NewIe310goThirftServiceInvokeArgs() *Ie310goThirftServiceInvokeArgs {
  return &Ie310goThirftServiceInvokeArgs{}
}


func (p *Ie310goThirftServiceInvokeArgs) GetData() string {
  return p.Data
}
func (p *Ie310goThirftServiceInvokeArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Ie310goThirftServiceInvokeArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Data = v
}
  return nil
}

func (p *Ie310goThirftServiceInvokeArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Invoke_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Ie310goThirftServiceInvokeArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("data", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:data: ", p), err) }
  if err := oprot.WriteString(string(p.Data)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.data (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:data: ", p), err) }
  return err
}

func (p *Ie310goThirftServiceInvokeArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Ie310goThirftServiceInvokeArgs(%+v)", *p)
}

// Attributes:
//  - Success
type Ie310goThirftServiceInvokeResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewIe310goThirftServiceInvokeResult() *Ie310goThirftServiceInvokeResult {
  return &Ie310goThirftServiceInvokeResult{}
}

var Ie310goThirftServiceInvokeResult_Success_DEFAULT string
func (p *Ie310goThirftServiceInvokeResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return Ie310goThirftServiceInvokeResult_Success_DEFAULT
  }
return *p.Success
}
func (p *Ie310goThirftServiceInvokeResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *Ie310goThirftServiceInvokeResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Ie310goThirftServiceInvokeResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *Ie310goThirftServiceInvokeResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Invoke_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Ie310goThirftServiceInvokeResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *Ie310goThirftServiceInvokeResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Ie310goThirftServiceInvokeResult(%+v)", *p)
}


