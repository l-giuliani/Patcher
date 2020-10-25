import sys
import patchService

serviceDict = {
    'generate': patchService.generatePatch,
    'apply': patchService.applyPatch,
    'generate-crc32': patchService.generatePatchWithCrc32,
    'apply-crc32': patchService.applyPatchWithCrc32
}

def main():
    print ("Awesome Patcher")
    if(len(sys.argv) < 2):
        print("Number params error")
        return
    if (sys.argv[1] in serviceDict):
        if(len(sys.argv) < 5):
            print("Number params error")
            return
        if(len(sys.argv) > 5):
            serviceDict[sys.argv[1]](sys.argv[2], sys.argv[3], sys.argv[4], sys.argv[5])
        else:
            serviceDict[sys.argv[1]](sys.argv[2], sys.argv[3], sys.argv[4])
    elif (sys.argv[1] == "update-service"):
        print("Service not implemented yet")
    else:
        print("Option error")

def test():
    patchService.applyPatchWithCrc32("C:/patchTests/patched_fer.patchg", 2, 3)

main()