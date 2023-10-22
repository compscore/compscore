import { Container } from "@mui/material";
import { SnackbarProvider } from "notistack";
import { useState } from "react";
import { Outlet } from "react-router";
import { cookies, removeCookie, setCookie } from "../models/Cookies";
import Drawer from "./Drawer";
import NavBar from "./NavBar";

type props = {
  cookies: cookies;
  removeCookie: removeCookie;
  setCookie: setCookie;
  mobile: boolean;
};

export default function Main({
  cookies,
  removeCookie,
  mobile,
  setCookie,
}: props) {
  const [drawerState, setDrawerState] = useState<boolean>(false);
  return (
    <>
      <Drawer
        drawerState={drawerState}
        setDrawerState={setDrawerState}
        setCookie={setCookie}
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
