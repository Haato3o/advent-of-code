def main():
    with open("input", "r") as f:
        input = f.read()
    
    stack = []
    stack_helper = []

    for lines in input.split("\n\n"):
        elf_calories = sum([int(calories) for calories in lines.split("\n")])

        if len(stack) == 0:
            stack.append(elf_calories)
            continue

        if stack[-1] >= elf_calories:
            continue

        while len(stack) > 0 and stack[-1] < elf_calories:
            stack_helper.append(stack.pop())

        stack.append(elf_calories)

        while len(stack) < 3 and len(stack_helper) > 0:
            stack.append(stack_helper.pop())

        stack_helper.clear()
    
    print(sum(stack))

if __name__ == "__main__":
    main()