import pickle

pickle_file = open('c://code//leetcode//python_code//learn//file//my_list.pkl', "rb")
my_list = pickle.load(pickle_file)
print(my_list)