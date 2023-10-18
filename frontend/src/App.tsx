import { CssBaseline } from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import { useCookies } from "react-cookie";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import Main from "./components/Main";
import Checks from "./pages/Checks";
import Index from "./pages/Index";
import Login from "./pages/Login";
import CheckScoreboard from "./pages/Scoreboard/Check";
import RoundScoreboard from "./pages/Scoreboard/Round";
import Scoreboard from "./pages/Scoreboard/Scoreboard";
import StatusScoreboard from "./pages/Scoreboard/Status";
import TeamScoreboard from "./pages/Scoreboard/Team";
import { useSystemTheme } from "./themes/Preference";

export default function App() {
  const [cookies, setCookie, removeCookie] = useCookies(["auth"]);

  return (
    <ThemeProvider theme={useSystemTheme()}>
      <CssBaseline />
      <BrowserRouter>
        <Routes>
          <Route
            path='/'
            element={
              <Main cookies={cookies} removeCookie={removeCookie}>
                <Index />
              </Main>
            }
          />
          <Route
            path='/login'
            element={
              <Main cookies={cookies} removeCookie={removeCookie}>
                <Login setCookie={setCookie} />
              </Main>
            }
          />
          <Route
            path='/checks'
            element={
              <Main cookies={cookies} removeCookie={removeCookie}>
                <Checks cookies={cookies} />
              </Main>
            }
          />
          <Route
            path='/scoreboard'
            element={
              <Main cookies={cookies} removeCookie={removeCookie}>
                <Scoreboard />
              </Main>
            }
          />
          <Route
            path='/scoreboard/team/:team'
            element={
              <Main cookies={cookies} removeCookie={removeCookie}>
                <TeamScoreboard />
              </Main>
            }
          />
          <Route
            path='/scoreboard/check/:check'
            element={
              <Main cookies={cookies} removeCookie={removeCookie}>
                <CheckScoreboard />
              </Main>
            }
          />
          <Route
            path='/scoreboard/round/:round'
            element={
              <Main cookies={cookies} removeCookie={removeCookie}>
                <RoundScoreboard />
              </Main>
            }
          />
          <Route
            path='/scoreboard/status/:team/:check'
            element={
              <Main cookies={cookies} removeCookie={removeCookie}>
                <StatusScoreboard />
              </Main>
            }
          />
          <Route path='iframe'>
            <Route path='scoreboard' element={<Scoreboard />} />
            <Route path='scoreboard/team/:team' element={<TeamScoreboard />} />
            <Route
              path='scoreboard/check/:check'
              element={<CheckScoreboard />}
            />
            <Route
              path='scoreboard/round/:round'
              element={<RoundScoreboard />}
            />
            <Route
              path='scoreboard/status/:team/:check'
              element={<StatusScoreboard />}
            />
          </Route>
          <Route
            path='*'
            element={
              <Main cookies={cookies} removeCookie={removeCookie}>
                <Index />
              </Main>
            }
          />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  );
}
