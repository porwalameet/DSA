def solution(A):
    print(set(range(1 ,len(A)+2)), set(A), set(range(1 ,len(A)+2)) - set(A))
    return min(set(range(1 ,len(A)+2)) - set(A))

# testcase 1
A = [1, 3, 6, 4, 1, 2]
print (solution(A))

# testcase 2
A = [1, 2, 3]
print (solution(A))

# testcase 3
A = [-1, -3]
print (solution(A))

# testcase 4
A = [2]
print (solution(A))