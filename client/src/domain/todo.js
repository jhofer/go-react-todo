import { observable } from "mobx";

export class Todo {
  @observable ID;
  @observable CreatedAt;
  @observable DeletedAt;
  @observable Completed;
  @observable Title;
  @observable UpdatedAt;
}
