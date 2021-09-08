const login = () => {
  const username = document.login_form.name.value;
  const password = document.login_form.password.value;

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
    credentials = 'same-origin'
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
      window.location.href = '/Users/mattsunkun/web4/frontend/login-mkac/mkpdf.html';

    })
    .catch((error)=>{
      console.error('エラー');
    });
};
