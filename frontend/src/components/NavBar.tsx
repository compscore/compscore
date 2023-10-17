import {
  AppBar,
  Avatar,
  Box,
  Button,
  Link,
  Toolbar,
  Typography,
} from "@mui/material";
import jwt_decode from "jwt-decode";
import { JWT } from "../models/JWT";

type Props = {
  cookies: {
    auth?: any;
  };
  removeCookie: (
    name: "auth",
    options?: import("universal-cookie").CookieSetOptions
  ) => void;
};

export default function NavBar({ cookies, removeCookie }: Props) {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position='static'>
        <Toolbar>
          <Button href='/'>
            <Avatar src='/compscore.svg' />
          </Button>
          <Box sx={{ m: 1 }} />
          <Link
            href='/'
            color='inherit'
            underline='none'
            variant='h6'
            sx={{ flexGrow: 1 }}
          >
            Compscore
          </Link>
          {cookies.auth ? (
            <>
              <Typography variant='h6'>
                {(jwt_decode(cookies.auth) as JWT).Team} -{" "}
                {(jwt_decode(cookies.auth) as JWT).Role}
              </Typography>
              <Box sx={{ m: 1 }} />
              <Button
                color='inherit'
                onClick={() => {
                  removeCookie("auth");
                  window.location.href = "/";
                }}
              >
                Logout
              </Button>
            </>
          ) : (
            <Button color='inherit' href='/login'>
              Login
            </Button>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}
