import { Container } from "@mui/material";
import { SnackbarProvider } from "notistack";
import { useState } from "react";
import { CookieSetOptions } from "universal-cookie";
import Drawer from "./Drawer";
import NavBar from "./NavBar";
import { Outlet } from "react-router";

type props = {
  cookies: {
    auth?: any;
  };
  removeCookie: (name: "auth", options?: CookieSetOptions | undefined) => void;
  mobile: boolean;
};

export default function Main({ cookies, removeCookie, mobile }: props) {
  const [drawerState, setDrawerState] = useState<boolean>(false);
  return (
    <>
      <Drawer
        drawerState={drawerState}
        setDrawerState={setDrawerState}
        removeCookie={removeCookie}
        cookies={cookies}
      />
      <NavBar
        mobile={mobile}
        cookies={cookies}
        setDrawerState={setDrawerState}
      />
      <Container component='main'>
        <SnackbarProvider maxSnack={3}>
          <Outlet />
        </SnackbarProvider>
      </Container>
    </>
  );
}
