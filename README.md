# Magic Hash

Magic hash calculator program written in go

Magic hashes are special hash digests that are when converted to a hexadecimal string, start with `0e` and contain only digits afterwards. This causes these hashes to be invalidly compared by weakly typed languages when used as a password digest when logging in.

Example in PHP:

```php
"0e462097431906509019562988736854" == "0e830400451993494058024219903391"
```

this returns true, which is the same as:

```php
md5("240610708") == "0e830400451993494058024219903391"
```

So if the hash match is not validated using the `password_verify()` or `hash_equals()` methods, there is a slight possibiliy that an attacker can take over an account, given that the account owner's password creates a magic hash.

## Calculation Speed

The chances of generating a magic hash exponentially decreases with the bit size of the hash digest. It can be calculated as follows:

$b$ = number of bits
```math
\frac{1}{16} * \frac{1}{16} * \frac{10}{16}^{\frac{b}{4}-2}
```

In the case of a 128 bit hash (eg: MD5), the chance for a magic hash is: 

$$ \approx \frac{1}{1,329,227}, $$

in case of a 256 bit hash (eg: SHA-256), however the chance is:

$$ \approx \frac{1}{4,523,128,485,832}. $$

This is $3,402,826$ times less than in case of md5, which is also exponantiated by the lenghtier calculation time of the SHA-256 algorithm compared to MD5.

A simple python brute-force solution ran for 32 seconds for MD5 before termination. On a M4 MacBook Pro, 11 core CPU. This means that even a relatively simper SHA1 calculation would have taken hours, or even days.

The go solution however, completed in 27 seconds, which is about a 16% improvement just by switching languages, and in 8 seconds when using multithreading with 8 cores, which is a further 337% improvement. This makes the calculations more feasible.

## Results

|Algorithm|Input|Digest|Time|
|-|-|-|-|
|CRC32/IEEE|`dubniczky-cs`|`0e984686`|0.24s|
|MD5|`dubniczky-37HAY`|`0e167783116409945604446462415162`|6s|
|SHA1|`dubniczky-k6yW74`|`0e42764555227832861957804727564640004574`|31.5m|
