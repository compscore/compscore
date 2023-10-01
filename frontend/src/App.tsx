import { ThemeProvider } from "@mui/material/styles";
import { Container, CssBaseline } from "@mui/material";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { SnackbarProvider } from "notistack";
import { useCookies } from "react-cookie";

import { useSystemTheme } from "./themes/Preference";
import Index from "./pages/Index";
import Login from "./pages/Login";
import NavBar from "./components/NavBar";
import Scoreboard from "./pages/Scoreboard";
import TeamScoreboard from "./pages/TeamScoreboard";

export default function App() {
  const [cookies, setCookie] = useCookies(["auth"]);

  return (
    <ThemeProvider theme={useSystemTheme()}>
      <CssBaseline />
      <NavBar cookies={cookies} setCookie={setCookie} />
      <Container component='main'>
        <BrowserRouter>
          <SnackbarProvider maxSnack={3}>
            <Routes>
              <Route path='/' element={<Index />} />
              <Route path='/login' element={<Login setCookie={setCookie} />} />
              <Route path='/scoreboard' element={<Scoreboard />} />
              <Route
                path='/scoreboard/team/:team'
                element={<TeamScoreboard />}
              />
            </Routes>
          </SnackbarProvider>
        </BrowserRouter>
      </Container>
    </ThemeProvider>
  );
}
