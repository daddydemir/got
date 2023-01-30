function createModel() {
    alert($('#code').val());
    fetch(
        'http://localhost:1111/upload',
        {
            method: 'post',
            body: JSON.stringify({
                data: $('#code').val(),
            }),
        }
    ).then(response => {
        console.log(response)
    })
        .catch(error => {
            console.log(error)
        })
}