fib(0,1).
fib(1,1).
fib(NTH,RES) :-
  NTH > 1,
  N1 is NTH-1,
  N2 is NTH-2,
  fib(N1,RES1),
  fib(N2,RES2),
  RES is RES1+RES2.
