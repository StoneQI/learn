#include <iostream>

#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/transport/TSocket.h>
#include <thrift/transport/TTransportUtils.h>

#include "gen-cpp/Calculator.h"

using namespace std;
using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;

using namespace tutorial;
using namespace shared;

int main() {

  // TBufferedTransport 为数据缓存Transport，没有发送数据的能力，因此需要 TSocket 去发送数据
  std::shared_ptr<TTransport> socket(new TSocket("localhost", 9090));

  // TBufferedTransport 结构为这个 boost::scoped_array<uint8_t> rBuf_; 当调用 TBufferedTransport的 flush时，底层调用TSocket的wirite写入数据到fd，然后调用TSocket的flush，默认什么也不做，TSocket也是一个Transport
  std::shared_ptr<TTransport> transport(new TBufferedTransport(socket));

  // 将transport 传入 protocol，即TBinaryProtocol 序列化后数据的保存方式
  std::shared_ptr<TProtocol> protocol(new TBinaryProtocol(transport));

  // 传入client 会吧client中 oprot_ = protocol
  CalculatorClient client(protocol);

  try {
    transport->open();

    client.ping();
    cout << "ping()" << endl;

    cout << "1 + 1 = " << client.add(1, 1) << endl;

    // // 下面为生成的代码
    // {
    //   int32_t CalculatorClient::add(const int32_t num1, const int32_t num2)
    //   {
    //     send_add(num1, num2);
    //     return recv_add();
    //   }

    //   void CalculatorClient::send_add(const int32_t num1, const int32_t num2)
    //   {
    //     int32_t cseqid = 0;
    //     // 写入消息头
    //     oprot_->writeMessageBegin("add", ::apache::thrift::protocol::T_CALL, cseqid);

    //     // 传的参数会生成一个结构体，然后生成write和read方法。如果有
    //     Calculator_add_pargs args;
    //     args.num1 = &num1;
    //     args.num2 = &num2;
    //     // 调用write方法写入到oprot_中
    //     args.write(oprot_);
    //     // 具体方法如下，依次写入每一个参数
    //     // uint32_t Calculator_add_pargs::write(::apache::thrift::protocol::TProtocol* oprot) const {
    //     // uint32_t xfer = 0;
    //     // ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
    //     // xfer += oprot->writeStructBegin("Calculator_add_pargs");

    //     // xfer += oprot->writeFieldBegin("num1", ::apache::thrift::protocol::T_I32, 1);
    //     // xfer += oprot->writeI32((*(this->num1)));
    //     // xfer += oprot->writeFieldEnd();

    //     // xfer += oprot->writeFieldBegin("num2", ::apache::thrift::protocol::T_I32, 2);
    //     // xfer += oprot->writeI32((*(this->num2)));
    //     // xfer += oprot->writeFieldEnd();

    //     // xfer += oprot->writeFieldStop();
    //     // xfer += oprot->writeStructEnd();


    //     // 写入消息结束
    //     oprot_->writeMessageEnd();

    //     // TTransport 写入结束
    //     oprot_->getTransport()->writeEnd();
    //     // 发送消息
    //     oprot_->getTransport()->flush();
    //   }
    // }

    Work work;
    work.op = Operation::DIVIDE;
    work.num1 = 1;
    work.num2 = 0;

    try {
      client.calculate(1, work);
      cout << "Whoa? We can divide by zero!" << endl;
    } catch (InvalidOperation& io) {
      cout << "InvalidOperation: " << io.why << endl;
      // or using generated operator<<: cout << io << endl;
      // or by using std::exception native method what(): cout << io.what() << endl;
    }

    work.op = Operation::SUBTRACT;
    work.num1 = 15;
    work.num2 = 10;
    int32_t diff = client.calculate(1, work);
    // // 生成代码
    // {
    //   int32_t CalculatorClient::calculate(const int32_t logid, const Work& w)
    //   {
    //     send_calculate(logid, w);
    //     return recv_calculate();
    //   }

    //   void CalculatorClient::send_calculate(const int32_t logid, const Work& w)
    //   {
    //     int32_t cseqid = 0;
    //     oprot_->writeMessageBegin("calculate", ::apache::thrift::protocol::T_CALL, cseqid);

    //     // 参数赋值
    //     Calculator_calculate_pargs args;
    //     args.logid = &logid;
    //     args.w = &w;

    //     // 参数具体写入
    //     args.write(oprot_);
    //     // 逻辑如下
    //     {
    //       uint32_t Calculator_calculate_pargs::write(::apache::thrift::protocol::TProtocol* oprot) const {
    //       uint32_t xfer = 0;
    //       ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
    //       xfer += oprot->writeStructBegin("Calculator_calculate_pargs");

    //       xfer += oprot->writeFieldBegin("logid", ::apache::thrift::protocol::T_I32, 1);
    //       xfer += oprot->writeI32((*(this->logid)));
    //       xfer += oprot->writeFieldEnd();

    //       xfer += oprot->writeFieldBegin("w", ::apache::thrift::protocol::T_STRUCT, 2);
    //       // 对应结构体的写入方法，生成代码中实现
    //       xfer += (*(this->w)).write(oprot);
    //       xfer += oprot->writeFieldEnd();

    //       xfer += oprot->writeFieldStop();
    //       xfer += oprot->writeStructEnd();
    //       return xfer;
    //       }
    //     }

    //     oprot_->writeMessageEnd();
    //     oprot_->getTransport()->writeEnd();
    //     oprot_->getTransport()->flush();
    //   }
    // }
    cout << "15 - 10 = " << diff << endl;

    // Note that C++ uses return by reference for complex types to avoid
    // costly copy construction
    SharedStruct ss;
    client.getStruct(ss, 1);
    cout << "Received log: " << ss << endl;

    transport->close();
  } catch (TException& tx) {
    cout << "ERROR: " << tx.what() << endl;
  }
}
