function createModel() {
    fetch(
        'http://localhost:1111/upload',
        {
            method: 'post',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                data: $('#code').val(),
            }),
        }
    ).then(response => {
        console.log(response)
    })
        .catch(error => {
            console.log(error)
        });
}
