# pip install grpcio # grpc
# pip install grpcio-tools  
# pip install protobuf
# python3 -m grpc_tools.protoc -I./proto/hello hello.proto --python_out=./grpc/py_client/ --grpc_python_out=./grpc/py_client/ 

# -I 指定proto所在目录
# -m 指定通过protoc生成py文件
# --python_out指定生成py文件的输出路径
# hello.proto 输入的proto文件
import grpc
import hello_pb2, hello_pb2_grpc

def run():
  conn = grpc.insecure_channel('localhost:50052')
  client = hello_pb2_grpc.HelloStub(channel=conn)   # 客户端使用Stub类发送请求,参数为频道,为了绑定链接
  response = client.SayHello(hello_pb2.HelloRequest(name='py grpc!'))   # 返回的结果就是proto中定义的类
  print(response.message)
  
if __name__ == '__main__':
  run()