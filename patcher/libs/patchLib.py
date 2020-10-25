import config
from ctypes import*
import zlib, os

patchDll = None

def init():
    global patchDll
    if(patchDll is None):
        patchDll = cdll.LoadLibrary(config.conf["dllPath"])

def generate(old, new, patch):
    init()
    res = patchDll.GeneratePatch(c_char_p(old.encode('utf-8')), c_char_p(new.encode('utf-8')), c_char_p(patch.encode('utf-8')))
    #res = patchDll.GeneratePatch(c_char_p("C:/patchTests/new".encode('utf-8')), c_char_p("C:/patchTests/old".encode('utf-8')), c_char_p("C:/patchTests/patch".encode('utf-8')))
    return res

def apply(old, patch, patched):
    init()
    res = patchDll.ApplyPatch(c_char_p(old.encode('utf-8')), c_char_p(patched.encode('utf-8')), c_char_p(patch.encode('utf-8')))
    return res

def computeCrc32(fileName):
    try:
        if not os.path.exists(fileName):
            return False, None
        file = open(fileName,"rb")
        prev = 0
        for eachLine in file:
            prev = zlib.crc32(eachLine, prev)
        file.close()
        return True, "%X"%(prev & 0xFFFFFFFF)
    except:
        return False, None