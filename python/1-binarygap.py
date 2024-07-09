def f(n):
    cnt = 0
    result = 0
    found_one = False

    i = n

    while i:
        if i & 1 == 1:
            if (found_one is False):
                found_one = True
            else:
                result=max(result, cnt)
            cnt = 0
        else:
            cnt += 1
        i >>= 1
    
    return result

def sol2(n):
    return len(max(format(n,'b').strip('0').split('1')))

n=32
print(sol2(n),bin(n)[2:])
