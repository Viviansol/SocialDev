$('#login').on('submit', Login);

function Login(event){
    event.preventDefault();

    $.ajax({
        url: "/login",
        method : "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function (){
        window.location = "/home"
    }).fail(function (err){
        alert("invalid email or password")
    })

}