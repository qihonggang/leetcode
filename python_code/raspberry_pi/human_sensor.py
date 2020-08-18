import time

class HumanSensor(object):
    def specific_request(self):
       print("人体传感器数据")
       # 人体传感器数据
       data = {'type': 'human_sensor', 'time': time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()), 'data': True,
               'unit': None, 'unusual': '0'}
       return data