import pickle

my_list = [123, 3.14, '小甲鱼', ['another list']]
pickle_file = open('c://code//leetcode//python_code//learn//file//my_list.pkl', "wb")
pickle.dump(my_list, pickle_file)
pickle_file.close()