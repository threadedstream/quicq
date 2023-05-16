import matplotlib.pyplot as plt
import numpy as np


X = np.array([1000, 2500, 5000])
QUIC_Y = np.array([1.39, 4.13, 4.91])
TCP_Y = np.array([1.46, 4.20, 5.81])

plt.xlabel('number of messages')
plt.ylabel('time (in minutes)')
plt.plot(X, QUIC_Y, 'b', label='QUIC')
plt.plot(X, TCP_Y, 'g', label='TCP')

plt.legend(loc="upper left")

plt.savefig(fname='quic-vs-tcp-message-processing.png')