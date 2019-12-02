
def total_fuel_requirements(module_masses):
    """
    Puzzle part one solution
    """
    return sum([mass // 3 - 2 for mass in module_masses])

def main():
    """
    AOC 2019 Day 1 Driver
    """
    module_masses = None
    with open('input') as file:
        module_masses = [int(mass) for _, mass in enumerate(file)]

    print('Total fuel requirements for all modules on spacecraft:',
          total_fuel_requirements(module_masses))

if __name__ == '__main__':
    main()

