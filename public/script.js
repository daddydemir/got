function createModel() {
    fetch(
        'http://localhost:7777/upload',
        {
            method: 'post',
            headers: {
                'Accept': 'text/plain',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                data: $('#code').val(),
                pkg: $('#package').val(),
            }),
        }
    ).then(response => {
        return response.text();
    }).then(function (data){
        let area = document.getElementById('output');
        area.value = '';
        area.value = data;
    });
}
