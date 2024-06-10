from itertools import product
import hashlib

prefix = "dubniczky-"
alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


def is_magic_hash(digest: str) -> bool:
    if not digest.startswith("0e"):
        return False
    if not digest[2:].isdigit():
        return False
    
    return True

def md5_search_magic_hash(prefix, alphabet, max_length):
    for i in range(1, max_length):
        for c in product(alphabet, repeat=i):
            s = prefix + ''.join(c)
            digest = hashlib.md5(s.encode('utf-8')).hexdigest()
            if is_magic_hash(digest):
                return (s, digest)

if __name__ == "__main__":
    result = md5_search_magic_hash(prefix, alphabet, 16)
    print(result[0], " -> ", result[1])
