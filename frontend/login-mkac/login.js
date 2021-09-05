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
  };
  fetch("https://nufes-teamc.herokuapp.com/login", param)
    .then((res) => {
      return res.json();
    })
    .then((json) => {
      console.log(json);
    });
};
