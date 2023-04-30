import matplotlib.pyplot as plt
import numpy as np

X = np.array([int(1e2), int(1e3), int(1e4), int(1e5)])
QUIC_Y = np.array([0.23, 2.58, 22.56, 242.66])
TCP_Y = np.array([0.45, 8.15, 35.32, 374.22])

QUIC_Y = np.array([0.20, 1.85, 18.28, 186.64])



plt.xlabel('number of tries')
plt.ylabel('time (in seconds)')
plt.plot(X, QUIC_Y, 'b', label='QUIC')
plt.plot(X, TCP_Y, 'g', label='TCP (no tls)')

plt.legend(loc="upper left")

plt.savefig(fname='quic-vs-tcp-conn-establishment.png')

