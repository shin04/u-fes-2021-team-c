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

  function callApi() {
    fetch("https://nufes-teamc.herokuapp.com/users")
      .then(function (res) {
        return res.json();
      })
      .then(function (users) {
        console.log(users);
      });
  }

  callApi();
}