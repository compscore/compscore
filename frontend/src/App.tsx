import { ThemeProvider } from "@mui/material/styles";
import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import { useSystemTheme } from "./themes/Preference";
import { BrowserRouter, Routes } from "react-router-dom";
import { SnackbarProvider } from "notistack";
import { Route } from "react-router-dom";

import Index from "./pages/Index";
import Login from "./pages/Login";
import NavBar from "./components/NavBar";
import Scoreboard from "./pages/Scoreboard";

import { useCookies } from "react-cookie";

export default function App() {
  const [cookies, setCookie] = useCookies(["auth"]);

  return (
    <ThemeProvider theme={useSystemTheme()}>
      <CssBaseline />
      <NavBar cookies={cookies} setCookie={setCookie} />
      <Container component='main' maxWidth='xs'>
        <BrowserRouter>
          <SnackbarProvider maxSnack={3}>
            <Routes>
              <Route path='/' element={<Index />} />
              <Route path='/login' element={<Login setCookie={setCookie} />} />
              <Route path='/scoreboard' element={<Scoreboard />} />
            </Routes>
          </SnackbarProvider>
        </BrowserRouter>
      </Container>
    </ThemeProvider>
  );
}
