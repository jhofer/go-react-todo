import * as React from "react";
import * as PropTypes from "prop-types";
import Todo from "./todo.jsx";
import { observer } from "mobx-react";

@observer
class TodoList extends React.Component {
  render() {
    return (
      <div>
        <ul>
          {this.props.context.todos.map(todo => (
            <Todo todo={todo} key={todo.id} />
          ))}
        </ul>
        Tasks left: {this.props.context.unfinishedTodoCount}
      </div>
    );
  }
}

export default TodoList;
