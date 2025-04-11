# Create an empty list to store numbers
numbers = []

# Define a function to calculate the sum of a given list of numbers
def sum_numbers(numbers):
    total_sum = 0
    for number in numbers:
        total_sum += number
    return total_sum

# Add some numbers to the list
numbers.append(3)
numbers.append(5)
numbers.append(7)

# Call the function and print the result
print(sum_numbers(numbers))
