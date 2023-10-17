import { Container, CssBaseline } from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import { SnackbarProvider } from "notistack";
import { useState } from "react";
import { useCookies } from "react-cookie";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import Drawer from "./components/Drawer";
import NavBar from "./components/NavBar";
import CheckScoreboard from "./pages/CheckScoreboard";
import Checks from "./pages/Checks";
import Index from "./pages/Index";
import Login from "./pages/Login";
import RoundScoreboard from "./pages/RoundScoreboard";
import Scoreboard from "./pages/Scoreboard";
import StatusScoreboard from "./pages/StatusScoreboard";
import TeamScoreboard from "./pages/TeamScoreboard";
import { useSystemTheme } from "./themes/Preference";

export default function App() {
  const [cookies, setCookie, removeCookie] = useCookies(["auth"]);
  const [drawerState, setDrawerState] = useState<boolean>(false);

  return (
    <ThemeProvider theme={useSystemTheme()}>
      <CssBaseline />
      <Drawer
        drawerState={drawerState}
        setDrawerState={setDrawerState}
        removeCookie={removeCookie}
      />
      <NavBar cookies={cookies} setDrawerState={setDrawerState} />
      <Container component='main'>
        <BrowserRouter>
          <SnackbarProvider maxSnack={3}>
            <Routes>
              <Route path='/' element={<Index />} />
              <Route path='/login' element={<Login setCookie={setCookie} />} />
              <Route path='/checks' element={<Checks cookies={cookies} />} />
              <Route path='/scoreboard' element={<Scoreboard />} />
              <Route
                path='/scoreboard/team/:team'
                element={<TeamScoreboard />}
              />
              <Route
                path='/scoreboard/check/:check'
                element={<CheckScoreboard />}
              />
              <Route
                path='/scoreboard/round/:round'
                element={<RoundScoreboard />}
              />
              <Route
                path='/scoreboard/status/:team/:check'
                element={<StatusScoreboard />}
              />
              <Route path='*' element={<Index />} />
            </Routes>
          </SnackbarProvider>
        </BrowserRouter>
      </Container>
    </ThemeProvider>
  );
}
