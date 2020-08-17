import time

class HumiditySensor(object):

    def specific_request(self):
       print("湿度传感器数据")
       # 湿度传感器数据
       data = {'type': 'humidity_sensor', 'time': time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()), 'data': '25', 'unit': '%rh' }
       return data
