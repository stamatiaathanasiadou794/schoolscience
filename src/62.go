def is_prime(num):
    if num <= 1:
        return False
    for i in range(2, int(num**0.5) + 1):
        if num % i == 0:
            return False
    return True

def sum_of_primes(n):
    total = 0
    for i in range(2, n+1):
        if is_prime(i):
            total += i
    return total

print(sum_of_primes(5))  # Example output: 16
