from raspberry_pi.adapter import Adapter
from raspberry_pi.human_sensor import HumanSensor
from raspberry_pi.humidity_sensor import HumiditySensor
from raspberry_pi.target import Target
from raspberry_pi.temperature_sensor import TemperatureSensor
import json

# 根据不同的key，获取需要被适配的类
def get_adaptee_class(key):
    adaptees = {
        'target': Target(),
        'temperature_sensor': TemperatureSensor(),
        'humidity_sensor': HumiditySensor(),
        'human_sensor': HumanSensor()
    }
    return adaptees.get(key, None)

if __name__ == "__main__":
    # 选择传感器类型
    adaptee_type = 'human_sensor'
    # 适配器进行适配
    adapter = Adapter(get_adaptee_class(adaptee_type))
    # 适配后的采集数据操作
    data = adapter.request()
    # 格式化为json格式
    json_data = json.dumps(data)
    print(json_data)
    # 传输到kafka