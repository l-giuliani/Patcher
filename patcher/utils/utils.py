import os

def readSmallFile(fileName):
    data = None
    try:
        if not os.path.exists(fileName):
            return False, None
        with open(fileName, 'r') as file:
            data = file.read()
        return True, data
    except:
        return False, None

def writeSmallFile(fileName, data):
    try:
        file = open(fileName,"w+")
        file.write(data)
        file.close()
        return True
    except:
        return False