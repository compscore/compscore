import { Container } from "@mui/material";
import { SnackbarProvider } from "notistack";
import { ReactNode, useState } from "react";
import { CookieSetOptions } from "universal-cookie";
import Drawer from "./Drawer";
import NavBar from "./NavBar";

type props = {
  children: ReactNode;
  cookies: {
    auth?: any;
  };
  removeCookie: (name: "auth", options?: CookieSetOptions | undefined) => void;
  mobile: boolean;
};

export default function Main({
  children,
  cookies,
  removeCookie,
  mobile,
}: props) {
  const [drawerState, setDrawerState] = useState<boolean>(false);
  return (
    <>
      <Drawer
        drawerState={drawerState}
        setDrawerState={setDrawerState}
        removeCookie={removeCookie}
      />
      <NavBar
        mobile={mobile}
        cookies={cookies}
        setDrawerState={setDrawerState}
      />
      <Container component='main'>
        <SnackbarProvider maxSnack={3}>{children}</SnackbarProvider>
      </Container>
    </>
  );
}
