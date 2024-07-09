def sol1(l):
    x = 0
    for n in l:
        x = x^n
    return x


def sol2(l):
    s = set()
    for n in l:
        s.add(n) if n not in s else s.remove(n)
    return(s)



l=[2,4,3,2,4,1,1,10,3,7,7]
print(sol2(l))