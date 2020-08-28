import axios from "axios";

axios.defaults.baseURL = "http://127.0.0.1:8888";
axios.defaults.timeout = 10000;
axios.defaults.withCredentials = true;

export default axios