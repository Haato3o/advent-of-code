def get_priority(c: str) -> int:
    if ord('z') >= ord(c) >= ord('a'):
        return ord(c) % 96
    else:
        return (ord(c) % 64) + 26
    
def main():
    with open("input", "r") as file:
        input = file.readlines()

    result = 0
    cases = [(x[0:len(x)//2], x[len(x)//2:]) for x in [line.strip("\n") for line in input]]
    for low_half, high_half in cases:
        seen = set()
        seen_other = set()
        for c in low_half:
            seen.add(c)
        
        for c in high_half:
            if c in seen and c not in seen_other:
                result += get_priority(c)
                seen_other.add(c)
    
    print(result)

if __name__ == "__main__":
    main()