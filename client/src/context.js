import Api from "./api";
import { observable, computed } from "mobx";

export default class Context {
  api = new Api();
  constructor() {
    this.onStartup();
  }

  onStartup = async () => {
    this.todos = await this.api.getTodos();
  };

  @observable todos = [];
  @computed
  get unfinishedTodoCount() {
    return this.todos.filter(todo => !todo.Completed).length;
  }
}
