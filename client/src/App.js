import React from "react"
import "./App.css";
import {container} from "semantic-ui-react";
import TodoList from "../To-Do-List";

function App(){
  return(
    <div>
    <container>
    <TodoList />
    </container>
    </div>
  );
  }

export default App;