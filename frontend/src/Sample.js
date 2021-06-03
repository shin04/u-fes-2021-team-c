import React from "react";

function Sample() {
  constructor(props) {
    super(props);
    this.state = {name:"ファイル形式"};
    this.state = {count: 0};
  }
handleClick(){
  this.setState({count:this.state.count +1});
}
handleClick(name){
  this.setState({name:name});
}
  render() {
    return (
      <div>
         <h1>ファイルを{this.state.name}に変更</h1>
         <button onClick={() => {this.setState({name:'PDF'})}}>PDF</button>
         <button onClick={() => {this.setState({name:'JPEG'})}}>JPEG</button>
        <button onClick={()=>{this.handleClick()}}>+</button>
      </div>
    );
  }
}

export default Sample;
