import React from "react";
import {
  Box,
  Link,
  Tooltip,
  Typography,
  TextField,
  Button,
  Container,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { LoginSuccess, LoginFailure } from "../models/Login";
import { CookieSetOptions } from "universal-cookie";
import PasswordInput from "../components/PasswordInput";

type Props = {
  setCookie: (
    name: "auth",
    value: any,
    options?: CookieSetOptions | undefined
  ) => void;
};

export default function Login({ setCookie }: Props) {
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const data = JSON.stringify({
      username: event.currentTarget.username.value,
      password: event.currentTarget.password.value,
    });

    fetch("http://localhost:8080/api/login", {
      method: "POST",
      body: data,
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(async (res) => {
        if (res.status === 200) {
          let response = (await res.json()) as LoginSuccess;
          enqueueSnackbar("Logged in", { variant: "success" });

          setCookie("auth", response.token, {
            path: response.path,
            domain: response.domain,
            secure: response.secure,
            httpOnly: response.httpOnly,
            expires: new Date(response.expiration * 1000),
          });

          window.location.href = "/";
        } else {
          let response = (await res.json()) as LoginFailure;

          enqueueSnackbar(response.error, { variant: "error" });
        }
      })
      .catch((err) => {
        enqueueSnackbar("Encountered an error" + err, { variant: "error" });
        console.log(err);
      });
  };

  return (
    <Container component='main' maxWidth='xs'>
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component='h1' variant='h3'>
          Sign in
        </Typography>
        <Box component='form' onSubmit={handleSubmit}>
          <TextField
            margin='normal'
            required
            fullWidth
            id='username'
            label='Team'
            name='username'
            autoFocus
          />
          <PasswordInput
            margin='normal'
            required
            name='password'
            label='Password'
            id='password'
          />
          <Tooltip title='Reach out to Black Team' arrow>
            <Link
              sx={{
                textDecoration: "none",
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
              }}
            >
              <Typography variant='body2' color='text.secondary'>
                Forgot password?
              </Typography>
            </Link>
          </Tooltip>
          <Button
            type='submit'
            fullWidth
            variant='contained'
            sx={{ mt: 3, mb: 2 }}
          >
            Sign In
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
