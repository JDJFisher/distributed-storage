
# Standard library
from __future__ import print_function

# Third party
import grpc

# Application
from common.helloworld_pb2 import HelloRequest, HelloReply


class GreeterStub(object):

  def __init__(self, channel):
    self.SayHello = channel.unary_unary(
        '/helloworld.Greeter/SayHello',
        request_serializer=HelloRequest.SerializeToString,
        response_deserializer=HelloReply.FromString,
    )


def run():
    with grpc.insecure_channel('server:50051') as channel:
        stub = GreeterStub(channel)
        response = stub.SayHello(HelloRequest(name='you'))

    print("Greeter client received: " + response.message)


if __name__ == '__main__':
    print('Hello world')
    run()