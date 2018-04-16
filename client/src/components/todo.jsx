import * as React from "react";
import { observer } from "mobx-react";

const Todo = observer(({ todo }) => (
  <li>
    <input
      type="checkbox"
      checked={todo.Completed}
      onClick={() => (todo.Completed = !todo.Completed)}
    />
    {todo.Title}
  </li>
));

export default Todo;
