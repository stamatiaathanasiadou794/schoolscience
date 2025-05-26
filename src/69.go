def example_function():
    # Randomly choose 2 random numbers and add them together
    total = 20 + random.randint(1, 10)
    
    # Add a specific value to the total
    if random.random() < 0.5:
        total += 3
    
    return total

def main():
    # Example usage of the function
    result = example_function()
    print(result)

if __name__ == "__main__":
    main()
