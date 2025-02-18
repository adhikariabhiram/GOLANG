import numpy as np
import matplotlib.pyplot as plt
l = [100, 110, 110, 110, 120, 120, 130, 140, 140, 150, 170, 220]
q1, q2, q3 = np.quantile(l, [0.25, 0.5, 0.75], method='midpoint')
print(q1, q2, q3)
iqr = q3-q1
print(iqr)

l2 = [100, 110, 110, 110, 120, 120, 130, 140, 140, 150, 170, 220]
l3 = [l, l2]
fig = plt.figure(figsize=(10, 7))
bp = plt.boxplot(l3, vert=0)
plt.show()

# np.random.seed(10)
# data = np.random.normal(50, 20, 200)
# fig = plt.figure(figsize=(10, 7))
# plt.boxplot(data, vert=0)
# plt.show()


def listmax(l):
    l1=[]
    for i in l:
        l1.append(sum(i)/len(i))
    return l1


def same(l1,l2):
    m=min(len(l1),len(l2))
    for i in range(m):
        if l1[i]!=l2:
            return False
    return True