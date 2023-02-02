function createModel() {
    fetch(
        'http://localhost:1111/upload',
        {
            method: 'post',
            headers: {
                'Accept': 'text/plain',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                data: $('#code').val(),
            }),
        }
    ).then(response => {
        console.log(response);
        let area = document.getElementById('output');
        area.value = '';
        area.value = response;
        alert(response);
    })
        .catch(error => {
            console.log(error)
        });
}
