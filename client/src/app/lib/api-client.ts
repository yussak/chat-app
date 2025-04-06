import Axios from "axios";

// todo: server componentからだとlocalhost:8080ではなくserver:8080じゃないとアクセすできないので対応する
export const api = Axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
});
