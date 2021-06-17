"use strict";

{
  const menuItems = document.querySelectorAll(".menu li a");
  const contents = document.querySelectorAll(".content");

  menuItems.forEach(clickedItem => {
    clickedItem.addEventListener("click", e => {
      e.preventDefault();

      menuItems.forEach(item => {
        item.classList.remove("active");
      });
      clickedItem.classList.add("active");

      contents.forEach(content => {
        content.classList.remove("active");
      });
      document.getElementById(clickedItem.dataset.id).classList.add("active");
    });
  });

  // async function callApi() {
  //   const res = await fetch("https://nufes-teamc.herokuapp.com/users");
  //   const users = await res.json();
  //   console.log(users);
  // }

  // callApi();

  // function callApi() {
  //   fetch("https://nufes-teamc.herokuapp.com/users")
  //     .then(function (res) {
  //       return res.json();
  //     })
  //     .then(function (users) {
  //       comsole.log(users);
  //     });
  // }

  // callApi();

  // function callApi() {
  //   const xhr = new XMLHttpRequest();
  //   xhr.open("GET", "https://nufes-teamc.herokuapp.com/users");
  //   xhr.responseType = "json";
  //   xhr.send();
  //   xhr.onload = function () {
  //     console.log(xhr.response);
  //   };
  // }

  // callApi();

  fetch("https://nufes-teamc.herokuapp.com/users")
    .then((res)=>{
      return( res.json() );
    })
    .then((json)=>{
    console.log(json);
    });
}