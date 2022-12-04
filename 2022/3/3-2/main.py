def get_priority(c: str) -> int:
    if ord('z') >= ord(c) >= ord('a'):
        return ord(c) % 96
    else:
        return (ord(c) % 64) + 26
    
def main():
    with open("input", "r") as file:
        input = file.readlines()

    result = 0
    cases = [line.strip("\n") for line in input]
    chunks = [cases[3*i: 3*i+3] for i in range(0, len(cases) // 3)]
    
    for chunk in chunks:
        badge = [a for a in chunk[0] if a in [b for b in chunk[1] if b in set(chunk[2])]][0]
        result += get_priority(badge)
    
    print(result)

if __name__ == "__main__":
    main()