import time
from concurrent import futures
import grpc
# import hello_pb2_grpc, hello_pb2
from hello_pb2 import HelloResponse
from hello_pb2_grpc import HelloServicer, add_HelloServicer_to_server

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

class TestService(HelloServicer):
    '''
    继承HelloServicer,实现hello方法
    '''
    def __init__(self):
        pass
      
    def SayHello(self, request, context):
        '''
        具体实现SayHello的方法，并按照pb的返回对象构造HelloResponse返回
        :param request:
        :param context:
        :return:
        '''
        result = request.name + " this is gprc test service"
        list_result = {"12": 1232}
        return HelloResponse(message=str(result))

def run():
    '''
    模拟服务启动
    :return:
    '''
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_HelloServicer_to_server(TestService(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    print("start service...")
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)
        
if __name__ == '__main__':
    run()