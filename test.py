import numpy as np
import random
random.seed(10)
mat=[]
vec=[]
for i in range(10):
        mat.append([])
        vec.append(random.randint(0,10))
        for j in range(10):
            mat[i].append(random.randint(0,100))
print("matrix is: \n", np.array(mat))
print("vector is: \n",np.array(vec))
res=np.array(mat)@np.array(vec)
print("Result is \n", res)