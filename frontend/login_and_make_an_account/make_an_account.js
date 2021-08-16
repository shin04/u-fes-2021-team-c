function green_tea() {
    const fetchForm = document.querySelector('.form');
const data = new FormData(fetchForm);
const param  = {
  method: "POST",
  headers: {
    "Content-Type": "application/json"
  },
  body: JSON.stringify(data)
};
fetch("https://nufes-teamc.herokuapp.com/user", param)
  .then((res)=>{
    return( res.json() );
  })
  .then((json)=>{
    console.log(json);
});
}
const btn = document.querySelector('.button');
btn.addEventListener('click', green_tea, false);