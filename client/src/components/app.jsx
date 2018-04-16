import * as React from "react";
import TodoList from "./todolist";
import Context from "../context";

const context = new Context();

export default class App extends React.Component {
  render() {
    return (
      <div>
        <TodoList context={context} />
      </div>
    );
  }
}
