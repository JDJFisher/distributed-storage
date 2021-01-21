# Standard library
from concurrent import futures
import logging
import time
import sys

# Third party
import grpc

# Application
from helloworld_pb2_grpc import GreeterStub, GreeterServicer, add_GreeterServicer_to_server
from helloworld_pb2 import HelloReply, HelloRequest


class Greeter(GreeterServicer):
    def SayHello(self, request, context):
        return HelloReply(message='Hello, %s!' % request.name)


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    logging.info('Starting...')

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:10000')
    server.start()

    logging.info('Waiting...')
    time.sleep(1)
    logging.info('Sending...')

    with grpc.insecure_channel('localhost:10000') as channel:
        stub = GreeterStub(channel)
        response = stub.SayHello(HelloRequest(name='you'))
        logging.info("Received: " + response.message)
    
    server.wait_for_termination()


