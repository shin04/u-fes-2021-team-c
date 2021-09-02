function hello() {
  console.log("hello!");
}
function ocha() {
  fetch("https://nufes-teamc.herokuapp.com/users")
    .then((res) => {
      return res.json();
    })
    .then((json) => {
      console.log(json);
    });
}

const fetchForm = document.querySelector("login_form");
const btn = document.querySelector(".button");

// function green_tea() {
//   const data = new FormData(fetchForm);
//   console.log(data);
//   console.log(JSON.stringify(data));
//   const param = {
//     method: "POST",
//     headers: {
//       "Content-Type": "application/json",
//     },
//     body: JSON.stringify(data),
//   };
//   fetch("https://nufes-teamc.herokuapp.com/login", param)
//     .then((res) => {
//       return res.json();
//     })
//     .then((json) => {
//       console.log(json);
//     });
// }
// btn.addEventListener("click", green_tea, false);

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
