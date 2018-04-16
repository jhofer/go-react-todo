import axios from "axios";

const data = obj => obj.data;

export default class Api {
  getTodos = () => axios.get("/api/v1/todos/").then(data);
}
