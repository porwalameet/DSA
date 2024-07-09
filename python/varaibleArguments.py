import logging
LOG_FORMAT = "'%(asctime)s'- '%(levelname)s' - '%(message)s'"
logging.basicConfig(level=logging.DEBUG, format=LOG_FORMAT )
def sum(*num) -> int:
    sum=0
    for n in num:
        sum +=n
    logging.info(sum)
    return sum

def intro(**data):
    for k,v in data.items():
        print(f'{k}: {v}')


print(f'Sum: {sum(1,4,2,2,20,4)}')
intro(Firstname="Sita", Lastname="Sharma", Age=22, Phone=1234567890)
intro(Firstname="John", Lastname="Wood", Email="johnwood@nomail.com", Country="Wakanda", Age=25, Phone=9876543210)

