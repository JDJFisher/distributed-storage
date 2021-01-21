
# Third party
import grpc

# Application
from common.helloworld_pb2 import HelloRequest, HelloReply


class GreeterServicer(object):
  def SayHello(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


class Greeter(GreeterServicer):
    def SayHello(self, request, context):
        return HelloReply(message='Hello, %s!' % request.name)


def add_GreeterServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SayHello': grpc.unary_unary_rpc_method_handler(
          servicer.SayHello,
          request_deserializer=HelloRequest.FromString,
          response_serializer=HelloReply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'helloworld.Greeter', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


def serve():
    server = grpc.server()
    add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    print('Hello world')
    # serve()

