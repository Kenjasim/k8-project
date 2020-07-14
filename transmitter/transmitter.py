# Program to perform a simple HTTP request to a server every 2 seconds

import requests, time, os

class Transmitter():
    """
    Class which sends a simple HTTP request to a proxy to be rerouted
    """

    def __init__(self):
        """
        Constructor class which defines the address and the port
        """

        # The address and port number to access
        try:
            self.address = "http://" + os.environ['TRANSMIT_IP']
            self.port_num = os.environ['TRANSMIT_PORT']
            self.ip = self.address+":"+self.port_num
        except:
            print("ERROR: Ensure you set the TRANSMIT_IP and TRANSMIT_PORT environment variables")
            os._exit(1)


    def request(self):
        """
        Sends a post of "Hello World" every 2 seconds
        """

        #Define the payload to send
        pload = {'String': 'Hello World'}

        # Send every two seconds
        while (True):
            try:
                r = requests.post(self.ip ,data = pload)
            except Exception as e:
                print("Exception: ", e)
            
            time.sleep(2)
        

t = Transmitter()
t.request()
