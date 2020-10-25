
const generatePatch = (data) => {
    const request = {
        RequestType: "generate-crc32",
        Data: JSON.stringify(data)
    }
    return new Promise((resolve,reject)=>{
        astilectron.sendMessage(request, function(message) {
            console.log("received " + message)
            if(message == true){
                resolve()
            } else {
                reject()
            }
        });
    })
    
}


const Services = {
    generatePatch
}