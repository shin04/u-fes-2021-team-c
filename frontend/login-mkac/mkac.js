const mkac = () => {
  const username = document.mkac_form.name.value;
  const password = document.mkac_form.password.value;

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
    return res.json();
  })
  .then((json) => {
    console.log(json);
  });
};