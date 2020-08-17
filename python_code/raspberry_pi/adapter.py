from raspberry_pi.target import Target


class Adapter(Target):
    def __init__(self, classname):
        self.classname = classname
    def request(self):
        if self.classname.__class__.__name__ == "Target" :
            return self.classname.request()
        else:
            return self.classname.specific_request()
