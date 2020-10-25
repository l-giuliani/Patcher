import patchLib, utils

def generatePatch(old, new, patch):
    res = patchLib.generate(old, new, patch)
    if(res != 0):
        print("Error Generating Patch")
        return False
    return True

def applyPatch(old, patch, patched):
    res = patchLib.apply(old, patch, patched)
    if(res != 0):
        print("Error Applying Patch")
        return False
    return True
    
def generatePatchWithCrc32(old, new, patch, crc32File):
    patchRes = generatePatch(old, new, patch)
    if (not patchRes):
        return False
    crcRes, crc32 = patchLib.computeCrc32(new)
    if (not crcRes):
        print("Error generating crc32")
        return False
    writeCrcFileRes = utils.writeSmallFile(crc32File, crc32)
    if (not writeCrcFileRes):
        print("Error writing crc32 file")
        return False

def applyPatchWithCrc32(old, patch, patched, crc32File):
    patchRes = applyPatch(old, patch, patched)
    if (not patchRes):
        return False
    crcRes, crc32 = patchLib.computeCrc32(patched)
    if (not crcRes):
        print("Error generating crc32")
        return False
    readCrcFileRes, crcNewFile = utils.readSmallFile(crc32File)
    if(not readCrcFileRes):
        print("Error reading crc32 file")
        return False
    if(crcNewFile == crc32):
        print("Success")
        return True
    print("crc32 not equals")
    return False

    
