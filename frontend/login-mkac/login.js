const login = () => {
  const username = document.login_form.name.value;
  const password = document.login_form.password.value;
  const username_bool = (username == username.match(/\w+/g));
  const password_bool = (password == password.match(/\w+/g));
  if (username_bool&&passwordbool) {
    const data = {
      name: username,
      password: password,
    };
    const param = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    fetch("https://nufes-teamc.herokuapp.com/login", param)
      .then((res) => {
        if( ! res.ok ) {
          throw new Error(`Fetch: ${res.status} ${res.statusText}`);
        }
        return( res.json() );
      })
      .then((json) => {
        console.log(json);
        window.location.href = './mkpdf.html';

      })
      .catch((error)=>{
        alert('エラー');
      });
  }else{
    console.error(username, password);
  };
};
