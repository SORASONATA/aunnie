def main():
    print(num(100))


def num(n1):
    if n1<=1:
        return 1
    else:
        return n1 + num(n1-1)
    
if __name__ == "__main__":
    main()