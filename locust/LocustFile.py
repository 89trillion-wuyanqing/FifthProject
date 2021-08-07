from locust import User, task, events, constant
import time
import websocket
import json
import jsonpath
# from protocol import talk_pb2

def eventType_success(eventType, recvText, total_time):
    events.request_success.fire(request_type="[RECV]",
                                name=eventType,
                                response_time=total_time,
                                response_length=len(recvText))

class WebSocketClient(object):
    _locust_environment = None

    def __init__(self, host):
        self.host = host
        # 连接
        self.ws = websocket.WebSocket()

    def connect(self, burl):
        start_time = time.time()
        try:
            self.conn = self.ws.connect(url=burl)
        except websocket.WebSocketTimeoutException as e:
            total_time = int((time.time() - start_time) * 1000)
            events.request_failure.fire(
                request_type="[Connect]", name='TimeOut', response_time=total_time, exception=e)
        else:
            total_time = int((time.time() - start_time) * 1000)
            events.request_success.fire(
                request_type="[Connect]", name='WebSocket', response_time=total_time, response_length=0)
        return self.conn

    def recv(self):
        return self.ws.recv()

    def send(self, msg):
        self.ws.send(msg)

class WebsocketUser(User):
    abstract = True
    def __init__(self, *args, **kwargs):
        super(WebsocketUser, self).__init__(*args, **kwargs)
        self.client = WebSocketClient(self.host)
        self.client._locust_environment = self.environment

class ApiUser(WebsocketUser):
    host = 'ws://127.0.0.1:8011/ws'
    wait_time = constant(0)

    @task(1)
    def pft(self):
        # wss 地址
        self.url = 'ws://127.0.0.1:8011/ws'
        self.data = {}
        self.client.connect(self.url)

        # # 发送的订阅请求
        # datas = talk_pb2.TaskInfo()
        # datas.code = 1
        # datas.name = "qq"
        # datas.msg = "聊天"
        # datas.cmd = "talk"
        # dataB = datas.SerializeToString()
        # self.client.send(dataB)

        while True:
            # 消息接收计时
            start_time = time.time()
            recv = self.client.recv()
            total_time = int((time.time() - start_time) * 1000)

            # 为每个推送过来的事件进行归类和独立计算性能指标
            try:
                # result = talk_pb2.TaskInfo
                # result.ParseFromString(recv)
                print(recv)
                # recv_j = json.loads(recv)
                # eventType_s = jsonpath.jsonpath(recv_j, expr='$.eventType')
                # eventType_success(eventType_s[0], recv, total_time)
            except websocket.WebSocketTimeoutException as e:
                events.request_failure.fire(request_type="[ERROR] WebSocketConnectionClosedException",
                                            name='Connection is already closed.',
                                            response_time=total_time,
                                            exception=e)
            except:
                print(recv)
                # 正常 OK 响应，或者其它心跳响应加入进来避免当作异常处理
                eventType_success('ok', 'ok', total_time)
                # if 'content' in recv:
                #     eventType_success('ok', 'ok', total_time)