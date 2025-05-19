import numpy as np
import matplotlib.pyplot as plt
import igraph as ig
from igraph import Graph
fig, ax = plt.subplots()


# Define a graph by specifying edges
g = Graph(n=10, edges=[[0,1],[2,3],[0,3],[0,2],[1,5],[2,4],[4,5],[4,1],[6,7],[4,6],[1,3],[3,5],[2,6],[7,9],[8,9],[7,8]])

# Define g to be random graph. size =  number of nodes
#g = Graph.Watts_Strogatz(dim=1, size=16, nei=2, p=0.3)


# Plot graph g
#ig.plot(g, target=ax,vertex_label=[str(i) for i in range(g.vcount())])
#plt.show()


# Define an old scores vector
# all ones means the old scores are evenly distributed
#x = np.ones(g.vcount(), dtype=float)

# this x tests how well the algorithm limits the scores of a new region. E.g. the graph has 7 nodes with score evenly distributed, and then 3 more nodes join
x = [1,1,1,1,1,1,1,0,0,0]


# Get the nxn adjacency matrix of the graph g
A = np.array(g.get_adjacency().data)
#print(A)


# Get the transition matrix W which defines a random walk on the graph
def compute_W(A):
    A = np.array(A)
    n = A.shape[0]
    d = np.sum(A, axis=1)
    W = np.zeros((n, n))
    for i in range(n):
        for j in range(n):
            delta = 1 if i == j else 0
            W[i, j] = 0.5 * (delta + (A[i, j] / d[i] if d[i] != 0 else 0))
    return W


# Get 
def compute_P(W, alpha, tol=1e-3, max_iter=1000):
    W = np.array(W)
    n = W.shape[0]
    P = np.eye(n)
    for _ in range(max_iter):
        P_new = alpha * np.eye(n) + (1 - alpha) * P @ W
        if np.linalg.norm(P_new - P, ord='fro') < tol:
            break
        P = P_new
    else:
        print("Warning: P did not converge within max_iter")
    return P


def compute_Q(P):
    P = np.array(P)
    Q = np.argsort(-P, axis=1)
    return Q

def compute_J(Q):
    Q = np.array(Q)
    n = Q.shape[0]
    J = np.zeros((n, n, n), dtype=int)
    for i in range(n):
        for k in range(n):
            J[i, k, i] = 1
            top_k = Q[i, :k]
            J[i, k, top_k] = 1
    return J

def compute_s(J, x):
    J = np.array(J)
    x = np.array(x)
    n = J.shape[0]
    s = np.zeros((n, n), dtype=int)
    for i in range(n):
        for k in range(n):
            j_vec = J[i, k]
            j_comp = 1 - j_vec
            if np.dot(j_vec, x) <= np.dot(j_comp, x):
                s[i, k] = 1
    return s

def compute_y(J, s, W):
    J = np.array(J)
    s = np.array(s)
    W = np.array(W)
    n = s.shape[0]
    y = np.full(n, np.inf)

    for i in range(n):
        for k in range(n):
            if s[i, k] == 1:
                j_vec = J[i, k]
                j_comp = 1 - j_vec
                numerator = j_comp @ W @ j_vec
                count_ones = np.sum(j_vec)
                if count_ones > 0:
                    value = numerator / count_ones
                    if value < y[i]:
                        y[i] = value
        if np.isinf(y[i]):
            y[i] = 0.0  # fallback
    return y



W = compute_W(A)
#print(W)

P = compute_P(W, 0.5, tol=1e-5, max_iter=1000)
#print(P)

Q = compute_Q(P)
#print(Q)

J = compute_J(Q)
#print(J)

s = compute_s(J, x)
#print(s)

y = compute_y(J, s, W)
y_round = np.round(y, 3)



# Plot graph g
ig.plot(g, target=ax, vertex_label=[str([i,y_round[i]]) for i in range(g.vcount())])
plt.show()
