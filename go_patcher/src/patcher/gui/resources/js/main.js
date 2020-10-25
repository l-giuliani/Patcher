document.addEventListener('astilectron-ready', function() {

    document.getElementById('newVersion').addEventListener('change', function () {
        let fileInput = document.getElementById('newVersion');
        let filename = fileInput.files[0].path
        document.getElementById('newVersionLabel').innerHTML=filename
    })

    document.getElementById('precVersion').addEventListener('change', function () {
        let fileInput = document.getElementById('precVersion');
        let filename = fileInput.files[0].path
        document.getElementById('precVersionLabel').innerHTML=filename
    })

    /*document.getElementById('outputFolder').addEventListener('change', function () {
        let fileInput = document.getElementById('outputFolder');
        let path = fileInput.files[0].path
        path = path.substring(0,path.lastIndexOf('\\'))
        document.getElementById('outputFolderLabel').innerHTML=path
    })*/
    
    document.getElementById('generatePatch').addEventListener('click', async function () {
        let newVersion = document.getElementById('newVersion');
        let precVersion = document.getElementById('precVersion');
        //let outputFolder = document.getElementById('outputFolder');
        let patchFileName = document.getElementById('patchFileName');
        let crc32FileName = document.getElementById('crc32FileName');
        
        if(newVersion.files[0] == null || precVersion.files[0] == null 
            /*|| outputFolder.files[0] == null*/ || patchFileName.value == ''
            || crc32FileName.value == ''){
            console.log("error param")
            return 
        }

        let nvFilename = newVersion.files[0].path
        let pvFilename = precVersion.files[0].path
        //let outputPath = outputFolder.files[0].path
        //outputPath = outputPath.substring(0,outputPath.lastIndexOf('\\'))
        let outputPath = "output\\"

        const data = {
            NewVersion: nvFilename,
            PrecVersion: pvFilename,
            Patch: outputPath + "\\" + patchFileName.value,
            Crc32: outputPath + "\\" + crc32FileName.value
        }

        try{
            ViewUtils.showLoadSpinner()
            await Services.generatePatch(data)
            console.log("Success")
            ViewUtils.showSuccessAlert()
            ViewUtils.hideLoadSpinner()
        } catch(err){
            console.log("Error")
            ViewUtils.showErrorAlert()
            ViewUtils.hideLoadSpinner()
        }
    });
})