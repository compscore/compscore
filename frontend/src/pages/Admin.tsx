import { Box, Button, ButtonGroup, Container, Typography } from "@mui/material";
import jwt_decode from "jwt-decode";
import { useState } from "react";
import AuthenticateAs from "../components/Admin/AuthenticateAs";
import EngineState from "../components/Admin/EngineState";
import PasswordReset from "../components/Admin/PasswordReset";
import { cookies, setCookie } from "../models/Cookies";
import { JWT } from "../models/JWT";

type props = {
  cookies: cookies;
  setCookie: setCookie;
};

export default function Admin({ cookies, setCookie }: props) {
  if (cookies.auth == undefined) {
    window.location.href = "/login";
  }

  if ((cookies.auth && (jwt_decode(cookies.auth) as JWT)).Role !== "admin") {
    window.location.href = "/";
  }

  const [state, setState] = useState("password");

  return (
    <Container component='main' maxWidth='sm'>
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component='h1' variant='h3'>
          Admin Panel
        </Typography>
        <Box sx={{ m: 2 }} />
        <ButtonGroup variant='contained' fullWidth sx={{ display: "flex" }}>
          <Button
            disabled={state === "password"}
            sx={{ flexGrow: 1 }}
            onClick={() => {
              setState("password");
            }}
          >
            Password Reset
          </Button>
          <Button
            disabled={state === "authenticate"}
            sx={{ flexGrow: 1 }}
            onClick={() => {
              setState("authenticate");
            }}
          >
            Authenticate As
          </Button>
          <Button
            disabled={state === "engine"}
            sx={{ flexGrow: 1 }}
            onClick={() => {
              setState("engine");
            }}
          >
            Engine State
          </Button>
        </ButtonGroup>
        <Box sx={{ m: 2 }} />
        {state === "password" ? (
          <PasswordReset />
        ) : state === "authenticate" ? (
          <AuthenticateAs setCookie={setCookie} />
        ) : state === "engine" ? (
          <EngineState />
        ) : (
          "Unknown State"
        )}
      </Box>
    </Container>
  );
}
