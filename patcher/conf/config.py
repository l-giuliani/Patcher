conf = {
    #'dllPath':  'C:/Users/lgiuliani/Projects/Patcher/dist/lib/patchLib.dll',
    'dllPath':  'lib/patchLib.dll',

    'projects': {
        'fer': {
            'enable': True,
            'operationalPath': 'C:/patchTests',
            'exePath': 'C:/patchTests',
            'exeFileName': 'new.txt'
        }
    },
    
    'generate': {
        'destFile': 'C:/patchTests/patch.patch'
    }   
}