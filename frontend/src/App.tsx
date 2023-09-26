import { ThemeProvider } from "@mui/material/styles";
import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import { useSystemTheme } from "./themes/Preference";
import {BrowserRouter, Routes} from 'react-router-dom';
import { SnackbarProvider } from 'notistack';
import { Route } from 'react-router-dom';
import Index from './pages/Index';
import Login from './pages/Login';
import NavBar from './components/NavBar';


export default function App() {
  return (
    <ThemeProvider theme={useSystemTheme()}>
      <CssBaseline />
      <NavBar/>
      <Container component='main' maxWidth='xs'>
        <BrowserRouter>
          <SnackbarProvider maxSnack={3}>
            <Routes>
              <Route path="/" element={<Index />} />
              <Route path="/login" element={<Login />} />
            </Routes>
          </SnackbarProvider>
        </BrowserRouter>
      </Container>
    </ThemeProvider>
  );
}
