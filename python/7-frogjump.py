def solution(distance, falls):
    if len(falls) ==1 and falls[0] and distance == 1:
        return 0
    elif len(falls) < distance:
        return -1

    positions = set(range(1, distance+1))
    for min, pos in enumerate(falls):
        if pos in positions:
            positions.remove(pos)
            print(positions, pos, min)
            if not positions:
                return min
    return -1




X = 1
#A = [1 ,3 ,1 ,4 ,2 ,3, 5, 4]
A=[1]

print(solution(X , A))