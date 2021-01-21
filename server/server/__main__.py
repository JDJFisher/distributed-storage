# Standard library
from concurrent import futures
import logging
import time
import sys
import threading

# Third party
import grpc
from flask import Flask

# Application
from helloworld_pb2_grpc import GreeterStub, GreeterServicer, add_GreeterServicer_to_server
from helloworld_pb2 import HelloReply, HelloRequest


class Greeter(GreeterServicer):
    def SayHello(self, request, context):
        return HelloReply(message='Hello, %s!' % request.name)


#REST API for communicating with the clients
app = Flask(__name__)

@app.route("/")
def hello():
    return "hello, world!"


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

    #Run the REST API within a thread (it's blocking otherwise)
    threading.Thread(target=app.run(host='0.0.0.0', port=5000, debug=True, use_reloader=False)).start()

    with grpc.insecure_channel('localhost:10000') as channel:
        stub = GreeterStub(channel)
        response = stub.SayHello(HelloRequest(name='you'))
        logging.info("Received: " + response.message)
    
    server.wait_for_termination()


