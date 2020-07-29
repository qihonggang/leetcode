class Ball:
    def setName(self, name):
        self.__name = name
    def kick(self):
        print("我叫 %s ,噢~谁踢我呢？！" % self.__name)

a = Ball()
a.setName("飞火流星")
b = Ball()
b.setName("团队之星")
c = Ball()
c.setName("土豆")
a.kick
b.kick()
c.kick()
print(c._Ball__name)