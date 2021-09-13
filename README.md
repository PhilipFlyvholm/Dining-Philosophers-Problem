# Dining-Philosophers-Problem

## New notes

### Possible algo

1. P requests F1 and F2
2. F1 and F2 add to queue
3. When F is available then send accept to P
4. When both F accept P then P start eating
5. P wait x sec + randomisation
6. When P done dismiss F
7. P wait x sec + randomisation
8. Repeat

## Previous notes

### Philosophor class

- Input channel
- Output channel
- Times eaten/thought
- State

### Fork class

- Input channel
- Output channel
- Times used
- State

### Dining table

- Forks array
- Philosophor array
- Init method

### Note

- Man skal samle to op ad gangen for ikke at ende i deadlock
- Når en person requester gaffler så låser vi hele systemet

### Loop

1. P[2] requests to eats
2. The forks get locked
3. Does nothing if only one is availble
4. Otherwise sets the forks to be unavailble and starts eating