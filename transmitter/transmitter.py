# Program to perform a simple HTTP request to a server every 2 seconds

import requests, time

class Transmitter():
    """
    Class which sends a simple HTTP request to a proxy to be rerouted
    """

    def __init__(self):
        """
        Constructor class which defines the address and the port
        """

        # The address and port number to access
        self.address = "http://11.0.0.32:"
        self.port_num = "30488"


    def request(self):
        """
        Sends a post of "Hello World" every 2 seconds
        """

        #Define the payload to send
        pload = {'String': 'Hello World'}

        # Send every two seconds
        while (True):
            try:
                r = requests.post(self.address+self.port_num,data = pload)
            except Exception as e:
                print("Exception: ", e)
            
            time.sleep(2)
        

t = Transmitter()
t.request()