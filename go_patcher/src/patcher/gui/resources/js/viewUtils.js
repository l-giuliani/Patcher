
const showLoadSpinner = () => {
    document.getElementById('cover-spin').style.display = ''
}

const hideLoadSpinner = () => {
    document.getElementById('cover-spin').style.display = 'none'
}

const showSuccessAlert = () => {    
    document.getElementById('alert-success').style.display = ''
    document.getElementById('alert-error').style.display = 'none'
}

const showErrorAlert = () => {
    document.getElementById('alert-success').style.display = 'none'
    document.getElementById('alert-error').style.display = ''
}

const ViewUtils = {
    showLoadSpinner: showLoadSpinner,
    hideLoadSpinner: hideLoadSpinner,
    showSuccessAlert: showSuccessAlert,
    showErrorAlert: showErrorAlert
}