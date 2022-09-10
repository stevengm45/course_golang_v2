from collections import Counter


numeros = [2,3,3,4,4,5,1]

freq = Counter(numeros)

for i in freq:
    if(freq[i] > 1):
        print(i)

mp = [0] * 100
n = len(numeros)
for i in range(0, n):
    mp[numeros[i]] += 1

for i in range(0, n):
    if(mp[numeros[i]] > 1):
        print(numeros[i])

        mp[numeros[i]] = 0