def sol1(l):
    minval=abs(sum(l[:1])-sum(l[1:]))
    val =0
    for i in range(1, len(l)-1):
        val=abs(sum(l[:i])-sum(l[i:]))
        minval = min(val, minval)
    return minval


def sol2(A):

    head = A[0]
    tail = sum(A[1:])
    diff = abs(head - tail)
    for number in A[1:-1]:
        head += number
        tail -= number
        diff = min(diff , abs(head -tail))

    return diff
#l = [3,1,2,4,3]
l = [-1000,1000]
print(sol2(l))