# Elliptic Curve Cryptography Golang

This repository contains the implementation of several Elliptic Curve Cryptography Algorithms. The algorithms rely on [Go-Ristretto](https://github.com/bwesterb/go-ristretto) to manage points and scalars.
I put each algorithm in a separate sub-module, so it can be used independently with other ones.
Supported algorithms include: [ElGamal](https://link.springer.com/book/10.1007/b97644), [Pedersen Commitment Scheme](https://link.springer.com/chapter/10.1007/3-540-46766-1_9), [ECIES](https://en.wikipedia.org/wiki/Integrated_Encryption_Scheme)

I initially implement these algorithms to learn elliptic curve cryptography and use them in my academic papers. Since there are many attacks in Elliptic Curve Cryptography Implementation that are not tested, this module should not be used in the industry.
