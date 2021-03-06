const mkac = () => {
  const username = document.mkac_form.name.value;
  const password = document.mkac_form.password.value;
  const username_bool = (username == username.match(/\w+/g));
  const password_bool = (password == password.match(/\w+/g));
  if (username_bool&&passwordbool) {
    const data = { 
      name: username, 
      password: password, 
    };
    const param = {
      method: 'POST', 
      headers: {
        'Content-Type': 'application/json', 
      }, 
      body: JSON.stringify(data), 
    };
    fetch('https://nufes-teamc.herokuapp.com/user', param)
    .then((res) => {
      if( ! res.ok ) {
        throw new Error(`Fetch: ${res.status} ${res.statusText}`);
      }
      return( res.json() );
    })
    .then((json) => {
      console.log(json);
      console.log('アカウント作成に成功しました．')
      window.location.href = './home.html';

    })
    .catch((error)=>{
      alert('エラー');
    });
  }else{
    console.error(username, password);
  }
};