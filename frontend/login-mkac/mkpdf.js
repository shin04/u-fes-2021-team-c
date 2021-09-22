const mkpdf = () => {
    const username = document.mkpdf_form.name.value;
    const number = document.mkpdf_form.number.value;
    coonst img = document.mkpdf_form.img

  
    const data = {
      name: username,
      number: number,
      img: img, 
    };
    const param = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    fetch("./mkpdf", param)
      .then((res) => {
        if( ! res.ok ) {
          throw new Error(`Fetch: ${res.status} ${res.statusText}`);
        }
        return( res.json() );
      })
      .then((json) => {
        console.log(json);
  
      })
      .catch((error)=>{
        alert('エラー');
      });
};

/* fetch のパスどうすればいいですか？ */