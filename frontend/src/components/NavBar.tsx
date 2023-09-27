import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Button from "@mui/material/Button";
import { Avatar, Typography } from "@mui/material";
import Link from "@mui/material/Link";
import jwt_decode from "jwt-decode";
import { JWT } from "../models/JWT";

type Props = {
  cookies: {
    auth?: any;
  };
  setCookie: (
    name: "auth",
    value: any,
    options?: import("universal-cookie").CookieSetOptions | undefined
  ) => void;
};

export default function ButtonAppBar({ cookies, setCookie }: Props) {
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
                {(jwt_decode(cookies.auth) as JWT).Team}
              </Typography>
              <Box sx={{ m: 1 }} />
              <Button
                color='inherit'
                onClick={() => {
                  setCookie("auth", "", { maxAge: -1 });
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
