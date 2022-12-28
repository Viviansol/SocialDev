$('#form-registration').on('submit', createUser);

function createUser (event ){
    event.preventDefault();

    if($('#password').val() !== $('#confirm-password').val()){
        alert("Passwords are different!")
        return;
    }

    $.ajax({
        url:"/users",
        method: "POST",
        data:{
            name :$('#name').val(),
            email: $('#email').val(),
            nickName:$('#nick').val(),
            password:$('#password').val(),
        }
    }).done(function (){
        alert("user registrated!")
    }).fail(function (){
        alert("cannot register user")
    })

}