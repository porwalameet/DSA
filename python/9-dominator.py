def solution(A):
    if len(A) == 0:
        return -1

    sort_a = sorted(A)
    l = len(A) // 2
    domi_candidate = sort_a[l]
    print(domi_candidate, A.count(domi_candidate))
    if A.count(domi_candidate) >= l:
        return domi_candidate
    return -1

A = [3 ,4 ,3 ,2 ,3 ,-1 ,3 ,3]
print(solution(A))

# testcase 2
A = []
print(solution(A))

# testcase 3
A = [0 ,1]
print(solution(A))

# testcase 4
A = [0,-2147483648 ,-2147483648 ,2147483648, 2147483648, 2147483648]
print(solution(A))