import jwt_decode from "jwt-decode";
import { JWT } from "../models/JWT";

type props = {
  cookies: {
    auth?: any;
  };
};

export default function Admin({ cookies }: props) {
  if (cookies.auth == undefined) {
    window.location.href = "/login";
  }

  if ((jwt_decode(cookies.auth) as JWT).Role !== "admin") {
    window.location.href = "/";
  }

  return <div>Admin</div>;
}
