
def total_fuel_requirements(module_masses):
    """
    Puzzle part one solution
    """
    return sum([mass // 3 - 2 for mass in module_masses])

def real_total_fuel_requirements(module_masses):
    """
    Puzzle part two solution
    """
    def fuel_of_mass(mass):
        return mass // 3 - 2

    total = 0
    for mass in module_masses:
        prospective_fuel = fuel_of_mass(mass)
        while prospective_fuel > 0:
            total += prospective_fuel
            prospective_fuel = fuel_of_mass(prospective_fuel)
    return total

def main():
    """
    AOC 2019 Day 1 Driver
    """
    module_masses = None
    with open('input') as file:
        module_masses = [int(mass) for _, mass in enumerate(file)]

    print('Total fuel requirements for all modules on spacecraft:',
          total_fuel_requirements(module_masses))
    print('Total fuel requirements for all modules on spacecraft taking ' +
          'into account the fuel\'s fuel requirements:',
          real_total_fuel_requirements(module_masses))

if __name__ == '__main__':
    main()

