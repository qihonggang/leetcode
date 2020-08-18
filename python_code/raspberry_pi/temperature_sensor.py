import time


class TemperatureSensor(object):

    def specific_request(self):
       print("温度传感器数据")
       # 温度传感器数据
       data = {'type': 'temperature_sensor', 'time': time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()), 'data': '25',
               'unit': '℃', 'unusual': '0'}
       return data