def sol(l):
    n = len(l)+1
    expected_sums = (n*(n+1)) // 2
    return expected_sums- sum(l)

l = [2]
print(sol(l))